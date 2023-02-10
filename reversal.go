package mpesa

import "context"

type (
	ReverseTxRequest struct {
		Country                  string `json:"input_Country"`                  //nolint:tagliatelle
		ReversalAmount           string `json:"input_ReversalAmount"`           //nolint:tagliatelle
		ServiceProviderCode      string `json:"input_ServiceProviderCode"`      //nolint:tagliatelle
		ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"` //nolint:tagliatelle
		TransactionID            string `json:"input_TransactionID"`            //nolint:tagliatelle
	}

	ReverseTxResponse struct {
		ResponseCode             string `json:"output_ResponseCode"`             //nolint:tagliatelle
		ResponseDesc             string `json:"output_ResponseDesc"`             //nolint:tagliatelle
		TransactionID            string `json:"output_TransactionID"`            //nolint:tagliatelle
		ConversationID           string `json:"output_ConversationID"`           //nolint:tagliatelle
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"` //nolint:tagliatelle
	}

	// reversal The Reversal API is used to reverse a successful transaction.
	// Using the Transaction ID of a previously successful transaction,
	// the OpenAPI will withdraw the funds from the recipient party’s mobile money
	// wallet and revert the funds to the mobile money wallet of the initiating
	// party of the original transaction.
	reversal interface {
		ReverseTx(ctx context.Context, m Mode, request ReverseTxRequest) (ReverseTxResponse, error)
	}

	ReversalFunc func(ctx context.Context, m Mode, request ReverseTxRequest) (ReverseTxResponse, error)
)

func (f ReversalFunc) ReverseTx(ctx context.Context, m Mode, request ReverseTxRequest) (ReverseTxResponse, error) {
	return f(ctx, m, request)
}
