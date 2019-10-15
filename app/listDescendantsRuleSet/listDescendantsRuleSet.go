package listDescendantsRuleSet

import (
	"bitbucket.verifone.com/validation-service/entityService"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
)

var (
	UnexpectedError = errors.New("unexpected error")
)

type ListDescendantsRuleSet interface {
	Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error)
}

type App struct {
	instrumentation *instrumentation
	repository      ruleSet.Repository
	entityService   entityService.EntityService
}

func NewListDescendantsRuleSet(logger *logger.Logger, ruleSetRepository ruleSet.Repository, entityService entityService.EntityService) *App {
	return &App{
		instrumentation: newInstrumentation(logger),
		repository:      ruleSetRepository,
		entityService:   entityService,
	}
}

func (app *App) Execute(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityId": entityId,
	})
	app.instrumentation.startListingDescendantsRuleSet()

	entityIds, err := app.entityService.GetDescendantsOf(entityId)

	if err != nil {
		app.instrumentation.failedGetDescendants(err)
		return nil, UnexpectedError
	}

	ruleSets, err := app.repository.ListByEntityIds(ctx, entityIds...)

	if err != nil {
		app.instrumentation.failedListingDescendantsRuleSet(err)
		return nil, UnexpectedError
	}

	app.instrumentation.finishListingDescendantsRuleSet()
	return ruleSets, nil
}
