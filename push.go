package mpesa

import (
	"context"
	"net/http"
)

var _ PushCallbackHandler = (*PushCallbackFunc)(nil)

type PushRequest struct {
}

type PushResponse struct {
}

type push interface {
	Push(ctx context.Context, m Mode, request PushRequest) (PushResponse, error)
}

type callbackServer interface {
	CallbackServeHTTP(writer http.ResponseWriter, request *http.Request)
}

type PushCallbackHandler interface {
	HandleCallback(request PushCallbackRequest) (PushCallbackResponse, error)
}

type PushCallbackFunc func(request PushCallbackRequest) (PushCallbackResponse, error)

func (p PushCallbackFunc) HandleCallback(request PushCallbackRequest) (PushCallbackResponse, error) {
	return p(request)
}
