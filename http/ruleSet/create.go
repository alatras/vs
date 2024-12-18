package ruleSet

import (
	"net/http"
	"validation-service/app/createRuleSet"
	appd "validation-service/appdynamics"
	"validation-service/enums/contextKey"
	"validation-service/http/errorResponse"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (payload CreateRuleSetPayload) Bind(r *http.Request) error {
	err := payload.Validate()
	return err
}

func (t CreateRuleSetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	app := rs.createRuleSetAppFactory()

	entityId := chi.URLParam(r, "id")

	payload := CreateRuleSetPayload{}

	if err := render.Bind(r, &payload); err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)
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

	tag := payload.Tag

	ruleSet, err := app.Execute(ctx, entityId, payload.Name, payload.Action, rules, tag)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)

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

	response := CreateRuleSetResponse{
		CreateRuleSetPayload: payload,
		Id:                   ruleSet.Id,
		Entity:               entityId,
	}

	render.Status(r, http.StatusCreated)

	err = render.Render(w, r, response)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)
		rs.logger.Error.WithError(err).Error("error rendering response")
	}
}
