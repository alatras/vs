package transaction

type Transaction struct {
	Amount       uint64
	CurrencyCode CurrencyCode
	CountryCode  string
	EntityId     string
}
