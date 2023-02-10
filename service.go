package mpesa

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/techcraftlabs/base"
)

var (
	_ service = (*Client)(nil)
)

type (
	service interface {
		QueryTx(ctx context.Context, req QueryTxParams) (QueryTxResponse, error)
		SessionID(ctx context.Context) (response SessionResponse, err error)
		PushAsync(ctx context.Context, request Request) (PushAsyncResponse, error)
		Disburse(ctx context.Context, request Request) (DisburseResponse, error)
		CallbackServeHTTP(w http.ResponseWriter, r *http.Request)
	}

	// Config contains details initialize in mpesa portal
	// Applications require the following details:
	//•	Application Name – human-readable name of the application
	//•	Version – version number of the application, allowing changes in API products to be managed in different versions
	//•	Description – Free text field to describe the use of the application
	//•	APIKey – Unique authorisation key used to authenticate the application on the first call. API Keys need to be encrypted in the first “Generate Session API Call” to create a valid session key to be used as an access token for future calls. Encrypting the API Key is documented in the GENERATE SESSION API page
	//•	SessionLifetime – The session key has a finite lifetime of availability that can be configured. Once a session key has expired, the session is no longer usable, and the caller will need to authenticate again.
	//•	TrustedSources – the originating caller can be limited to specific IP address(es) as an additional security measure.
	//•	Products / Scope / Limits – the required API products for the application can be enabled and limits defined for each call.
	Config struct {
		Endpoints              *Endpoints
		Name                   string
		Version                string
		Description            string
		BasePath               string
		Market                 Market
		Platform               Platform
		APIKey                 string
		PublicKey              string
		SessionLifetimeMinutes int64
		ServiceProvideCode     string
		TrustedSources         []string
	}

	Endpoints struct {
		AuthEndpoint     string
		PushEndpoint     string
		DisburseEndpoint string
		QueryEndpoint    string
	}

	Client struct {
		Conf              *Config
		base              *base.Client
		encryptedAPIKey   *string
		sessionID         *string
		sessionExpiration time.Time
		pushCallbackFunc  PushCallbackHandler
		requestAdapter    *requestAdapter
		rp                base.Replier
		rv                base.Receiver
	}
)

func (c *Client) QueryTx(ctx context.Context, req QueryTxParams) (QueryTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewClient(conf *Config, callbacker PushCallbackHandler, opts ...ClientOption) *Client {
	enc := new(string)
	ses := new(string)

	client := new(Client)

	basePath := conf.BasePath

	client = &Client{
		Conf:              conf,
		base:              base.NewClient(),
		encryptedAPIKey:   enc,
		sessionID:         ses,
		sessionExpiration: time.Now(),
		pushCallbackFunc:  callbacker,
	}

	for _, opt := range opts {
		opt(client)
	}

	platform := client.Conf.Platform
	market := client.Conf.Market

	platformStr, marketStr := platform.String(), market.URLContextValue()
	p := fmt.Sprintf("https://%s/%s/ipg/v2/%s/", basePath, platformStr, marketStr)
	client.Conf.BasePath = p
	client.requestAdapter = &requestAdapter{
		platform:            platform,
		market:              market,
		serviceProviderCode: conf.ServiceProvideCode,
	}

	rp := base.NewReplier(client.base.Logger, client.base.DebugMode)
	rv := base.NewReceiver(client.base.Logger, client.base.DebugMode)
	client.rp = rp
	client.rv = rv
	return client
}

func (c *Client) SessionID(ctx context.Context) (response SessionResponse, err error) {

	token, err := c.getEncryptionKey()
	if err != nil {
		return response, err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	var opts []base.RequestOption
	headersOpt := base.WithRequestHeaders(headers)
	opts = append(opts, headersOpt)
	re := c.makeInternalRequest(sessionID, nil, opts...)
	res, err := c.base.Do(ctx, re, &response)
	if err != nil {
		return response, err
	}

	resErr := res.Error
	if resErr != nil {
		return SessionResponse{}, fmt.Errorf("could not fetch session id: %w", resErr)
	}

	//save the session id
	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not fetch session id: %s", response.OutputErr)
		return response, err1
	}

	sessLifeTimeMin := c.Conf.SessionLifetimeMinutes
	sessID := response.ID
	up := time.Duration(sessLifeTimeMin) * time.Minute
	expiration := time.Now().Add(up)
	c.sessionExpiration = expiration
	c.sessionID = &sessID

	return response, nil
}

func (c *Client) PushAsync(ctx context.Context, request Request) (response PushAsyncResponse, err error) {
	sess, err := c.checkSessionID()
	if err != nil {
		return response, err
	}
	token, err := encryptKey(sess, c.Conf.PublicKey)
	if err != nil {
		return response, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	payload, err := c.requestAdapter.adapt(pushPay, request)
	if err != nil {
		return PushAsyncResponse{}, err
	}

	var opts []base.RequestOption
	headersOpt := base.WithRequestHeaders(headers)
	opts = append(opts, headersOpt)
	re := c.makeInternalRequest(pushPay, payload, opts...)
	res, err := c.base.Do(ctx, re, &response)

	if err != nil {
		return response, err
	}
	fmt.Printf("pushpay response: %s: %v\n", pushPay.String(), res)

	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not perform c2b single stage request: %s", response.OutputErr)
		return response, err1
	}

	return response, nil
}

func (c *Client) Disburse(ctx context.Context, request Request) (response DisburseResponse, err error) {
	sess, err := c.checkSessionID()
	if err != nil {
		return response, err
	}
	token, err := encryptKey(sess, c.Conf.PublicKey)
	if err != nil {
		return response, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	payload, err := c.requestAdapter.adapt(disburse, request)
	if err != nil {
		return DisburseResponse{}, err
	}

	var opts []base.RequestOption
	headersOpt := base.WithRequestHeaders(headers)
	opts = append(opts, headersOpt)
	re := c.makeInternalRequest(disburse, payload, opts...)
	res, err := c.base.Do(ctx, re, &response)

	if err != nil {
		return response, err
	}
	fmt.Printf("disburse response: %s: %v\n", disburse.String(), res)

	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not perform disburse request: %s", response.OutputErr)
		return response, err1
	}

	return response, nil
}

func (c *Client) CallbackServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	body := new(PushCallbackRequest)
	_, err := c.rv.Receive(ctx, "mpesa push callback", request, body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	reqBody := *body

	resp, err := c.pushCallbackFunc.HandleCallback(reqBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	hs := base.WithMoreResponseHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	response := base.NewResponse(200, resp, hs)
	c.rp.Reply(writer, response)
}
