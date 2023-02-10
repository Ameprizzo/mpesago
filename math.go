package mpesa

import (
	"math"
)

const percentile = 100

func twoDecimalPlaces(req requestType, amount float64) float64 {
	ceil := func(x float64) float64 {
		return math.Ceil(x*percentile) / percentile
	}

	floor := func(x float64) float64 {
		return math.Floor(x*percentile) / percentile
	}

	round := func(x float64) float64 {
		return math.Round(x*percentile) / percentile
	}

	switch req { //nolint:exhaustive
	case disburse:
		return round(amount)

	case pushPay:
		return ceil(amount)

	default:
		return floor(amount)
	}
}
