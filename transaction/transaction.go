package transaction

type Transaction struct {
	EntityId            string
	Amount              uint64
	MinorUnits          int
	CurrencyCode        CurrencyCode
	CustomerCountryCode CountryCodeIso31661Alpha2
}
