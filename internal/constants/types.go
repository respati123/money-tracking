package constants

type TransactionType int

const (
	Debit TransactionType = iota
	Credit
)

var TransactionTypes = struct {
	Debit  TransactionType
	Credit TransactionType
}{
	Debit:  Debit,
	Credit: Credit,
}

func (t TransactionType) String() string {
	switch t {
	case Debit:
		return "debit"
	case Credit:
		return "credit"
	default:
		return "Unknown"
	}
}
