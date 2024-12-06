package enums

type DebtorTransactionStatus int

const (
	Pending DebtorTransactionStatus = iota + 1
	Success
	Failed
	Voided
)
