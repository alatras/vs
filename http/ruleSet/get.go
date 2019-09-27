package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func (resp GetRuleSetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rs Resource) Get(w http.ResponseWriter, r *http.Request) {
	app := rs.getRuleSetAppFactory()

	ctx := r.Context()
	entityId := chi.URLParam(r, "id")
	ruleSetId := chi.URLParam(r, "ruleSetId")

	fetchedRuleSet, err := app.Execute(ctx, entityId, ruleSetId)

	if err != nil {
		switch err {
		case getRuleSet.NotFound:
			_ = render.Render(w, r, errorResponse.ResourceNotFound("ruleSet", ruleSetId))
		case getRuleSet.UnexpectedError:
			_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		}
		return
	}

	rules := make([]RulePayload, len(fetchedRuleSet.RuleMetadata))

	for index, currentRule := range fetchedRuleSet.RuleMetadata {
		appRule := RulePayload{
			Key:      string(currentRule.Property),
			Operator: string(currentRule.Operator),
			Value:    currentRule.Value,
		}

		rules[index] = appRule
	}

	response := GetRuleSetResponse{
		Id:     fetchedRuleSet.Id,
		Entity: fetchedRuleSet.EntityId,
		Name:   fetchedRuleSet.Name,
		Action: string(fetchedRuleSet.Action),
		Rules:  rules,
	}

	render.Status(r, http.StatusOK)

	err = render.Render(w, r, response)

	if err != nil {
		rs.logger.Error.WithError(err).Error("error rendering response")
	}
}