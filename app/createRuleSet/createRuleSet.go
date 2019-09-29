package createRuleSet

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
	"errors"
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
	Execute(ctx context.Context, entityId string, name string, action string, rules []Rule) (*ruleSet.RuleSet, error)
}

type App struct {
	instrumentation   *instrumentation
	ruleSetRepository ruleSet.Repository
}

func NewCreateRuleSet(logger *logger.Logger, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation:   newInstrumentation(logger),
		ruleSetRepository: ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityId string, name string, action string, rules []Rule) (*ruleSet.RuleSet, error) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"entityId": entityId,
		"name":     name,
		"action":   action,
		"rules":    rules,
	})

	app.instrumentation.startCreatingRuleset()

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
	)

	if err := app.ruleSetRepository.Create(ctx, newRuleSet); err != nil {
		app.instrumentation.rulesetCreationFailed(err)
		return nil, UnexpectedError
	}

	app.instrumentation.finishCreatingRuleset(newRuleSet)

	return &newRuleSet, nil
}
