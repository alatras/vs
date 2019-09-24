package ruleset

import (
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

func (payload CreateRulesetPayload) Bind(r *http.Request) error {
	err := payload.Validate()
	return err
}

func (t CreateRulesetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entityId := chi.URLParam(r, "id")

	payload := CreateRulesetPayload{}

	if err := render.Bind(r, &payload); err != nil {
		_ = render.Render(w, r, errorResponse.MalformedParameters(err.Error()))
		return
	}

	ruleMetadataArray := make([]rule.Metadata, len(payload.Rules))

	for index, payloadRule := range payload.Rules {
		ruleMetadata := rule.Metadata{
			Property: rule.Property(payloadRule.Key),
			Operator: rule.Operator(payloadRule.Operator),
			Value:    payloadRule.Value,
		}

		ruleMetadataArray[index] = ruleMetadata
	}

	action := ruleSet.Action(strings.ToLower(payload.Action))

	newRuleSet, err := ruleSet.New(
		entityId,
		payload.Name,
		action,
		ruleMetadataArray,
	)

	if err != nil {
		_ = render.Render(w, r, errorResponse.UnexpectedError(err))
		return
	}

	err = rs.ruleSetRepository.Create(ctx, newRuleSet)

	if err != nil {
		_ = render.Render(w, r, errorResponse.UnexpectedError(err))
		return
	}

	response := CreateRulesetResponse{
		CreateRulesetPayload: payload,
		Id:                   newRuleSet.Id,
		Entity:               entityId,
	}

	render.Status(r, http.StatusOK)

	err = render.Render(w, r, response)

	if err != nil {
		rs.logger.Error.WithError(err).Error("error rendering response")
	}
}
