package listAncestors

import (
	"context"
	"validation-service/app/listAncestorsRuleSet"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

const (
	unexpectedErrorMessage  = "Unexpected error: if the error persists, please contact an administrator, quoting the code and timestamp of this error"
	resourceNotFoundMessage = "The requested resource, or one of its sub-resources, can't be " +
		"found. If the submitted query is valid, this error is likely to be caused by a problem with a nested " +
		"resource that has been deleted or modified. Check the details property for additional insights."
	malformedParametersMessage = "At least one parameter is invalid. Examine the details " +
		"property for more information. Invalid parameters are listed and prefixed accordingly: body for parameters " +
		"submitted in the request's body, query for parameters appended to the request's URL, and params for " +
		"templated parameters of the request's URL."
)

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

func (app *successApp) Execute(ctx context.Context, entityIds []string) ([]ruleSet.RuleSet, listAncestorsRuleSet.AppError) {
	return mockRuleSets, listAncestorsRuleSet.AppError{}
}

type errorApp struct {
	error listAncestorsRuleSet.AppError
}

func (app *errorApp) Execute(ctx context.Context, entityIds []string) ([]ruleSet.RuleSet, listAncestorsRuleSet.AppError) {
	return []ruleSet.RuleSet{}, app.error
}
