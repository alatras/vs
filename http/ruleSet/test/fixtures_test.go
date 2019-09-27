package test

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
)

const (
	malformedParametersErrorMessage = "At least one parameter is invalid. Examine the details property for more information. Invalid parameters are listed and prefixed accordingly: body for parameters submitted in the request's body, query for parameters appended to the request's URL, and params for templated parameters of the request's URL."
)

var (
	mockRuleSet = ruleSet.New("12345", "Test", ruleSet.Block, []rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	})
)

type successApp struct {
}

func (app *successApp) Execute(ctx context.Context, entityId string, name string, action string, rules []createRuleSet.Rule) (*ruleSet.RuleSet, error) {
	return &mockRuleSet, nil
}

type errorApp struct {
	error error
}

func (app *errorApp) Execute(ctx context.Context, entityId string, name string, action string, rules []createRuleSet.Rule) (*ruleSet.RuleSet, error) {
	return nil, app.error
}
