package mpesa

import (
	"fmt"
	"math"
)

type (
	requestAdapter struct {
		platform            Platform
		market              Market
		serviceProviderCode string
	}
)

func (a *requestAdapter) adapt(requestType requestType, request Request) (interface{}, error) {
	amount := math.Floor(request.Amount * 100 / 100)
	if requestType == pushPay {
		response := pushPayRequest{
			Amount:                   fmt.Sprintf("%0.2f", amount),
			Country:                  a.market.Country(),
			Currency:                 a.market.Currency(),
			CustomerMSISDN:           request.MSISDN,
			ServiceProviderCode:      a.serviceProviderCode,
			ThirdPartyConversationID: request.ThirdPartyID,
			TransactionReference:     request.Reference,
			PurchasedItemsDesc:       request.Description,
		}
		return response, nil
	}

	if requestType == disburse {
		response := disburseRequest{
			Amount:                   fmt.Sprintf("%0.2f", amount),
			Country:                  a.market.Country(),
			Currency:                 a.market.Currency(),
			CustomerMSISDN:           request.MSISDN,
			ServiceProviderCode:      a.serviceProviderCode,
			ThirdPartyConversationID: request.ThirdPartyID,
			TransactionReference:     request.Reference,
			PaymentItemsDesc:         request.Description,
		}

		return response, nil

	}
	return nil, fmt.Errorf("unknown request type: accespted types are pushpay and disburse")
}
