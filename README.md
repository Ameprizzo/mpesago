# mpesa



## example
```go

package main

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/mpesa"
	"net/http"
	"os"
	"time"
)

func main() {

	var (
		opts []mpesa.ClientOption
	)

	ctx,cancel := context.WithTimeout(context.Background(),time.Minute)
	defer cancel()
	config := &mpesa.Config{
		Endpoints:              nil,
		Name:                   "",
		Version:                "",
		Description:            "",
		BasePath:               "",
		Market:                 0,
		Platform:               0,
		APIKey:                 "",
		PublicKey:              "",
		SessionLifetimeMinutes: 0,
		ServiceProvideCode:     "",
		TrustedSources:         nil,
	}

	debugOption := mpesa.WithDebugMode(true)
	marketOption := mpesa.WithMarket(mpesa.GhanaMarket)
	//marketOption := mpesa.WithMarket(mpesa.TanzaniaMarket)
	platformOption := mpesa.WithApiPlatform(mpesa.OPENAPI)
	//platformOption := mpesa.WithApiPlatform(mpesa.SANDBOX)
	loggerOption := mpesa.WithLogger(os.Stderr)
	httpOption := mpesa.WithHTTPClient(http.DefaultClient)

	opts = append(opts,debugOption,marketOption,platformOption,loggerOption,httpOption)

	c := mpesa.NewClient(config,opts...)

	sessionID, err := c.SessionID(ctx)
	if err != nil {
		return
	}

	fmt.Printf("session id: %s\n",sessionID)
	http.HandleFunc("/callbacks/mpesa",c.CallbackServeHTTP)
	request := mpesa.Request{
		ThirdPartyID: "",
		Reference:    "",
		Amount:       1000.79,
		MSISDN:       "",
		Description:  "",
	}
	disburseResponse, err := c.Disburse(ctx,request)
	if err != nil {
		return
	}

	fmt.Printf("disburse response: %v\n",disburseResponse)

	pushResponse, err := c.PushAsync(ctx,request)
	if err != nil {
		return
	}

	fmt.Printf("push response: %v\n",pushResponse)
}

```
