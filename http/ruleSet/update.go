package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func (payload UpdateRuleSetPayload) Bind(r *http.Request) error {
	err := payload.Validate()
	return err
}

func (u UpdateRuleSetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rs Resource) Update(w http.ResponseWriter, r *http.Request) {
	app := rs.updateRuleSetAppFactory()

	ctx := r.Context()
	entityId := chi.URLParam(r, "id")
	ruleSetId := chi.URLParam(r, "ruleSetId")

	payload := UpdateRuleSetPayload{}

	if err := render.Bind(r, &payload); err != nil {
		_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		return
	}

	rules := make([]updateRuleSet.Rule, len(payload.Rules))

	for i, rule := range payload.Rules {
		appRule := updateRuleSet.Rule{
			Key:      rule.Key,
			Operator: rule.Operator,
			Value:    rule.Value,
		}

		rules[i] = appRule
	}

	ruleSet, err := app.Execute(ctx, entityId, ruleSetId, payload.Name, payload.Action, rules)

	if err != nil {
		switch err {
		case updateRuleSet.InvalidAction:
			_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		case updateRuleSet.InvalidRule:
			_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		case updateRuleSet.UnexpectedError:
			_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		case updateRuleSet.NotFound:
			_ = render.Render(w, r, errorResponse.ResourceNotFound("ruleSet", ruleSetId))
		}
		return
	}

	response := UpdateRuleSetResponse{
		UpdateRuleSetPayload: payload,
		Id:                   ruleSet.Id,
		Entity:               entityId,
	}

	render.Status(r, http.StatusOK)

	err = render.Render(w, r, response)

	if err != nil {
		rs.logger.Error.WithError(err).Error("error rendering response")
	}
}
