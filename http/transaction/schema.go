package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
)

type ValidateTransactionPayload struct {
	Transaction transaction `json:"transaction"`
}

type transaction struct {
	Amount   amount   `json:"amount"`
	Merchant merchant `json:"merchant"`
	Customer customer `json:"customer"`
}

type amount struct {
	Value        uint64 `json:"value"`
	MinorUnits   uint64 `json:"minorUnits"`
	CurrencyCode string `json:"currencyCode"`
}

type ValidateTransactionResponse struct {
	report.Report
}

type merchant struct {
	Id string `json:"id"`
}

type customer struct {
	Country string `json:"country"`
}
