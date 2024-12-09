package enums

type DebtorInstallmentLineStatus int

const (
	Upcoming DebtorInstallmentLineStatus = iota + 1
	Paid
	Overdue
)

func (e DebtorInstallmentLineStatus) String() string {
	switch e {
	case Upcoming:
		return "Upcoming"
	case Paid:
		return "Paid"
	case Overdue:
		return "Overdue"
	default:
		return "Unknown"
	}
}
