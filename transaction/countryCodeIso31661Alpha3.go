package transaction

type CountryCodeIso31661Alpha3 string

func IsCountryCodeIso31661Alpha3(countryCode string) bool {
	return countryCodeIso31661Alpha3Set[countryCode]
}

const (
	NZL = "NZL"
	FJI = "FJI"
	PNG = "PNG"
	GLP = "GLP"
	STP = "STP"
	BSB = "BSB"
	MHL = "MHL"
	WLF = "WLF"
	CUB = "CUB"
	SDN = "SDN"
	GMB = "GMB"
	CUW = "CUW"
	MYS = "MYS"
	MYT = "MYT"
	TWN = "TWN"
	POL = "POL"
	OMN = "OMN"
	SUR = "SUR"
	ARE = "ARE"
	KEN = "KEN"
	ARG = "ARG"
	GNB = "GNB"
	ARM = "ARM"
	UZB = "UZB"
	SEN = "SEN"
	TGO = "TGO"
	IRL = "IRL"
	FLK = "FLK"
	IRN = "IRN"
	QAT = "QAT"
	BDI = "BDI"
	NLD = "NLD"
	IRQ = "IRQ"
	SVK = "SVK"
	SVN = "SVN"
	GNQ = "GNQ"
	THA = "THA"
	ABW = "ABW"
	ASM = "ASM"
	SWE = "SWE"
	ISL = "ISL"
	BEL = "BEL"
	ISR = "ISR"
	KWT = "KWT"
	LIE = "LIE"
	BEN = "BEN"
	DZA = "DZA"
	ATA = "ATA"
	BES = "BES"
	RUS = "RUS"
	ATF = "ATF"
	ATG = "ATG"
	ITA = "ITA"
	SWZ = "SWZ"
	TZA = "TZA"
	PAK = "PAK"
	BFA = "BFA"
	CXR = "CXR"
	PAN = "PAN"
	SGP = "SGP"
	UKR = "UKR"
	JEY = "JEY"
	KGZ = "KGZ"
	BVT = "BVT"
	DJI = "DJI"
	REU = "REU"
	CHL = "CHL"
	PRI = "PRI"
	CHN = "CHN"
	PRK = "PRK"
	SXM = "SXM"
	MLI = "MLI"
	BWA = "BWA"
	HRV = "HRV"
	KHM = "KHM"
	IDN = "IDN"
	PRT = "PRT"
	MLT = "MLT"
	TJK = "TJK"
	VNM = "VNM"
	CYM = "CYM"
	PRY = "PRY"
	SHN = "SHN"
	CYP = "CYP"
	SYC = "SYC"
	RWA = "RWA"
	BGD = "BGD"
	AUS = "AUS"
	AUT = "AUT"
	LKA = "LKA"
	PSE = "PSE"
	GAB = "GAB"
	ZWE = "ZWE"
	BGR = "BGR"
	SYR = "SYR"
	CZE = "CZE"
	NOR = "NOR"
	UMI = "UMI"
	CIV = "CIV"
	MMR = "MMR"
	TKL = "TKL"
	KIR = "KIR"
	TKM = "TKM"
	GRD = "GRD"
	GRC = "GRC"
	PCN = "PCN"
	HTI = "HTI"
	GRL = "GRL"
	YEM = "YEM"
	AFG = "AFG"
	MNE = "MNE"
	MNG = "MNG"
	NPL = "NPL"
	BHS = "BHS"
	BHR = "BHR"
	MNP = "MNP"
	GBR = "GBR"
	SJM = "SJM"
	DMA = "DMA"
	TLS = "TLS"
	BIH = "BIH"
	HUN = "HUN"
	AGO = "AGO"
	WSM = "WSM"
	FRA = "FRA"
	MOZ = "MOZ"
	NAM = "NAM"
	PER = "PER"
	DNK = "DNK"
	GTM = "GTM"
	FRO = "FRO"
	SLB = "SLB"
	VAT = "VAT"
	SLE = "SLE"
	NRU = "NRU"
	AIA = "AIA"
	GUF = "GUF"
	ZZZ = "ZZZ"
	SLV = "SLV"
	GUM = "GUM"
	FSM = "FSM"
	DOM = "DOM"
	CMR = "CMR"
	GUY = "GUY"
	AZE = "AZE"
	MAC = "MAC"
	GEO = "GEO"
	TON = "TON"
	MAF = "MAF"
	NCL = "NCL"
	SMR = "SMR"
	ERI = "ERI"
	KNA = "KNA"
	MAR = "MAR"
	BLM = "BLM"
	VCT = "VCT"
	BLR = "BLR"
	MRT = "MRT"
	BLZ = "BLZ"
	PHL = "PHL"
	COD = "COD"
	COG = "COG"
	ESH = "ESH"
	PYF = "PYF"
	URY = "URY"
	COK = "COK"
	COM = "COM"
	COL = "COL"
	USA = "USA"
	ESP = "ESP"
	EST = "EST"
	BMU = "BMU"
	MSR = "MSR"
	ZMB = "ZMB"
	KOR = "KOR"
	SOM = "SOM"
	VUT = "VUT"
	ALA = "ALA"
	ECU = "ECU"
	ALB = "ALB"
	ETH = "ETH"
	GGY = "GGY"
	MCO = "MCO"
	NER = "NER"
	LAO = "LAO"
	VEN = "VEN"
	GHA = "GHA"
	CPV = "CPV"
	MDA = "MDA"
	MTQ = "MTQ"
	MDG = "MDG"
	SPM = "SPM"
	NFK = "NFK"
	LBN = "LBN"
	LBR = "LBR"
	BOL = "BOL"
	MDV = "MDV"
	GIB = "GIB"
	LBY = "LBY"
	HKG = "HKG"
	CAF = "CAF"
	LSO = "LSO"
	NGA = "NGA"
	MUS = "MUS"
	IMN = "IMN"
	LCA = "LCA"
	JOR = "JOR"
	GIN = "GIN"
	VGB = "VGB"
	CAN = "CAN"
	TCA = "TCA"
	TCD = "TCD"
	AND = "AND"
	ROU = "ROU"
	CRI = "CRI"
	IND = "IND"
	MEX = "MEX"
	SRB = "SRB"
	KAZ = "KAZ"
	SAU = "SAU"
	JPN = "JPN"
	LTU = "LTU"
	TTO = "TTO"
	PLW = "PLW"
	HMD = "HMD"
	MWI = "MWI"
	SSD = "SSD"
	NIC = "NIC"
	CCK = "CCK"
	FIN = "FIN"
	TUN = "TUN"
	LUX = "LUX"
	UGA = "UGA"
	IOT = "IOT"
	BRA = "BRA"
	TUR = "TUR"
	TUV = "TUV"
	DEU = "DEU"
	EGY = "EGY"
	LVA = "LVA"
	JAM = "JAM"
	NIU = "NIU"
	VIR = "VIR"
	ZAF = "ZAF"
	BRN = "BRN"
	HND = "HND"
)

var countryCodeIso31661Alpha3Set = map[string]bool{
	"NZL": true,
	"FJI": true,
	"PNG": true,
	"GLP": true,
	"STP": true,
	"BSB": true,
	"MHL": true,
	"WLF": true,
	"CUB": true,
	"SDN": true,
	"GMB": true,
	"CUW": true,
	"MYS": true,
	"MYT": true,
	"TWN": true,
	"POL": true,
	"OMN": true,
	"SUR": true,
	"ARE": true,
	"KEN": true,
	"ARG": true,
	"GNB": true,
	"ARM": true,
	"UZB": true,
	"BTN": true,
	"SEN": true,
	"TGO": true,
	"IRL": true,
	"FLK": true,
	"IRN": true,
	"QAT": true,
	"BDI": true,
	"NLD": true,
	"IRQ": true,
	"SVK": true,
	"SVN": true,
	"GNQ": true,
	"THA": true,
	"ABW": true,
	"ASM": true,
	"SWE": true,
	"ISL": true,
	"MKD": true,
	"BEL": true,
	"ISR": true,
	"KWT": true,
	"LIE": true,
	"BEN": true,
	"DZA": true,
	"ATA": true,
	"BES": true,
	"RUS": true,
	"ATF": true,
	"ATG": true,
	"ITA": true,
	"SWZ": true,
	"TZA": true,
	"PAK": true,
	"BFA": true,
	"CXR": true,
	"PAN": true,
	"SGP": true,
	"UKR": true,
	"JEY": true,
	"KGZ": true,
	"BVT": true,
	"CHE": true,
	"DJI": true,
	"REU": true,
	"CHL": true,
	"PRI": true,
	"CHN": true,
	"PRK": true,
	"SXM": true,
	"MLI": true,
	"BWA": true,
	"HRV": true,
	"KHM": true,
	"IDN": true,
	"PRT": true,
	"MLT": true,
	"TJK": true,
	"VNM": true,
	"CYM": true,
	"PRY": true,
	"SHN": true,
	"CYP": true,
	"SYC": true,
	"RWA": true,
	"BGD": true,
	"AUS": true,
	"AUT": true,
	"LKA": true,
	"PSE": true,
	"GAB": true,
	"ZWE": true,
	"BGR": true,
	"SYR": true,
	"CZE": true,
	"NOR": true,
	"UMI": true,
	"CIV": true,
	"MMR": true,
	"TKL": true,
	"KIR": true,
	"TKM": true,
	"GRD": true,
	"GRC": true,
	"PCN": true,
	"HTI": true,
	"GRL": true,
	"YEM": true,
	"AFG": true,
	"MNE": true,
	"MNG": true,
	"NPL": true,
	"BHS": true,
	"BHR": true,
	"MNP": true,
	"GBR": true,
	"SJM": true,
	"DMA": true,
	"TLS": true,
	"BIH": true,
	"HUN": true,
	"AGO": true,
	"WSM": true,
	"FRA": true,
	"MOZ": true,
	"NAM": true,
	"PER": true,
	"DNK": true,
	"GTM": true,
	"FRO": true,
	"SLB": true,
	"VAT": true,
	"SLE": true,
	"NRU": true,
	"AIA": true,
	"GUF": true,
	"ZZZ": true,
	"SLV": true,
	"GUM": true,
	"FSM": true,
	"DOM": true,
	"CMR": true,
	"GUY": true,
	"AZE": true,
	"MAC": true,
	"GEO": true,
	"TON": true,
	"MAF": true,
	"NCL": true,
	"SMR": true,
	"ERI": true,
	"KNA": true,
	"MAR": true,
	"BLM": true,
	"VCT": true,
	"BLR": true,
	"MRT": true,
	"BLZ": true,
	"PHL": true,
	"COD": true,
	"COG": true,
	"ESH": true,
	"PYF": true,
	"URY": true,
	"COK": true,
	"COM": true,
	"COL": true,
	"USA": true,
	"ESP": true,
	"EST": true,
	"BMU": true,
	"MSR": true,
	"ZMB": true,
	"KOR": true,
	"SOM": true,
	"VUT": true,
	"ALA": true,
	"ECU": true,
	"ALB": true,
	"ETH": true,
	"GGY": true,
	"MCO": true,
	"NER": true,
	"LAO": true,
	"VEN": true,
	"GHA": true,
	"CPV": true,
	"MDA": true,
	"MTQ": true,
	"MDG": true,
	"SPM": true,
	"NFK": true,
	"LBN": true,
	"LBR": true,
	"BOL": true,
	"MDV": true,
	"GIB": true,
	"LBY": true,
	"HKG": true,
	"CAF": true,
	"LSO": true,
	"NGA": true,
	"MUS": true,
	"IMN": true,
	"LCA": true,
	"JOR": true,
	"GIN": true,
	"VGB": true,
	"CAN": true,
	"TCA": true,
	"TCD": true,
	"AND": true,
	"ROU": true,
	"CRI": true,
	"IND": true,
	"MEX": true,
	"SRB": true,
	"KAZ": true,
	"SAU": true,
	"JPN": true,
	"LTU": true,
	"TTO": true,
	"PLW": true,
	"HMD": true,
	"MWI": true,
	"SSD": true,
	"NIC": true,
	"CCK": true,
	"FIN": true,
	"TUN": true,
	"LUX": true,
	"UGA": true,
	"IOT": true,
	"BRA": true,
	"TUR": true,
	"TUV": true,
	"DEU": true,
	"EGY": true,
	"LVA": true,
	"JAM": true,
	"NIU": true,
	"VIR": true,
	"ZAF": true,
	"BRN": true,
	"HND": true,
}
