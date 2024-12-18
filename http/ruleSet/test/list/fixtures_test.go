package list

import (
	"context"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
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
	"TEST TAG",
)}

type successApp struct {
}

func (app *successApp) Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	return mockRuleSets, nil
}

type errorApp struct {
	error error
}

func (app *errorApp) Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	return []ruleSet.RuleSet{}, app.error
}
