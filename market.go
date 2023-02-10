package mpesa

import (
	"strings"
)

var _ market = (*Market)(nil)

const (
	GhanaMarket    = Market(0)
	TanzaniaMarket = Market(1)
)

type market interface {
	URLContextValue() string
	Country() string
	Currency() string
	Description() string
}

type Market int

func MarketFmt(marketString string) Market {
	if strings.ToLower(marketString) == "ghana" {
		return GhanaMarket
	}

	if strings.ToLower(marketString) == "tanzania" {
		return TanzaniaMarket
	}

	return Market(-1)
}

func (m Market) URLContextValue() string {
	switch m {

	//ghana
	case 0:
		return "vodafoneGHA"
		//tanzania
	case 1:
		return "vodacomTZN"
	default:
		return ""
	}
}

func (m Market) Country() string {
	switch m {

	//ghana
	case 0:
		return "GHA"
		//tanzania
	case 1:
		return "TZN"
	default:
		return ""
	}
}

func (m Market) Currency() string {
	switch m {

	//ghana
	case 0:
		return "GHS"
		//tanzania
	case 1:
		return "TZS"
	default:
		return ""
	}
}

func (m Market) Description() string {
	switch m {

	//ghana
	case 0:
		return "Vodafone Ghana"
		//tanzania
	case 1:
		return "Vodacom Tanzania"
	default:
		return ""
	}
}
