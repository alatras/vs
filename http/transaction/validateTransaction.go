package transaction

import (
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"bitbucket.verifone.com/validation-service/report"
	trx "bitbucket.verifone.com/validation-service/transaction"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

/*
	Required to be implemented so that chi can bind the data to the payload struct
	Validate the request and return error if validation fails
*/
func (t ValidateTransactionPayload) Bind(r *http.Request) error {
	if t.Transaction.Amount.Value == 0 {
		return errors.New("amount required")
	}

	if t.Transaction.Entity == "" {
		return errors.New("entity required")
	}

	return nil
}

func (t ValidateTransactionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func response(report report.Report) *ValidateTransactionResponse {
	return &ValidateTransactionResponse{report}
}

func (rs Resource) Validate(w http.ResponseWriter, r *http.Request) {
	var details interface{}

	trxPayload := ValidateTransactionPayload{}

	if err := render.Bind(r, &trxPayload); err != nil {
		_ = render.Render(w, r, errorResponse.MalformedParameters(err))
		return
	}

	ctx := r.Context()

	t := trx.Transaction{
		Amount:       trxPayload.Transaction.Amount.Value,
		CurrencyCode: trx.CurrencyCode(trxPayload.Transaction.Amount.CurrencyCode),
		EntityId:     trxPayload.Transaction.Entity,
	}

	reportChan, errChan := rs.app.Enqueue(ctx, t)

	select {
	case rep := <-reportChan:
		render.Status(r, http.StatusOK)
		err := render.Render(w, r, response(rep))
		if err != nil {
			rs.logger.Error.WithError(err).Error("error rendering response")
		}
	case err := <-errChan:
		rs.logger.Error.WithError(err).Error("error validating transaction")
		e := render.Render(w, r, errorResponse.UnexpectedError(details))
		if e != nil {
			rs.logger.Error.WithError(e).Error("error rendering response")
		}
	}
}
