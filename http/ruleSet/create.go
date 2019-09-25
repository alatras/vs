package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func (payload CreateRulesetPayload) Bind(r *http.Request) error {
	err := payload.Validate()
	return err
}

func (t CreateRulesetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	app := rs.createRulesetAppFactory()

	ctx := r.Context()
	entityId := chi.URLParam(r, "id")

	payload := CreateRulesetPayload{}

	if err := render.Bind(r, &payload); err != nil {
		_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		return
	}

	rules := make([]createRuleSet.Rule, len(payload.Rules))

	for index, currentRule := range payload.Rules {
		appRule := createRuleSet.Rule{
			Key:      currentRule.Key,
			Operator: currentRule.Operator,
			Value:    currentRule.Value,
		}

		rules[index] = appRule
	}

	ruleSet, err := app.Execute(ctx, entityId, payload.Name, payload.Action, rules)

	if err != nil {
		switch err {
		case createRuleSet.InvalidAction:
			_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		case createRuleSet.InvalidRule:
			_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		case createRuleSet.UnexpectedError:
			_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		}
		return
	}

	response := CreateRulesetResponse{
		CreateRulesetPayload: payload,
		Id:                   ruleSet.Id,
		Entity:               entityId,
	}

	render.Status(r, http.StatusOK)

	err = render.Render(w, r, response)

	if err != nil {
		rs.logger.Error.WithError(err).Error("error rendering response")
	}
}
