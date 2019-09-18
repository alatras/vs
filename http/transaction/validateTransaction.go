package transaction

import (
	"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/common/httpError"
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type validateTransactionPayload struct {
	*transaction.Transaction
}

type validateTransactionResponse struct {
	Status string `json:"status,omitempty"`
}

func (t validateTransactionPayload) Bind(r *http.Request) error {
	if t.Transaction == nil {
		return errors.New("missing required Transaction fields")
	}

	return nil
}

func (t validateTransactionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ValidateTransactionResponse(report report.Report) *validateTransactionResponse {
	resp := &validateTransactionResponse{
		Status: "true",
	}

	return resp
}

func (rs Resource) Validate(w http.ResponseWriter, r *http.Request) {
	ruleSetRepository, err := ruleSet.NewStubRuleSetRepository()

	var details interface{}

	if err != nil {
		_ = render.Render(w, r, httpError.UnexpectedError(details))
	}

	trxPayload := validateTransactionPayload{}

	if err := render.Bind(r, &trxPayload); err != nil {
		_ = render.Render(w, r, httpError.UnexpectedError(err))
		return
	}

	validator := app.NewValidatorService(6, ruleSetRepository)

	trx := transaction.Transaction{
		Amount: trxPayload.Amount,
		Organization: trxPayload.Organization,
	}

	rpt := <-validator.Enqueue(trx)

	render.Status(r, http.StatusOK)
	_ = render.Render(w, r, ValidateTransactionResponse(rpt))
	return
}