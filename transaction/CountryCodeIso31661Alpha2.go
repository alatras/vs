package transaction

import "fmt"

type CountryCodeIso31661Alpha2 string

func IsCountryCodeIso31661Alpha2(countryCode string) bool {
	res := countryCodeIso31661Alpha2Set[countryCode]
	fmt.Println("res")
	return !res
}

const (
	PR CountryCodeIso31661Alpha2 = "PR"
	PS CountryCodeIso31661Alpha2 = "PS"
	PT CountryCodeIso31661Alpha2 = "PT"
	PW CountryCodeIso31661Alpha2 = "PW"
	PY CountryCodeIso31661Alpha2 = "PY"
	QA CountryCodeIso31661Alpha2 = "QA"
	AD CountryCodeIso31661Alpha2 = "AD"
	AE CountryCodeIso31661Alpha2 = "AE"
	AF CountryCodeIso31661Alpha2 = "AF"
	AG CountryCodeIso31661Alpha2 = "AG"
	AI CountryCodeIso31661Alpha2 = "AI"
	AL CountryCodeIso31661Alpha2 = "AL"
	AM CountryCodeIso31661Alpha2 = "AM"
	AO CountryCodeIso31661Alpha2 = "AO"
	AQ CountryCodeIso31661Alpha2 = "AQ"
	AR CountryCodeIso31661Alpha2 = "AR"
	AS CountryCodeIso31661Alpha2 = "AS"
	AT CountryCodeIso31661Alpha2 = "AT"
	RE CountryCodeIso31661Alpha2 = "RE"
	AU CountryCodeIso31661Alpha2 = "AU"
	AW CountryCodeIso31661Alpha2 = "AW"
	AX CountryCodeIso31661Alpha2 = "AX"
	AZ CountryCodeIso31661Alpha2 = "AZ"
	RO CountryCodeIso31661Alpha2 = "RO"
	BA CountryCodeIso31661Alpha2 = "BA"
	BB CountryCodeIso31661Alpha2 = "BB"
	RS CountryCodeIso31661Alpha2 = "RS"
	BD CountryCodeIso31661Alpha2 = "BD"
	BE CountryCodeIso31661Alpha2 = "BE"
	RU CountryCodeIso31661Alpha2 = "RU"
	BF CountryCodeIso31661Alpha2 = "BF"
	BG CountryCodeIso31661Alpha2 = "BG"
	RW CountryCodeIso31661Alpha2 = "RW"
	BH CountryCodeIso31661Alpha2 = "BH"
	BI CountryCodeIso31661Alpha2 = "BI"
	BJ CountryCodeIso31661Alpha2 = "BJ"
	BL CountryCodeIso31661Alpha2 = "BL"
	BM CountryCodeIso31661Alpha2 = "BM"
	BN CountryCodeIso31661Alpha2 = "BN"
	BO CountryCodeIso31661Alpha2 = "BO"
	SA CountryCodeIso31661Alpha2 = "SA"
	BQ CountryCodeIso31661Alpha2 = "BQ"
	SB CountryCodeIso31661Alpha2 = "SB"
	BR CountryCodeIso31661Alpha2 = "BR"
	SC CountryCodeIso31661Alpha2 = "SC"
	BS CountryCodeIso31661Alpha2 = "BS"
	SD CountryCodeIso31661Alpha2 = "SD"
	BT CountryCodeIso31661Alpha2 = "BT"
	SE CountryCodeIso31661Alpha2 = "SE"
	BV CountryCodeIso31661Alpha2 = "BV"
	SG CountryCodeIso31661Alpha2 = "SG"
	BW CountryCodeIso31661Alpha2 = "BW"
	SH CountryCodeIso31661Alpha2 = "SH"
	SI CountryCodeIso31661Alpha2 = "SI"
	BY CountryCodeIso31661Alpha2 = "BY"
	SJ CountryCodeIso31661Alpha2 = "SJ"
	BZ CountryCodeIso31661Alpha2 = "BZ"
	SK CountryCodeIso31661Alpha2 = "SK"
	SL CountryCodeIso31661Alpha2 = "SL"
	SM CountryCodeIso31661Alpha2 = "SM"
	SN CountryCodeIso31661Alpha2 = "SN"
	SO CountryCodeIso31661Alpha2 = "SO"
	CA CountryCodeIso31661Alpha2 = "CA"
	SR CountryCodeIso31661Alpha2 = "SR"
	CC CountryCodeIso31661Alpha2 = "CC"
	SS CountryCodeIso31661Alpha2 = "SS"
	CD CountryCodeIso31661Alpha2 = "CD"
	ST CountryCodeIso31661Alpha2 = "ST"
	CF CountryCodeIso31661Alpha2 = "CF"
	SV CountryCodeIso31661Alpha2 = "SV"
	CG CountryCodeIso31661Alpha2 = "CG"
	CH CountryCodeIso31661Alpha2 = "CH"
	SX CountryCodeIso31661Alpha2 = "SX"
	CI CountryCodeIso31661Alpha2 = "CI"
	SY CountryCodeIso31661Alpha2 = "SY"
	SZ CountryCodeIso31661Alpha2 = "SZ"
	CK CountryCodeIso31661Alpha2 = "CK"
	CL CountryCodeIso31661Alpha2 = "CL"
	CM CountryCodeIso31661Alpha2 = "CM"
	CN CountryCodeIso31661Alpha2 = "CN"
	CO CountryCodeIso31661Alpha2 = "CO"
	CR CountryCodeIso31661Alpha2 = "CR"
	TC CountryCodeIso31661Alpha2 = "TC"
	TD CountryCodeIso31661Alpha2 = "TD"
	CU CountryCodeIso31661Alpha2 = "CU"
	TF CountryCodeIso31661Alpha2 = "TF"
	CV CountryCodeIso31661Alpha2 = "CV"
	TG CountryCodeIso31661Alpha2 = "TG"
	CW CountryCodeIso31661Alpha2 = "CW"
	TH CountryCodeIso31661Alpha2 = "TH"
	CX CountryCodeIso31661Alpha2 = "CX"
	CY CountryCodeIso31661Alpha2 = "CY"
	TJ CountryCodeIso31661Alpha2 = "TJ"
	CZ CountryCodeIso31661Alpha2 = "CZ"
	TK CountryCodeIso31661Alpha2 = "TK"
	TL CountryCodeIso31661Alpha2 = "TL"
	TM CountryCodeIso31661Alpha2 = "TM"
	TN CountryCodeIso31661Alpha2 = "TN"
	TO CountryCodeIso31661Alpha2 = "TO"
	TR CountryCodeIso31661Alpha2 = "TR"
	TT CountryCodeIso31661Alpha2 = "TT"
	DE CountryCodeIso31661Alpha2 = "DE"
	TV CountryCodeIso31661Alpha2 = "TV"
	TW CountryCodeIso31661Alpha2 = "TW"
	DJ CountryCodeIso31661Alpha2 = "DJ"
	TZ CountryCodeIso31661Alpha2 = "TZ"
	DK CountryCodeIso31661Alpha2 = "DK"
	DM CountryCodeIso31661Alpha2 = "DM"
	DO CountryCodeIso31661Alpha2 = "DO"
	UA CountryCodeIso31661Alpha2 = "UA"
	UG CountryCodeIso31661Alpha2 = "UG"
	DZ CountryCodeIso31661Alpha2 = "DZ"
	UM CountryCodeIso31661Alpha2 = "UM"
	EC CountryCodeIso31661Alpha2 = "EC"
	US CountryCodeIso31661Alpha2 = "US"
	EE CountryCodeIso31661Alpha2 = "EE"
	EG CountryCodeIso31661Alpha2 = "EG"
	EH CountryCodeIso31661Alpha2 = "EH"
	UY CountryCodeIso31661Alpha2 = "UY"
	UZ CountryCodeIso31661Alpha2 = "UZ"
	VA CountryCodeIso31661Alpha2 = "VA"
	ER CountryCodeIso31661Alpha2 = "ER"
	VC CountryCodeIso31661Alpha2 = "VC"
	ES CountryCodeIso31661Alpha2 = "ES"
	ET CountryCodeIso31661Alpha2 = "ET"
	VE CountryCodeIso31661Alpha2 = "VE"
	VG CountryCodeIso31661Alpha2 = "VG"
	VI CountryCodeIso31661Alpha2 = "VI"
	VN CountryCodeIso31661Alpha2 = "VN"
	VU CountryCodeIso31661Alpha2 = "VU"
	FI CountryCodeIso31661Alpha2 = "FI"
	FJ CountryCodeIso31661Alpha2 = "FJ"
	FK CountryCodeIso31661Alpha2 = "FK"
	FM CountryCodeIso31661Alpha2 = "FM"
	FO CountryCodeIso31661Alpha2 = "FO"
	FR CountryCodeIso31661Alpha2 = "FR"
	WF CountryCodeIso31661Alpha2 = "WF"
	GA CountryCodeIso31661Alpha2 = "GA"
	GB CountryCodeIso31661Alpha2 = "GB"
	WS CountryCodeIso31661Alpha2 = "WS"
	GD CountryCodeIso31661Alpha2 = "GD"
	GE CountryCodeIso31661Alpha2 = "GE"
	GF CountryCodeIso31661Alpha2 = "GF"
	GG CountryCodeIso31661Alpha2 = "GG"
	GH CountryCodeIso31661Alpha2 = "GH"
	GI CountryCodeIso31661Alpha2 = "GI"
	GL CountryCodeIso31661Alpha2 = "GL"
	GM CountryCodeIso31661Alpha2 = "GM"
	GN CountryCodeIso31661Alpha2 = "GN"
	GP CountryCodeIso31661Alpha2 = "GP"
	GQ CountryCodeIso31661Alpha2 = "GQ"
	GR CountryCodeIso31661Alpha2 = "GR"
	GS CountryCodeIso31661Alpha2 = "GS"
	GT CountryCodeIso31661Alpha2 = "GT"
	GU CountryCodeIso31661Alpha2 = "GU"
	GW CountryCodeIso31661Alpha2 = "GW"
	GY CountryCodeIso31661Alpha2 = "GY"
	HK CountryCodeIso31661Alpha2 = "HK"
	HM CountryCodeIso31661Alpha2 = "HM"
	HN CountryCodeIso31661Alpha2 = "HN"
	HR CountryCodeIso31661Alpha2 = "HR"
	HT CountryCodeIso31661Alpha2 = "HT"
	YE CountryCodeIso31661Alpha2 = "YE"
	HU CountryCodeIso31661Alpha2 = "HU"
	ID CountryCodeIso31661Alpha2 = "ID"
	YT CountryCodeIso31661Alpha2 = "YT"
	IE CountryCodeIso31661Alpha2 = "IE"
	IL CountryCodeIso31661Alpha2 = "IL"
	IM CountryCodeIso31661Alpha2 = "IM"
	IN CountryCodeIso31661Alpha2 = "IN"
	IO CountryCodeIso31661Alpha2 = "IO"
	ZA CountryCodeIso31661Alpha2 = "ZA"
	IQ CountryCodeIso31661Alpha2 = "IQ"
	IR CountryCodeIso31661Alpha2 = "IR"
	IS CountryCodeIso31661Alpha2 = "IS"
	IT CountryCodeIso31661Alpha2 = "IT"
	ZM CountryCodeIso31661Alpha2 = "ZM"
	JE CountryCodeIso31661Alpha2 = "JE"
	ZW CountryCodeIso31661Alpha2 = "ZW"
	ZZ CountryCodeIso31661Alpha2 = "ZZ"
	JM CountryCodeIso31661Alpha2 = "JM"
	JO CountryCodeIso31661Alpha2 = "JO"
	JP CountryCodeIso31661Alpha2 = "JP"
	KE CountryCodeIso31661Alpha2 = "KE"
	KG CountryCodeIso31661Alpha2 = "KG"
	KH CountryCodeIso31661Alpha2 = "KH"
	KI CountryCodeIso31661Alpha2 = "KI"
	KM CountryCodeIso31661Alpha2 = "KM"
	KN CountryCodeIso31661Alpha2 = "KN"
	KP CountryCodeIso31661Alpha2 = "KP"
	KR CountryCodeIso31661Alpha2 = "KR"
	KW CountryCodeIso31661Alpha2 = "KW"
	KY CountryCodeIso31661Alpha2 = "KY"
	KZ CountryCodeIso31661Alpha2 = "KZ"
	LA CountryCodeIso31661Alpha2 = "LA"
	LB CountryCodeIso31661Alpha2 = "LB"
	LC CountryCodeIso31661Alpha2 = "LC"
	LI CountryCodeIso31661Alpha2 = "LI"
	LK CountryCodeIso31661Alpha2 = "LK"
	LR CountryCodeIso31661Alpha2 = "LR"
	LS CountryCodeIso31661Alpha2 = "LS"
	LT CountryCodeIso31661Alpha2 = "LT"
	LU CountryCodeIso31661Alpha2 = "LU"
	LV CountryCodeIso31661Alpha2 = "LV"
	LY CountryCodeIso31661Alpha2 = "LY"
	MA CountryCodeIso31661Alpha2 = "MA"
	MC CountryCodeIso31661Alpha2 = "MC"
	MD CountryCodeIso31661Alpha2 = "MD"
	ME CountryCodeIso31661Alpha2 = "ME"
	MF CountryCodeIso31661Alpha2 = "MF"
	MG CountryCodeIso31661Alpha2 = "MG"
	MH CountryCodeIso31661Alpha2 = "MH"
	MK CountryCodeIso31661Alpha2 = "MK"
	ML CountryCodeIso31661Alpha2 = "ML"
	MM CountryCodeIso31661Alpha2 = "MM"
	MN CountryCodeIso31661Alpha2 = "MN"
	MO CountryCodeIso31661Alpha2 = "MO"
	MP CountryCodeIso31661Alpha2 = "MP"
	MQ CountryCodeIso31661Alpha2 = "MQ"
	MR CountryCodeIso31661Alpha2 = "MR"
	MS CountryCodeIso31661Alpha2 = "MS"
	MT CountryCodeIso31661Alpha2 = "MT"
	MU CountryCodeIso31661Alpha2 = "MU"
	MV CountryCodeIso31661Alpha2 = "MV"
	MW CountryCodeIso31661Alpha2 = "MW"
	MX CountryCodeIso31661Alpha2 = "MX"
	MY CountryCodeIso31661Alpha2 = "MY"
	MZ CountryCodeIso31661Alpha2 = "MZ"
	NA CountryCodeIso31661Alpha2 = "NA"
	NC CountryCodeIso31661Alpha2 = "NC"
	NE CountryCodeIso31661Alpha2 = "NE"
	NF CountryCodeIso31661Alpha2 = "NF"
	NG CountryCodeIso31661Alpha2 = "NG"
	NI CountryCodeIso31661Alpha2 = "NI"
	NL CountryCodeIso31661Alpha2 = "NL"
	NO CountryCodeIso31661Alpha2 = "NO"
	NP CountryCodeIso31661Alpha2 = "NP"
	NR CountryCodeIso31661Alpha2 = "NR"
	NU CountryCodeIso31661Alpha2 = "NU"
	NZ CountryCodeIso31661Alpha2 = "NZ"
	OM CountryCodeIso31661Alpha2 = "OM"
	PA CountryCodeIso31661Alpha2 = "PA"
	PE CountryCodeIso31661Alpha2 = "PE"
	PF CountryCodeIso31661Alpha2 = "PF"
	PG CountryCodeIso31661Alpha2 = "PG"
	PH CountryCodeIso31661Alpha2 = "PH"
	PK CountryCodeIso31661Alpha2 = "PK"
	PL CountryCodeIso31661Alpha2 = "PL"
	PM CountryCodeIso31661Alpha2 = "PM"
	PN CountryCodeIso31661Alpha2 = "PN"
)

var countryCodeIso31661Alpha2Set = map[string]bool{
	"PR": true,
	"PS": true,
	"PT": true,
	"PW": true,
	"PY": true,
	"QA": true,
	"AD": true,
	"AE": true,
	"AF": true,
	"AG": true,
	"AI": true,
	"AL": true,
	"AM": true,
	"AO": true,
	"AQ": true,
	"AR": true,
	"AS": true,
	"AT": true,
	"RE": true,
	"AU": true,
	"AW": true,
	"AX": true,
	"AZ": true,
	"RO": true,
	"BA": true,
	"BB": true,
	"RS": true,
	"BD": true,
	"BE": true,
	"RU": true,
	"BF": true,
	"BG": true,
	"RW": true,
	"BH": true,
	"BI": true,
	"BJ": true,
	"BL": true,
	"BM": true,
	"BN": true,
	"BO": true,
	"SA": true,
	"BQ": true,
	"SB": true,
	"BR": true,
	"SC": true,
	"BS": true,
	"SD": true,
	"BT": true,
	"SE": true,
	"BV": true,
	"SG": true,
	"BW": true,
	"SH": true,
	"SI": true,
	"BY": true,
	"SJ": true,
	"BZ": true,
	"SK": true,
	"SL": true,
	"SM": true,
	"SN": true,
	"SO": true,
	"CA": true,
	"SR": true,
	"CC": true,
	"SS": true,
	"CD": true,
	"ST": true,
	"CF": true,
	"SV": true,
	"CG": true,
	"CH": true,
	"SX": true,
	"CI": true,
	"SY": true,
	"SZ": true,
	"CK": true,
	"CL": true,
	"CM": true,
	"CN": true,
	"CO": true,
	"CR": true,
	"TC": true,
	"TD": true,
	"CU": true,
	"TF": true,
	"CV": true,
	"TG": true,
	"CW": true,
	"TH": true,
	"CX": true,
	"CY": true,
	"TJ": true,
	"CZ": true,
	"TK": true,
	"TL": true,
	"TM": true,
	"TN": true,
	"TO": true,
	"TR": true,
	"TT": true,
	"DE": true,
	"TV": true,
	"TW": true,
	"DJ": true,
	"TZ": true,
	"DK": true,
	"DM": true,
	"DO": true,
	"UA": true,
	"UG": true,
	"DZ": true,
	"UM": true,
	"EC": true,
	"US": true,
	"EE": true,
	"EG": true,
	"EH": true,
	"UY": true,
	"UZ": true,
	"VA": true,
	"ER": true,
	"VC": true,
	"ES": true,
	"ET": true,
	"VE": true,
	"VG": true,
	"VI": true,
	"VN": true,
	"VU": true,
	"FI": true,
	"FJ": true,
	"FK": true,
	"FM": true,
	"FO": true,
	"FR": true,
	"WF": true,
	"GA": true,
	"GB": true,
	"WS": true,
	"GD": true,
	"GE": true,
	"GF": true,
	"GG": true,
	"GH": true,
	"GI": true,
	"GL": true,
	"GM": true,
	"GN": true,
	"GP": true,
	"GQ": true,
	"GR": true,
	"GS": true,
	"GT": true,
	"GU": true,
	"GW": true,
	"GY": true,
	"HK": true,
	"HM": true,
	"HN": true,
	"HR": true,
	"HT": true,
	"YE": true,
	"HU": true,
	"ID": true,
	"YT": true,
	"IE": true,
	"IL": true,
	"IM": true,
	"IN": true,
	"IO": true,
	"ZA": true,
	"IQ": true,
	"IR": true,
	"IS": true,
	"IT": true,
	"ZM": true,
	"JE": true,
	"ZW": true,
	"ZZ": true,
	"JM": true,
	"JO": true,
	"JP": true,
	"KE": true,
	"KG": true,
	"KH": true,
	"KI": true,
	"KM": true,
	"KN": true,
	"KP": true,
	"KR": true,
	"KW": true,
	"KY": true,
	"KZ": true,
	"LA": true,
	"LB": true,
	"LC": true,
	"LI": true,
	"LK": true,
	"LR": true,
	"LS": true,
	"LT": true,
	"LU": true,
	"LV": true,
	"LY": true,
	"MA": true,
	"MC": true,
	"MD": true,
	"ME": true,
	"MF": true,
	"MG": true,
	"MH": true,
	"MK": true,
	"ML": true,
	"MM": true,
	"MN": true,
	"MO": true,
	"MP": true,
	"MQ": true,
	"MR": true,
	"MS": true,
	"MT": true,
	"MU": true,
	"MV": true,
	"MW": true,
	"MX": true,
	"MY": true,
	"MZ": true,
	"NA": true,
	"NC": true,
	"NE": true,
	"NF": true,
	"NG": true,
	"NI": true,
	"NL": true,
	"NO": true,
	"NP": true,
	"NR": true,
	"NU": true,
	"NZ": true,
	"OM": true,
	"PA": true,
	"PE": true,
	"PF": true,
	"PG": true,
	"PH": true,
	"PK": true,
	"PL": true,
	"PM": true,
	"PN": true,
}
