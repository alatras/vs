package transaction

type Transaction struct {
	EntityId            string
	Amount              uint64
	MinorUnits          uint64
	CurrencyCode        CurrencyCode
	CustomerCountryCode CountryCodeIso31661Alpha2
}
