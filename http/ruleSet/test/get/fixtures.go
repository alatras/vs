package get

import (
	"bitbucket.verifone.com/validation-service/http/ruleSet/test"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
)

type SuccessApp struct {
}

func (app *SuccessApp) Execute(ctx context.Context, entityId, ruleSetId string) (*ruleSet.RuleSet, error) {
	return &test.MockRuleSet, nil
}

type ErrorApp struct {
	Error error
}

func (app *ErrorApp) Execute(ctx context.Context, entityId, ruleSetId string) (*ruleSet.RuleSet, error) {
	return nil, app.Error
}
