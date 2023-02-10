package mpesa

import (
	"fmt"
	"github.com/techcraftlabs/base"
	"strings"
)

func (eps *Endpoints) Get(requestType requestType) string {
	switch requestType {
	case sessionID:
		return eps.AuthEndpoint

	case pushPay:
		return eps.PushEndpoint

	case disburse:
		return eps.DisburseEndpoint
	}

	return ""
}

func (c *Client) makeInternalRequest(requestType requestType, payload interface{}, opts ...base.RequestOption) *base.Request {
	baseURL := c.Conf.BasePath
	endpoints := c.Conf.Endpoints
	edps := endpoints
	url := appendEndpoint(baseURL, edps.Get(requestType))
	method := requestType.Method()

	return base.NewRequest(requestType.String(), method, url, payload, opts...)
}

func appendEndpoint(url string, endpoint string) string {
	url, endpoint = strings.TrimSpace(url), strings.TrimSpace(endpoint)
	urlHasSuffix, endpointHasPrefix := strings.HasSuffix(url, "/"), strings.HasPrefix(endpoint, "/")

	bothTrue := urlHasSuffix && endpointHasPrefix
	bothFalse := !urlHasSuffix && !endpointHasPrefix
	notEqual := urlHasSuffix != endpointHasPrefix

	if notEqual {
		return fmt.Sprintf("%s%s", url, endpoint)
	}

	if bothFalse {
		return fmt.Sprintf("%s/%s", url, endpoint)
	}

	if bothTrue {
		endp := strings.TrimPrefix(endpoint, "/")

		return fmt.Sprintf("%s%s", url, endp)
	}

	return ""
}
