package mpesa

import (
	"context"
	"net/http"
)

type (
	// QueryTxParams is the parameters for querying a transaction
	QueryTxParams struct {
		Reference           string
		ServiceProviderCode string
		ConversationID      string
		CountryCode         string
	}

	// QueryTxResponse is the response from querying a transaction
	QueryTxResponse struct {
		ConversationID            string `json:"output_ConversationID"`
		ResponseCode              string `json:"output_ResponseCode"`
		ResponseDesc              string `json:"output_ResponseDesc"`
		ResponseTransactionStatus string `json:"output_ResponseTransactionStatus"`
		ThirdPartyConversationID  string `json:"output_ThirdPartyConversationID"`
	}

	querier interface {
		QueryTx(ctx context.Context, m Mode, req QueryTxParams) (QueryTxResponse, error)
		QueryCallbackServeHTTP(w http.ResponseWriter, r *http.Request)
	}

	QueryTxFunc       func(ctx context.Context, m Mode, req QueryTxParams) (QueryTxResponse, error)
	QueryCallbackFunc func(ctx context.Context, req QueryTxParams) (QueryTxResponse, error)
)
