package enums

type DebtorTransactionStatus int

const (
	Pending DebtorTransactionStatus = iota + 1
	Success
	Failed
	Voided
)

func (e DebtorTransactionStatus) String() string {
	switch e {
	case Pending:
		return "Pending"
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	case Voided:
		return "Voided"
	default:
		return "Unknown"
	}
}
