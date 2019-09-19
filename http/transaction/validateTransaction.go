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

const numOfWorkers = 6

/*
	Required to be implemented so that chi can bind the data to the payload struct
	Validate the request and return error if validation fails
 */
func (t validateTransactionPayload) Bind(r *http.Request) error {
	if t.Amount == 0 {
		return errors.New("amount required")
	}

	if t.Organization == "" {
		return errors.New("organization required")
	}

	return nil
}

func (t validateTransactionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func response(report report.Report) *validateTransactionResponse {

	blocked := []ruleSetResponse{}
	tagged := []ruleSetResponse{}

	for i:=0;i<len(report.BlockedRuleSets); i++ {
		brs := ruleSetResponse{
			Name: report.BlockedRuleSets[i].Name,
		}
		var md []metadata

		for j:= 0; j < len(report.BlockedRuleSets[i].Metadata); j++ {
			md = append(md, metadata(report.BlockedRuleSets[i].Metadata[j]))
		}

		brs.Metadata = md

		blocked = append(blocked, brs)
	}

	for i:=0;i<len(report.TaggedRuleSets); i++ {
		trs := ruleSetResponse{
			Name: report.TaggedRuleSets[i].Name,
		}

		var md []metadata

		for j:= 0; j < len(report.TaggedRuleSets[i].Metadata); j++ {
			md = append(md, metadata(report.TaggedRuleSets[i].Metadata[j]))
		}

		trs.Metadata = md

		tagged = append(tagged, trs)
	}

	resp := &validateTransactionResponse{
		Action: report.Action,
		BlockedRuleSets: blocked,
		TaggedRuleSets: tagged,
	}

	return resp
}

func (rs Resource) Validate(w http.ResponseWriter, r *http.Request) {
	ruleSetRepository, err := ruleSet.NewStubRuleSetRepository()

	details := make(map[string]interface{})

	if err != nil {
		_ = render.Render(w, r, httpError.UnexpectedError(details))
	}

	trxPayload := validateTransactionPayload{}

	if err := render.Bind(r, &trxPayload); err != nil {
		details["error"]= err.Error()
		_ = render.Render(w, r, httpError.MalformedParameters(details))
		return
	}

	validator := app.NewValidatorService(numOfWorkers, ruleSetRepository)

	trx := transaction.Transaction{
		Amount: trxPayload.Amount,
		Organization: trxPayload.Organization,
	}

	rpt := <-validator.Enqueue(trx)

	render.Status(r, http.StatusOK)
	_ = render.Render(w, r, response(rpt))
}