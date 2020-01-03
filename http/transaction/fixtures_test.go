package transaction

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/report"
	trx "bitbucket.verifone.com/validation-service/transaction"
	"context"
)

const (
	notFoundMessage        = "The entity ID in the transaction body was not found"
	unexpectedErrorMessage = "Unexpected error: if the error persists, please contact an administrator, quoting the code and timestamp of this error"
)

type app struct {
	rep *report.Report
	err *validateTransaction.ValidationError
}

func (a *app) Enqueue(ctx context.Context, tx trx.Transaction) (chan report.Report, chan validateTransaction.ValidationError) {
	resp := make(chan report.Report)
	error := make(chan validateTransaction.ValidationError)

	go func() {
		if a.err != nil {
			error <- *a.err
		} else {
			resp <- *a.rep
		}
	}()

	return resp, error
}

func (a *app) ResizeWorkers(numOfWorkers int) {}
