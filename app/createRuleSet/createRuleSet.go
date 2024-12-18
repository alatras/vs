package createRuleSet

import (
	"context"
	"errors"
	"fmt"

	"validation-service/logger"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

var (
	InvalidAction   = errors.New("action should be TAG or BLOCK")
	InvalidRule     = errors.New("invalid rule")
	UnexpectedError = errors.New("unexpected error")
)

type Rule struct {
	Key      string
	Operator string
	Value    string
}

type CreateRuleSet interface {
	Execute(ctx context.Context, entityId string, name string, action string, rules []Rule, tag string) (*ruleSet.RuleSet, error)
}

type App struct {
	instrumentation   *instrumentation
	ruleSetRepository ruleSet.Repository
}

func NewCreateRuleSet(logger *logger.Logger, record *logger.LogRecord, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation:   newInstrumentation(logger, record),
		ruleSetRepository: ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityId string, name string, action string, rules []Rule, tag string) (*ruleSet.RuleSet, error) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityId": entityId,
		"name":     name,
		"action":   action,
		"rules":    rules,
		"tag":      tag,
	})

	app.instrumentation.startCreatingRuleSet()

	rulesetAction := ruleSet.Action(action)

	if rulesetAction != ruleSet.Tag && rulesetAction != ruleSet.Block {
		app.instrumentation.invalidAction(action)
		return nil, InvalidAction
	}

	ruleMetadataArray := make([]rule.Metadata, len(rules))

	for index, currentRule := range rules {
		ruleMetadata := rule.Metadata{
			Property: rule.Property(currentRule.Key),
			Operator: rule.Operator(currentRule.Operator),
			Value:    currentRule.Value,
		}

		if app.isPropertyBlacklisted(ruleMetadata.Property) {
			err := fmt.Errorf("creation of rules for key '%s' is not allowed", currentRule.Key)
			app.instrumentation.ruleMetadataInvalid(ruleMetadata, err)
			return nil, InvalidRule
		}

		if _, err := rule.NewValidator(ruleMetadata); err != nil {
			app.instrumentation.ruleMetadataInvalid(ruleMetadata, err)
			return nil, InvalidRule
		}

		ruleMetadataArray[index] = ruleMetadata
	}

	newRuleSet := ruleSet.New(
		entityId,
		name,
		ruleSet.Action(action),
		ruleMetadataArray,
		tag,
	)

	if err := app.ruleSetRepository.Create(ctx, newRuleSet); err != nil {
		app.instrumentation.rulesetCreationFailed(err)
		return nil, UnexpectedError
	}

	app.instrumentation.finishCreatingRuleSet(newRuleSet)

	return &newRuleSet, nil
}

func (app App) isPropertyBlacklisted(property rule.Property) bool {
	switch property {
	case rule.PropertyCard:
		return true
	}

	return false
}
