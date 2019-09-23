package transaction

import (
	"bitbucket.verifone.com/validation-service/common/rest"
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

/*
	Required to be implemented so that chi can bind the data to the payload struct
	Validate the request and return error if validation fails
*/
func (t ValidateTransactionPayload) Bind(r *http.Request) error {
	if t.Amount == 0 {
		return errors.New("amount required")
	}

	if t.Entity == "" {
		return errors.New("entity required")
	}

	return nil
}

func (t ValidateTransactionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func response(report report.Report) *ValidateTransactionResponse {

	blocked := []RuleSetResponse{}

	tagged := []RuleSetResponse{}

	for i := 0; i < len(report.BlockedRuleSets); i++ {
		brs := RuleSetResponse{
			Name: report.BlockedRuleSets[i].Name,
		}
		var md []Metadata

		for j := 0; j < len(report.BlockedRuleSets[i].Metadata); j++ {
			md = append(md, Metadata(report.BlockedRuleSets[i].Metadata[j]))
		}

		brs.Metadata = md

		blocked = append(blocked, brs)
	}

	for i := 0; i < len(report.TaggedRuleSets); i++ {
		trs := RuleSetResponse{
			Name: report.TaggedRuleSets[i].Name,
		}

		var md []Metadata

		for j := 0; j < len(report.TaggedRuleSets[i].Metadata); j++ {
			md = append(md, Metadata(report.TaggedRuleSets[i].Metadata[j]))
		}

		trs.Metadata = md

		tagged = append(tagged, trs)
	}

	resp := &ValidateTransactionResponse{
		Action:          report.Action,
		BlockedRuleSets: blocked,
		TaggedRuleSets:  tagged,
	}

	return resp
}

func (rs Resource) Validate(w http.ResponseWriter, r *http.Request) {
	details := make(map[string]interface{})

	trxPayload := ValidateTransactionPayload{}

	if err := render.Bind(r, &trxPayload); err != nil {
		details["error"] = err.Error()
		_ = render.Render(w, r, rest.MalformedParameters(details))
		return
	}

	ctx := rest.GetContextWithTraceId(r)

	trx := transaction.Transaction{
		Amount: trxPayload.Amount,
		Entity: trxPayload.Entity,
	}

	rpt := <-rs.app.Enqueue(trx, ctx)

	render.Status(r, http.StatusOK)
	_ = render.Render(w, r, response(rpt))
}
