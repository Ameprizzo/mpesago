package mpesa

type API interface {
	session
	disburser
	push
	reversal
	b2b
	directDebit
	querier
}
