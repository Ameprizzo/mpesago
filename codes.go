package mpesa

import (
	"sync"
)

const SUCCESS_CODE = "INS-0"

// responseCodes are codes returned by the M-Pesa API
var responseCodes = map[string]string{ //nolint:gochecknoglobals
	"INS-0":    "Request processed successfully",
	"INS-1":    "Internal Error",
	"INS-6":    "Transaction Failed",
	"INS-9":    "Request timeout",
	"INS-10":   "Duplicate Transaction",
	"INS-13":   "Invalid Shortcode Used",
	"INS-15":   "Invalid Amount Used",
	"INS-17":   "Invalid Transaction Reference. Length Should Be Between 1 and 20.",
	"INS-18":   "Invalid TransactionID Used",
	"INS-20":   "Not All Parameters Provided. Please try again.",
	"INS-21":   "Parameter validations failed. Please try again.",
	"INS-26":   "Invalid Currency Used",
	"INS-28":   "Invalid ThirdPartyConversationID Used",
	"INS-30":   "Invalid Purchased Items Description Used",
	"INS-35":   "This transaction do not belong to you",
	"INS-36":   "Invalid Agreed TC Used",
	"INS-39":   "Invalid First Payment Date Used",
	"INS-40":   "Invalid Frequency Used",
	"INS-41":   "Invalid Start Range Days Used",
	"INS-42":   "Invalid End Range Days Used",
	"INS-43":   "Invalid Expiry Date Used",
	"INS-50":   "MSISDN Token and MSISDN provided does not Match",
	"INS-990":  "Customer Transaction Value Limit Breached",
	"INS-991":  "Customer Transaction Count Limit Breached",
	"INS-992":  "Multiple Limits Breached",
	"INS-993":  "Organization Transaction Count Limit Breached",
	"INS-994":  "Organization Transaction Value Limit Breached",
	"INS-995":  "API Single Transaction Limit Breached",
	"INS-996":  "API Being Used Outside Of Usage Time",
	"INS-997":  "API Not Enabled",
	"INS-998":  "Invalid Market",
	"INS-2006": "Insufficient balance",
	"INS-2051": "MSISDN invalid.",
}

type codes struct {
	mu       sync.RWMutex
	codesMap map[string]string
}

func (c *codes) get(code string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// check if the code is available and return it
	// if not return "unknown code"
	if val, ok := c.codesMap[code]; ok {
		return val
	}

	return "unknown code"
}

var (
	once           sync.Once //nolint:gochecknoglobals
	singleInstance *codes    //nolint:gochecknoglobals
)

func getInstance() *codes {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &codes{
					mu:       sync.RWMutex{},
					codesMap: responseCodes,
				}
			})
	}

	return singleInstance
}

// ResponseCode returns the description of a response code
func ResponseCode(code string) string {
	return getInstance().get(code)
}
