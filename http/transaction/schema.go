package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
	trx "bitbucket.verifone.com/validation-service/transaction"
)

type ValidateTransactionPayload struct {
	Transaction transaction `json:"transaction"`
}

type transaction struct {
	Amount amount `json:"amount"`
	Entity string `json:"entity"`
}

type amount struct {
	Value        uint             `json:"value"`
	CurrencyCode trx.CurrencyCode `json:"currencyCode"`
}

type ValidateTransactionResponse struct {
	report.Report
}
