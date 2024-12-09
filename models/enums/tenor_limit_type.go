package enums

type TenorLimitType int

const (
	Monthly TenorLimitType = iota + 1
	Yearly
)

func (e TenorLimitType) String() string {
	switch e {
	case Monthly:
		return "Monthly"
	case Yearly:
		return "Yearly"
	default:
		return "Unknown"
	}
}
