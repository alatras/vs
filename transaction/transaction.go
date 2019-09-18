package transaction

type Transaction struct {
	Amount int    `json:"amount"`
	Entity string `json:"entity"`
}
