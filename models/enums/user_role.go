package enums

type UserRole int

const (
	Admin UserRole = iota + 1
	Debtor
)

func (e UserRole) String() string {
	switch e {
	case Admin:
		return "admin"
	case Debtor:
		return "debtor"
	default:
		return "unknown"
	}
}
