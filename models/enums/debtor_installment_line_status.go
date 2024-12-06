package enums

type DebtorInstallmentLineStatus int

const (
	Upcoming DebtorInstallmentLineStatus = iota + 1
	Paid
	Overdue
)
