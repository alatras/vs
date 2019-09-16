package transaction

type Transaction struct {
	Amount       int    `json:"amount"`
	Organization string `json:"organization"`
}
