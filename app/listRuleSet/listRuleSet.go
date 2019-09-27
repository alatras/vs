package listRuleSet

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
)

type ListRuleSet interface {
	Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error)
}

type App struct {
	instrumentation *instrumentation
	repository      ruleSet.Repository
}

func NewListRuleSet(logger *logger.Logger, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation: newInstrumentation(logger),
		repository:      ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityId": entityId,
	})
	app.instrumentation.startListingRuleSet()

	ruleSets, err := app.repository.ListByEntityId(ctx, entityId)

	if err != nil {
		app.instrumentation.failedListingRuleSet(err)
		return nil, errors.New("unexpected error")
	}

	app.instrumentation.finishListingRuleSet()
	return ruleSets, nil
}
