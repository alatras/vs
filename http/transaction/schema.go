package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
)

type ValidateTransactionPayload struct {
	Amount int    `json:"amount"`
	Entity string `json:"entity"`
}

type ValidateTransactionResponse struct {
	report.Report
}
