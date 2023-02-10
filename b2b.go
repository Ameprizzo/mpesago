package mpesa

import "context"

type B2BRequest struct {
	Amount                   string `json:"input_Amount"`
	Country                  string `json:"input_Country"`
	Currency                 string `json:"input_Currency"`
	PrimaryPartyCode         string `json:"input_PrimaryPartyCode"`
	ReceiverPartyCode        string `json:"input_ReceiverPartyCode"`
	ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
	TransactionReference     string `json:"input_TransactionReference"`
	PurchasedItemsDesc       string `json:"input_PurchasedItemsDesc"`
}

type B2BResponse struct {
	ConversationID           string `json:"output_ConversationID"`
	ResponseCode             string `json:"output_ResponseCode"`
	ResponseDesc             string `json:"output_ResponseDesc"`
	TransactionID            string `json:"output_TransactionID"`
	ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
}

// b2b The B2B API Call is used for business-to-business transactions. Funds from
// the business’ mobile money wallet will be deducted and transferred to the mobile
// money wallet of the other business. Use cases for the B2C includes:
// • Stock purchases
// • Bill payment
// • Ad-hoc payment
// • Business to business transactions
type b2b interface {
	B2BPush(ctx context.Context, m Mode, b2bReq B2BRequest) (B2BResponse, error)
}

type B2BPushFunc func(ctx context.Context, m Mode, b2bReq B2BRequest) (B2BResponse, error)
