package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
)

type ValidateTransactionPayload struct {
	Transaction transaction `json:"transaction"`
}

type transaction struct {
	Amount amount `json:"amount"`
	Entity string `json:"entity"`
}

type amount struct {
	Value        uint64 `json:"value"`
	CurrencyCode string `json:"currencyCode"`
}

type ValidateTransactionResponse struct {
	report.Report
}
