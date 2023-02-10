package api

const (
	C2B Category = iota
	B2C
	B2B
	DIRECT_DEBIT_CREATE
	DIRECT_DEBIT_PAY
	REVERSAL
	TX_STATUS
)

type Category int
