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
func (body ValidateTransactionPayload) Bind(r *http.Request) error {
	if body.Transaction.Amount.Value == 0 {
		return errors.New("amount required")
	}

	if body.Transaction.Amount.MinorUnits == 0 {
		return errors.New("minor units required")
	}

	if body.Transaction.Amount.CurrencyCode == "" {
		return errors.New("currency code required")
	}

	if body.Transaction.Merchant.Id == "" {
		return errors.New("merchant ID required")
	}

	if body.Transaction.Customer.Country == "" {
		return errors.New("customer country code required")
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
		MinorUnits:   trxPayload.Transaction.Amount.MinorUnits,
		CurrencyCode: trx.CurrencyCode(trxPayload.Transaction.Amount.CurrencyCode),
		EntityId:     trxPayload.Transaction.Merchant.Id,
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
