package create

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/http/ruleSet/test"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
)

type SuccessApp struct {
}

func (app *SuccessApp) Execute(ctx context.Context, entityId string, name string, action string, rules []createRuleSet.Rule) (*ruleSet.RuleSet, error) {
	return &test.MockRuleSet, nil
}

type ErrorApp struct {
	Error error
}

func (app *ErrorApp) Execute(ctx context.Context, entityId string, name string, action string, rules []createRuleSet.Rule) (*ruleSet.RuleSet, error) {
	return nil, app.Error
}
