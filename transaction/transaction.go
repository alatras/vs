package transaction

type Transaction struct {
	Amount       uint
	CurrencyCode CurrencyCode
	CountryCode  string
	EntityId     string
}
