package transaction

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	appd "bitbucket.verifone.com/validation-service/appdynamics"
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"bitbucket.verifone.com/validation-service/report"
	trx "bitbucket.verifone.com/validation-service/transaction"
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"strings"
)

const cardInstrument = "CARD"

/*
	Required to be implemented so that chi can bind the data to the payload struct
	Validate the request and return error if validation fails
*/
func (body ValidateTransactionPayload) Bind(r *http.Request) error {
	if body.Transaction.Amount.Value == "" {
		return errors.New("amount required")
	}

	if body.Transaction.Amount.CurrencyCode == "" {
		return errors.New("currency code required")
	}

	if body.Transaction.Merchant.Organisation.UUID == "" {
		return errors.New("merchant organisation UUID required")
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
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	var details interface{}

	trxPayload := ValidateTransactionPayload{}

	if err := render.Bind(r, &trxPayload); err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), false)
		_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		return
	}

	minorUnits := 0

	amountComponents := strings.Split(trxPayload.Transaction.Amount.Value, ".")

	amount, err := strconv.ParseUint(amountComponents[0], 10, 64)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), false)
		_ = render.Render(w, r, errorResponse.MalformedParameters(err))
		return
	}

	numberOfAmountComponents := len(amountComponents)

	if numberOfAmountComponents == 2 {
		decimalAmountString := amountComponents[1]

		decimalAmount, err := strconv.ParseUint(decimalAmountString, 10, 64)

		if err != nil {
			appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), false)
			_ = render.Render(w, r, errorResponse.MalformedParameters(err))
			return
		}

		minorUnits = len(decimalAmountString)

		for i := 0; i < minorUnits; i++ {
			amount *= 10
		}

		amount += decimalAmount
	} else if numberOfAmountComponents > 2 {
		errMessage := "amount can contain only one decimal point"
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, errMessage, false)
		_ = render.Render(w, r, errorResponse.MalformedParameters(errMessage))
		return
	}

	entityId := trxPayload.Transaction.Merchant.Organisation.UUID

	t := trx.Transaction{
		Amount:       amount,
		MinorUnits:   minorUnits,
		CurrencyCode: trx.CurrencyCode(trxPayload.Transaction.Amount.CurrencyCode),
		EntityId:     entityId,
	}

	if trxPayload.Transaction.Customer != (customer{}) {
		trxCustomer := trxPayload.Transaction.Customer

		if trxCustomer.Country != "" {
			t.CustomerCountryCode = trx.CountryCodeIso31661Alpha2(trxCustomer.Country)
		}

		t.CustomerId = trxCustomer.CustomerIdentification.CustomerId
		t.CustomerIP = trxCustomer.IP
		t.CustomerIPCountry = trxCustomer.IPCountry
	}

	for i := range trxPayload.Transaction.Instrument {
		if trxPayload.Transaction.Instrument[i].Type == cardInstrument {
			t.Card = trxPayload.Transaction.Instrument[i].CardNumber
			t.IssuerCountryCode = trx.CountryCodeIso31661Alpha3(trxPayload.Transaction.Instrument[i].Country)
		}
	}

	reportChan, errChan := rs.app.Enqueue(ctx, t)

	select {
	case rep := <-reportChan:
		render.Status(r, http.StatusOK)
		err := render.Render(w, r, response(rep))
		if err != nil {
			appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), false)
			rs.logger.Error.WithError(err).Error("error rendering response")
		}
	case validationError := <-errChan:
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, validationError.Error(), false)
		rs.logger.Error.WithError(validationError).Error("error validating transaction")

		var e error

		if validationError.Is(validateTransaction.EntityIdNotFoundErr) {
			e = render.Render(w, r, errorResponse.ResourceNotFound("entity", entityId))
		} else if validationError.Is(validateTransaction.EntityIdFormatIncorrectErr) {
			e = render.Render(w, r, errorResponse.MalformedParameters(map[string]string{
				"body.transaction.merchant.organisation.UUID": validationError.Error(),
			}))
		} else {
			e = render.Render(w, r, errorResponse.UnexpectedError(details))
		}

		if e != nil {
			appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, e.Error(), false)
			rs.logger.Error.WithError(e).Error("error rendering response")
		}
	}
}
