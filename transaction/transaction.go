package transaction

type Transaction struct {
	Amount      int    `json:"amount"`
	CountryCode string `json:"country_code"`
	EntityId    string `json:"entity"`
}
