package delete

import (
	"context"
)

type SuccessApp struct {
}

func (app *SuccessApp) Execute(ctx context.Context, entityId, ruleSetId string) error {
	return nil
}

type ErrorApp struct {
	Error error
}

func (app *ErrorApp) Execute(ctx context.Context, entityId, ruleSetId string) error {
	return app.Error
}
