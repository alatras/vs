package transaction

type Transaction struct {
	EntityId            string
	Amount              uint64
	MinorUnits          int
	CurrencyCode        CurrencyCode
	CustomerCountryCode CountryCodeIso31661Alpha2
	Card                string
	IssuerCountryCode   CountryCodeIso31661Alpha2
	CustomerId          string
	CustomerIP          string
	CustomerIPCountry   string
}
