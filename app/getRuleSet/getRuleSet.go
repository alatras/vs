package getRuleSet

import (
	"context"
	"errors"
	"validation-service/logger"
	"validation-service/ruleSet"
)

var (
	NotFound        = errors.New("rule set not found")
	UnexpectedError = errors.New("unexpected error")
)

type GetRuleSet interface {
	Execute(ctx context.Context, entityId, ruleSetId string) (*ruleSet.RuleSet, error)
}

type App struct {
	instrumentation   *instrumentation
	ruleSetRepository ruleSet.Repository
}

func NewGetRuleSet(logger *logger.Logger, record *logger.LogRecord, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation:   newInstrumentation(logger, record),
		ruleSetRepository: ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityId, ruleSetId string) (*ruleSet.RuleSet, error) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityId":  entityId,
		"ruleSetId": ruleSetId,
	})

	app.instrumentation.startFetchingRuleSet()

	fetchedRuleSet, err := app.ruleSetRepository.GetById(ctx, entityId, ruleSetId)

	if err != nil {
		app.instrumentation.ruleSetFetchFailed(err)
		return nil, UnexpectedError
	}

	if fetchedRuleSet == nil {
		app.instrumentation.ruleSetNotFound()
		return nil, NotFound
	}

	app.instrumentation.finishFetchingRuleSet(fetchedRuleSet)

	return fetchedRuleSet, nil
}
