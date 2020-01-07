package transaction

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/report"
	trx "bitbucket.verifone.com/validation-service/transaction"
	"context"
)

const (
	notFoundMessage = "The requested resource, or one of its sub-resources, can't be " +
		"found. If the submitted query is valid, this error is likely to be caused by a problem with a nested " +
		"resource that has been deleted or modified. Check the details property for additional insights."
	unexpectedErrorMessage     = "Unexpected error: if the error persists, please contact an administrator, quoting the code and timestamp of this error"
	malformedParametersMessage = "At least one parameter is invalid. Examine the details " +
		"property for more information. Invalid parameters are listed and prefixed accordingly: body for parameters " +
		"submitted in the request's body, query for parameters appended to the request's URL, and params for " +
		"templated parameters of the request's URL."
)

type app struct {
	rep *report.Report
	err *validateTransaction.AppError
}

func (a *app) Enqueue(ctx context.Context, tx trx.Transaction) (chan report.Report, chan validateTransaction.AppError) {
	resp := make(chan report.Report)
	error := make(chan validateTransaction.AppError)

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
