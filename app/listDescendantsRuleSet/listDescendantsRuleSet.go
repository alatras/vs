package listDescendantsRuleSet

import (
	"context"
	"validation-service/logger"
	"validation-service/ruleSet"
)

type ListDescendantsRuleSet interface {
	Execute(ctx context.Context, entityIds []string) ([]ruleSet.RuleSet, AppError)
}

type App struct {
	instrumentation *instrumentation
	repository      ruleSet.Repository
}

func NewListDescendantsRuleSet(logger *logger.Logger, record *logger.LogRecord, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation: newInstrumentation(logger, record),
		repository:      ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityIds []string) ([]ruleSet.RuleSet, AppError) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityIds": entityIds,
	})
	app.instrumentation.startListingDescendantsRuleSet()

	ruleSets, err := app.repository.ListByEntityIds(ctx, entityIds...)

	if err != nil {
		app.instrumentation.failedListingDescendantsRuleSet(err)
		return nil, NewError(UnexpectedErr, err)
	}

	app.instrumentation.finishListingDescendantsRuleSet()

	return ruleSets, AppError{}
}
