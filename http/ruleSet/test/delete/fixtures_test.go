package delete

import (
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
)

const (
	unexpectedErrorMessage = "Unexpected error: if the error persists, please contact an administrator, quoting the code and timestamp of this error"
	notFoundErrorMessage   = "The requested resource, or one of its sub-resources, can't be " +
		"found. If the submitted query is valid, this error is likely to be caused by a problem with a nested " +
		"resource that has been deleted or modified. Check the details property for additional insights."
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

func (app *successApp) Execute(ctx context.Context, entityId, ruleSetId string) error {
	return nil
}

type errorApp struct {
	error error
}

func (app *errorApp) Execute(ctx context.Context, entityId, ruleSetId string) error {
	return app.error
}
