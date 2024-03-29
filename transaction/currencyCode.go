package transaction

type CurrencyCode string

func IsCurrencyCode(countryCode string) bool {
	return currencyCodeSet[countryCode]
}

const (
	FJD CurrencyCode = "FJD"
	MXN CurrencyCode = "MXN"
	STD CurrencyCode = "STD"
	LVL CurrencyCode = "LVL"
	SCR CurrencyCode = "SCR"
	CDF CurrencyCode = "CDF"
	BBD CurrencyCode = "BBD"
	GTQ CurrencyCode = "GTQ"
	CLP CurrencyCode = "CLP"
	HNL CurrencyCode = "HNL"
	UGX CurrencyCode = "UGX"
	MXV CurrencyCode = "MXV"
	ZAR CurrencyCode = "ZAR"
	TND CurrencyCode = "TND"
	CUC CurrencyCode = "CUC"
	BSD CurrencyCode = "BSD"
	SLL CurrencyCode = "SLL"
	SDG CurrencyCode = "SDG"
	IQD CurrencyCode = "IQD"
	CUP CurrencyCode = "CUP"
	GMD CurrencyCode = "GMD"
	TWD CurrencyCode = "TWD"
	RSD CurrencyCode = "RSD"
	DOP CurrencyCode = "DOP"
	UYI CurrencyCode = "UYI"
	KMF CurrencyCode = "KMF"
	MYR CurrencyCode = "MYR"
	FKP CurrencyCode = "FKP"
	XOF CurrencyCode = "XOF"
	GEL CurrencyCode = "GEL"
	UYU CurrencyCode = "UYU"
	BTC CurrencyCode = "BTC"
	MAD CurrencyCode = "MAD"
	CVE CurrencyCode = "CVE"
	TOP CurrencyCode = "TOP"
	AZN CurrencyCode = "AZN"
	OMR CurrencyCode = "OMR"
	PGK CurrencyCode = "PGK"
	KES CurrencyCode = "KES"
	SEK CurrencyCode = "SEK"
	BTN CurrencyCode = "BTN"
	UAH CurrencyCode = "UAH"
	GNF CurrencyCode = "GNF"
	ERN CurrencyCode = "ERN"
	MZN CurrencyCode = "MZN"
	SVC CurrencyCode = "SVC"
	ARS CurrencyCode = "ARS"
	QAR CurrencyCode = "QAR"
	IRR CurrencyCode = "IRR"
	MRO CurrencyCode = "MRO"
	XPD CurrencyCode = "XPD"
	CNY CurrencyCode = "CNY"
	THB CurrencyCode = "THB"
	UZS CurrencyCode = "UZS"
	XPF CurrencyCode = "XPF"
	BDT CurrencyCode = "BDT"
	LYD CurrencyCode = "LYD"
	BMD CurrencyCode = "BMD"
	KWD CurrencyCode = "KWD"
	PHP CurrencyCode = "PHP"
	XXX CurrencyCode = "XXX"
	XPT CurrencyCode = "XPT"
	RUB CurrencyCode = "RUB"
	PYG CurrencyCode = "PYG"
	ISK CurrencyCode = "ISK"
	JMD CurrencyCode = "JMD"
	COP CurrencyCode = "COP"
	MKD CurrencyCode = "MKD"
	USD CurrencyCode = "USD"
	COU CurrencyCode = "COU"
	DZD CurrencyCode = "DZD"
	PAB CurrencyCode = "PAB"
	SGD CurrencyCode = "SGD"
	USN CurrencyCode = "USN"
	ETB CurrencyCode = "ETB"
	USS CurrencyCode = "USS"
	KGS CurrencyCode = "KGS"
	SOS CurrencyCode = "SOS"
	VEF CurrencyCode = "VEF"
	VUV CurrencyCode = "VUV"
	LAK CurrencyCode = "LAK"
	BND CurrencyCode = "BND"
	ZMK CurrencyCode = "ZMK"
	XAF CurrencyCode = "XAF"
	LRD CurrencyCode = "LRD"
	XAG CurrencyCode = "XAG"
	CHF CurrencyCode = "CHF"
	HRK CurrencyCode = "HRK"
	ALL CurrencyCode = "ALL"
	CHE CurrencyCode = "CHE"
	DJF CurrencyCode = "DJF"
	ZMW CurrencyCode = "ZMW"
	TZS CurrencyCode = "TZS"
	VND CurrencyCode = "VND"
	XAU CurrencyCode = "XAU"
	AUD CurrencyCode = "AUD"
	ILS CurrencyCode = "ILS"
	CHW CurrencyCode = "CHW"
	GHS CurrencyCode = "GHS"
	GYD CurrencyCode = "GYD"
	KPW CurrencyCode = "KPW"
	BOB CurrencyCode = "BOB"
	KHR CurrencyCode = "KHR"
	MDL CurrencyCode = "MDL"
	IDR CurrencyCode = "IDR"
	XBA CurrencyCode = "XBA"
	KYD CurrencyCode = "KYD"
	AMD CurrencyCode = "AMD"
	XBC CurrencyCode = "XBC"
	XBB CurrencyCode = "XBB"
	BWP CurrencyCode = "BWP"
	SHP CurrencyCode = "SHP"
	TRY CurrencyCode = "TRY"
	LBP CurrencyCode = "LBP"
	XBD CurrencyCode = "XBD"
	TJS CurrencyCode = "TJS"
	JOD CurrencyCode = "JOD"
	AED CurrencyCode = "AED"
	HKD CurrencyCode = "HKD"
	RWF CurrencyCode = "RWF"
	EUR CurrencyCode = "EUR"
	LSL CurrencyCode = "LSL"
	DKK CurrencyCode = "DKK"
	CAD CurrencyCode = "CAD"
	BGN CurrencyCode = "BGN"
	BOV CurrencyCode = "BOV"
	MMK CurrencyCode = "MMK"
	MUR CurrencyCode = "MUR"
	NOK CurrencyCode = "NOK"
	SYP CurrencyCode = "SYP"
	GIP CurrencyCode = "GIP"
	RON CurrencyCode = "RON"
	LKR CurrencyCode = "LKR"
	NGN CurrencyCode = "NGN"
	CRC CurrencyCode = "CRC"
	CZK CurrencyCode = "CZK"
	PKR CurrencyCode = "PKR"
	XCD CurrencyCode = "XCD"
	ANG CurrencyCode = "ANG"
	HTG CurrencyCode = "HTG"
	BHD CurrencyCode = "BHD"
	KZT CurrencyCode = "KZT"
	SRD CurrencyCode = "SRD"
	SZL CurrencyCode = "SZL"
	LTL CurrencyCode = "LTL"
	SAR CurrencyCode = "SAR"
	TTD CurrencyCode = "TTD"
	YER CurrencyCode = "YER"
	MVR CurrencyCode = "MVR"
	AFN CurrencyCode = "AFN"
	INR CurrencyCode = "INR"
	AWG CurrencyCode = "AWG"
	KRW CurrencyCode = "KRW"
	NPR CurrencyCode = "NPR"
	JPY CurrencyCode = "JPY"
	MNT CurrencyCode = "MNT"
	AOA CurrencyCode = "AOA"
	PLN CurrencyCode = "PLN"
	GBP CurrencyCode = "GBP"
	SBD CurrencyCode = "SBD"
	XTS CurrencyCode = "XTS"
	HUF CurrencyCode = "HUF"
	BYR CurrencyCode = "BYR"
	BIF CurrencyCode = "BIF"
	MWK CurrencyCode = "MWK"
	MGA CurrencyCode = "MGA"
	XDR CurrencyCode = "XDR"
	BZD CurrencyCode = "BZD"
	BAM CurrencyCode = "BAM"
	EGP CurrencyCode = "EGP"
	MOP CurrencyCode = "MOP"
	NAD CurrencyCode = "NAD"
	SSP CurrencyCode = "SSP"
	NIO CurrencyCode = "NIO"
	PEN CurrencyCode = "PEN"
	NZD CurrencyCode = "NZD"
	WST CurrencyCode = "WST"
	TMT CurrencyCode = "TMT"
	CLF CurrencyCode = "CLF"
	BRL CurrencyCode = "BRL"
)

var currencyCodeSet = map[string]bool{
	"FJD": true,
	"MXN": true,
	"STD": true,
	"LVL": true,
	"SCR": true,
	"CDF": true,
	"BBD": true,
	"GTQ": true,
	"CLP": true,
	"HNL": true,
	"UGX": true,
	"MXV": true,
	"ZAR": true,
	"TND": true,
	"CUC": true,
	"BSD": true,
	"SLL": true,
	"SDG": true,
	"IQD": true,
	"CUP": true,
	"GMD": true,
	"TWD": true,
	"RSD": true,
	"DOP": true,
	"UYI": true,
	"KMF": true,
	"MYR": true,
	"FKP": true,
	"XOF": true,
	"GEL": true,
	"UYU": true,
	"BTC": true,
	"MAD": true,
	"CVE": true,
	"TOP": true,
	"AZN": true,
	"OMR": true,
	"PGK": true,
	"KES": true,
	"SEK": true,
	"BTN": true,
	"UAH": true,
	"GNF": true,
	"ERN": true,
	"MZN": true,
	"SVC": true,
	"ARS": true,
	"QAR": true,
	"IRR": true,
	"MRO": true,
	"XPD": true,
	"CNY": true,
	"THB": true,
	"UZS": true,
	"XPF": true,
	"BDT": true,
	"LYD": true,
	"BMD": true,
	"KWD": true,
	"PHP": true,
	"XXX": true,
	"XPT": true,
	"RUB": true,
	"PYG": true,
	"ISK": true,
	"JMD": true,
	"COP": true,
	"MKD": true,
	"USD": true,
	"COU": true,
	"DZD": true,
	"PAB": true,
	"SGD": true,
	"USN": true,
	"ETB": true,
	"USS": true,
	"KGS": true,
	"SOS": true,
	"VEF": true,
	"VUV": true,
	"LAK": true,
	"BND": true,
	"ZMK": true,
	"XAF": true,
	"LRD": true,
	"XAG": true,
	"CHF": true,
	"HRK": true,
	"ALL": true,
	"CHE": true,
	"DJF": true,
	"ZMW": true,
	"TZS": true,
	"VND": true,
	"XAU": true,
	"AUD": true,
	"ILS": true,
	"CHW": true,
	"GHS": true,
	"GYD": true,
	"KPW": true,
	"BOB": true,
	"KHR": true,
	"MDL": true,
	"IDR": true,
	"XBA": true,
	"KYD": true,
	"AMD": true,
	"XBC": true,
	"XBB": true,
	"BWP": true,
	"SHP": true,
	"TRY": true,
	"LBP": true,
	"XBD": true,
	"TJS": true,
	"JOD": true,
	"AED": true,
	"HKD": true,
	"RWF": true,
	"EUR": true,
	"LSL": true,
	"DKK": true,
	"CAD": true,
	"BGN": true,
	"BOV": true,
	"MMK": true,
	"MUR": true,
	"NOK": true,
	"SYP": true,
	"GIP": true,
	"RON": true,
	"LKR": true,
	"NGN": true,
	"CRC": true,
	"CZK": true,
	"PKR": true,
	"XCD": true,
	"ANG": true,
	"HTG": true,
	"BHD": true,
	"KZT": true,
	"SRD": true,
	"SZL": true,
	"LTL": true,
	"SAR": true,
	"TTD": true,
	"YER": true,
	"MVR": true,
	"AFN": true,
	"INR": true,
	"AWG": true,
	"KRW": true,
	"NPR": true,
	"JPY": true,
	"MNT": true,
	"AOA": true,
	"PLN": true,
	"GBP": true,
	"SBD": true,
	"XTS": true,
	"HUF": true,
	"BYR": true,
	"BIF": true,
	"MWK": true,
	"MGA": true,
	"XDR": true,
	"BZD": true,
	"BAM": true,
	"EGP": true,
	"MOP": true,
	"NAD": true,
	"SSP": true,
	"NIO": true,
	"PEN": true,
	"NZD": true,
	"WST": true,
	"TMT": true,
	"CLF": true,
	"BRL": true,
}
