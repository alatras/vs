package transaction

type Transaction struct {
	Amount   int    `json:"amount"`
	EntityId string `json:"entity"`
}
