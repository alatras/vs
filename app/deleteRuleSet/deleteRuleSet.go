package deleteRuleSet

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

type DeleteRuleSet interface {
	Execute(ctx context.Context, entityId, ruleSetId string) error
}

type App struct {
	instrumentation   *instrumentation
	ruleSetRepository ruleSet.Repository
}

func NewDeleteRuleSet(logger *logger.Logger, record *logger.LogRecord, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation:   newInstrumentation(logger, record),
		ruleSetRepository: ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityId, ruleSetId string) error {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityId":  entityId,
		"ruleSetId": ruleSetId,
	})

	app.instrumentation.startDeletingRuleSet()

	deleted, err := app.ruleSetRepository.Delete(ctx, entityId, ruleSetId)

	if err != nil {
		app.instrumentation.ruleSetDeletionFailed(err)
		return UnexpectedError
	}

	if !deleted {
		app.instrumentation.ruleSetNotFound()
		return NotFound
	}

	app.instrumentation.ruleSetDeleted()

	return nil
}
