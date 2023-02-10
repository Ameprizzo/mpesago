package async

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	xio "github.com/techcraftlabs/base/io"
)

type PushRequest struct {
	OriginalConversationID   string `json:"input_OriginalConversationID,omitempty"`
	TransactionID            string `json:"input_TransactionID,omitempty"`
	ResultCode               string `json:"input_ResultCode,omitempty"`
	ResultDesc               string `json:"input_ResultDesc,omitempty"`
	ThirdPartyConversationID string `json:"input_ThirdPartyConversationID,omitempty"`
}

type PushResponse struct {
	OriginalConversationID   string `json:"output_OriginalConversationID"`
	ResponseCode             string `json:"output_ResponseCode"`
	ResponseDesc             string `json:"output_ResponseDesc"`
	ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
}

type PushCallbackHandler interface {
	HandlePush(request PushRequest) (PushResponse, error)
}

type PushCallbackFunc func(request PushRequest) (response PushResponse, err error)

func (f PushCallbackFunc) HandlePush(request PushRequest) (PushResponse, error) {
	return f(request)
}

func PushServeHTTP(ctx context.Context, fn PushCallbackHandler, opts ...OptionFunc) http.HandlerFunc {
	o := &options{
		logger: xio.Stderr,
		debug:  false,
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request  PushRequest
			response PushResponse
			err      error
			body     []byte
		)

		body, err = io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal(body, &request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if response, err = fn.HandlePush(request); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
