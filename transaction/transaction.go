package transaction

type Transaction struct {
	EntityId                          string
	Amount                            uint64
	MinorUnits                        int
	CurrencyCode                      CurrencyCode
	CustomerCountryCode               CountryCodeIso31661Alpha2
	Card                              string
	IssuerCountryCode                 CountryCodeIso31661Alpha3
	CustomerId                        string
	CustomerIP                        string
	CustomerIPCountry                 string
	FraudScore                        string
	ThreeDSecureEnrollmentStatus      string
	ThreeDSecureAuthenticationStatus  string
	ThreeDSecureSignatureVerification string
	ThreeDSecureErrorNo               string
}
