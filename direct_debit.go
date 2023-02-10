package mpesa

import "context"

const (
	ONCE_OFF  DirectDebitFrequency = "01"
	DAILY     DirectDebitFrequency = "02"
	WEEKLY    DirectDebitFrequency = "03"
	MONTHLY   DirectDebitFrequency = "04"
	QUARTER   DirectDebitFrequency = "05"
	HALF_YEAR DirectDebitFrequency = "06"
	YEARLY    DirectDebitFrequency = "07"
	ON_DEMAND DirectDebitFrequency = "08"
)

func (f DirectDebitFrequency) String() string {
	return string(f)
}

// DirectDebitFrequency is the frequency of direct debit
// The below needs to be considered when making use of the DirectDebitCreateRequest.Frequency parameter:
// When Frequency is not used:  FirstPaymentDate ,  StartRangeOfDays , and  EndRangeOfDays
// must be left empty.
//
// When Frequency is set to OneOff, Daily, Weekly or On Demand:  FirstPaymentDate
// (default value: current day) must be set and  StartRangeOfDays  and  EndRangeOfDays
// must be empty.
//
// When Frequency is set to Monthly, Quarterly, HalfYearly, or Yearly:  FirstPaymentDate
// (default value: current day) must be set and  StartRangeOfDays  and  EndRangeOfDays
// are optional.
type DirectDebitFrequency string

// directDebit
// Direct Debits are payments in M-Pesa that are initiated by the Payee alone without any
// Payer interaction, but permission must first be granted by the Payer. The granted permission
// from the Payer to Payee is commonly termed a ‘Mandate’, and M-Pesa must hold details of this Mandate.
//
// The Direct Debit API set allows an organisation to get the initial consent of their customers to create
// the Mandate that allows the organisation to debit customer's account at an agreed frequency and amount'
// for services rendered. After the initial consent, the debit of the account will not involve any customer
// interaction. The Direct Debit feature makes use of the following API calls:
//
// •	Create a Direct Debit mandate
// •	Pay a mandate
//
// The customer is able to view and cancel the Direct Debit mandate from G2 menu accessible via USSD menu
// or the Smartphone Application.
type directDebit interface {
	DirectDebitCreate(ctx context.Context, m Mode, req DirectDebitCreateRequest) (DirectDebitCreateResponse, error)
	DirectDebitPay(ctx context.Context, m Mode, req DirectDebitPayRequest) (DirectDebitPayResponse, error)
}

// DirectDebitCreateRequest is the request body for creating a direct debit
// CustomerMSISDN	The MSISDN of the customer where funds will be debitted from.	True	^[0-9]{12,14}$	254707161122
// Country	The country of the mobile money platform where the transaction needs happen on.	True	N/A	GHA
// ServiceProviderCode	The shortcode of the organization where funds will be creditted to.	True	^([0-9A-Za-z]{4,12})$	ORG001
// ThirdPartyReference	The direct debit's mandate reference	True	^[0-9a-zA-Z]{1,32}$	Test123
// ThirdPartyConversationID	The third party's transaction reference on their system.	True	^[0-9a-zA-Z \w+]{1,40}$	1e9b774d1da34af78412a498cbc28f5e
// AgreedTC	The customer agreed to the terms and conditions. Can only use 1 or 0.	True	^[0-1]{1}$	1
// FirstPaymentDate	The Start date of the Mandate.	False	^[0-9]{8}$	20190205
// Frequency	The frequency of the payments [see table below]	False	^[0-9]{2}$	02
// StartRangeOfDays	The start range of days in the month.	False	^[0-9]{2}$	01
// EndRangeOfDays	The end range of days in the month.	False	^[0-9]{2}$	22
// ExpiryDate	The expiry date of the Mandate.	False	^[0-9]{8}$	20190410
type DirectDebitCreateRequest struct {
	AgreedTC                 string `json:"input_AgreedTC"`
	Country                  string `json:"input_Country"`
	CustomerMSISDN           string `json:"input_CustomerMSISDN"`
	EndRangeOfDays           string `json:"input_EndRangeOfDays"`
	ExpiryDate               string `json:"input_ExpiryDate"`
	FirstPaymentDate         string `json:"input_FirstPaymentDate"`
	Frequency                string `json:"input_Frequency"`
	ServiceProviderCode      string `json:"input_ServiceProviderCode"`
	StartRangeOfDays         string `json:"input_StartRangeOfDays"`
	ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
	ThirdPartyReference      string `json:"input_ThirdPartyReference"`
}

// DirectDebitCreateResponse is the response body for creating a direct debit
// ResponseCode	The result code for the transaction.	INS-0
// ResponseDesc	The result description for the transaction.	Request processed successfully
// TransactionReference	The transaction reference or mandate from the Mobile Money Platform.	vgisfyn4b22w6tmqjftatq75lyuie6vc
// MsisdnToken	The encrypted MSISDN Token built from the provided MSISDN. Only returned in successful messages.	cvgwUBZ3lAO9ivwhWAFeng==
// ConversationID	The OpenAPI platform generates this as a reference to the transaction.	fd1e9143d22544459f7c66e1860ef276
// ThirdPartyConversationID	The incoming reference from the third party system. When there are queries about transactions, this will usually be used to track a transaction.	1e9b774d1da34af78412a498cbc28f5e
type DirectDebitCreateResponse struct {
	ResponseCode             string `json:"output_ResponseCode"`
	ResponseDesc             string `json:"output_ResponseDesc"`
	TransactionReference     string `json:"output_TransactionReference"`
	MsisdnToken              string `json:"output_MsisdnToken"`
	ConversationID           string `json:"output_ConversationID"`
	ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
}

// DirectDebitPayRequest is the request body for paying a direct debit
//  MsisdnToken	The previously returned encrypted MSISDN of the customer where funds will be debitted from. 	False	^[0-9a-zA-Z \w=]{1,26}$	cvgwUBZ3lAO9ivwhWAFeng==
// CustomerMSISDN	The MSISDN of the customer where funds will be debitted from.	False	^[0-9]{12,14}$	254707161122
// Country	The country of the mobile money platform where the transaction needs happen on.	True	N/A	GHA
// ServiceProviderCode	The shortcode of the organization where funds will be creditted to.	True	^([0-9A-Za-z]{4,12})$	ORG001
// ThirdPartyReference	The direct debit's mandate reference	True	^[0-9a-zA-Z]{1,32}$	Test123
// ThirdPartyConversationID	The third party's transaction reference on their system.	True	^[0-9a-zA-Z \w+]{1,40}$	1e9b774d1da34af78412a498cbc28f5e
// Amount	The transaction amount. This amount will be moved from the organization's account to the customer's account.	True	^\d*\.?\d+$	10.00
// Currency	The currency in which the transaction should take place.	True	^[a-zA-Z]{1,3}$	GHS
type DirectDebitPayRequest struct {
	MsisdnToken              string `json:"input_MsisdnToken"`
	Amount                   string `json:"input_Amount"`
	Country                  string `json:"input_Country"`
	Currency                 string `json:"input_Currency"`
	CustomerMSISDN           string `json:"input_CustomerMSISDN"`
	ServiceProviderCode      string `json:"input_ServiceProviderCode"`
	ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
	ThirdPartyReference      string `json:"input_ThirdPartyReference"`
}

// DirectDebitPayResponse is the response body for paying a direct debit
//  ResponseCode	The result code for the transaction.	INS-0
// ResponseDesc	The result description for the transaction.	Request processed successfully
// TransactionID	The transaction identifier that gets generated on the Mobile Money platform. This is used to query transactions on the Mobile Money Platform.	hv9ahxcg4ccv
// MsisdnToken	The encrypted MSISDN Token, which can be used as an identifier. Only returned in successful messages.	cvgwUBZ3lAO9ivwhWAFeng==
// ConversationID	The OpenAPI platform generates this as a reference to the transaction.	fd1e9143d22544459f7c66e1860ef276
// ThirdPartyConversationID	The incoming reference from the third party system. When there are queries about transactions, this will usually be used to track a transaction.	1e9b774d1da34af78412a498cbc28f5e
type DirectDebitPayResponse struct {
	ResponseCode             string `json:"output_ResponseCode"`
	ResponseDesc             string `json:"output_ResponseDesc"`
	TransactionReference     string `json:"output_TransactionReference"`
	MsisdnToken              string `json:"output_MsisdnToken"`
	ConversationID           string `json:"output_ConversationID"`
	ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
}
