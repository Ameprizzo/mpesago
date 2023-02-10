package mpesa

import (
	"fmt"
	"net/http"

	"github.com/techcraftlabs/base"
)

var (
	_ base.RequestInformer = (*requestType)(nil)
)

const (
	sessionID requestType = iota
	pushPay
	disburse
	queryTxn
)

type (
	requestType int
)

func (r requestType) Endpoint() string {
	switch r {

	case sessionID:
		return "/getSession/"

	case pushPay:
		return "/c2bPayment/singleStage/"

	case disburse:
		return "/b2cPayment/"

	case queryTxn:
		return "/queryTransactionStatus/"

	default:
		return ""
	}
}

func (r requestType) Method() string {
	switch r {

	case sessionID:
		return http.MethodGet

	case pushPay:
		return http.MethodPost

	default:
		return http.MethodPost

	}
}

func (r requestType) Name() string {
	return []string{"get session id", "ussd push",
		"disbursement"}[r]
}

func (r requestType) MNO() string {
	return "vodacom"
}

func (r requestType) Group() string {
	switch r {
	case sessionID:
		return "Authorization"

	case pushPay:
		return "collection"

	case disburse:
		return "disbursement"

	default:
		return ""
	}
}

func (r requestType) RecipientMNO() string {
	return ""
}

func (r requestType) String() string {

	return fmt.Sprintf("mno=%s:: group=%s:: request=%s:", r.MNO(), r.Group(), r.Name())
}
