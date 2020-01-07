package listAncestors

import (
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
)

const unexpectedErrorMessage = "Unexpected error: if the error persists, please contact an administrator, quoting the code and timestamp of this error"

var mockRuleSets = []ruleSet.RuleSet{ruleSet.New(
	"12345",
	"Test",
	ruleSet.Block,
	[]rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	},
)}

type successApp struct {
}

func (app *successApp) Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, listAncestorsRuleSet.AppError) {
	return mockRuleSets, listAncestorsRuleSet.AppError{}
}

type errorApp struct {
	error listAncestorsRuleSet.AppError
}

func (app *errorApp) Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, listAncestorsRuleSet.AppError) {
	return []ruleSet.RuleSet{}, app.error
}
