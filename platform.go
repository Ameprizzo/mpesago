package mpesa

import "strings"

const (
	SANDBOX Platform = iota
	OPENAPI
)

type Platform int

func PlatformFmt(platformString string) Platform {
	if strings.ToLower(platformString) == "openapi" {
		return OPENAPI
	}

	if strings.ToLower(platformString) == "sandbox" {
		return SANDBOX
	}

	return Platform(-1)
}

func (p Platform) String() string {
	if p == OPENAPI {
		return "openapi"
	}

	return "sandbox"
}
