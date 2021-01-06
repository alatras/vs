package updateRuleSet

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
	NotFound        = errors.New("rule set not found")
)

type Rule struct {
	Key      string
	Operator string
	Value    string
}

type UpdateRuleSet interface {
	Execute(ctx context.Context, entityId string, ruleSetId string, name string, action string, rules []Rule) (*ruleSet.RuleSet, error)
}

type App struct {
	instrumentation   *instrumentation
	ruleSetRepository ruleSet.Repository
}

func NewUpdateRuleSet(logger *logger.Logger, ruleSetRepository ruleSet.Repository) *App {
	return &App{
		instrumentation:   newInstrumentation(logger),
		ruleSetRepository: ruleSetRepository,
	}
}

func (app *App) Execute(ctx context.Context, entityId string, ruleSetId string, name string, action string, rules []Rule) (*ruleSet.RuleSet, error) {
	app.instrumentation.setContext(ctx)
	app.instrumentation.setMetadata(logger.Metadata{
		"ruleSetId": ruleSetId,
		"entityId":  entityId,
		"name":      name,
		"action":    action,
		"rules":     rules,
	})

	app.instrumentation.startUpdatingRuleSet()

	ruleSetAction := ruleSet.Action(action)

	if ruleSetAction != ruleSet.Tag && ruleSetAction != ruleSet.Block {
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
			err := fmt.Errorf("update of rules with key '%s' is not allowed", currentRule.Key)
			app.instrumentation.ruleMetadataInvalid(ruleMetadata, err)
			return nil, InvalidRule
		}

		if _, err := rule.NewValidator(ruleMetadata); err != nil {
			app.instrumentation.ruleMetadataInvalid(ruleMetadata, err)
			return nil, InvalidRule
		}

		ruleMetadataArray[index] = ruleMetadata
	}

	replaceRuleSet := ruleSet.From(
		ruleSetId,
		entityId,
		name,
		ruleSet.Action(action),
		ruleMetadataArray,
	)

	replaced, err := app.ruleSetRepository.Replace(ctx, entityId, replaceRuleSet)

	if err != nil {
		app.instrumentation.ruleSetUpdateFailed(err)
		return nil, UnexpectedError
	}

	if !replaced {
		app.instrumentation.ruleSetNotReplaced(replaceRuleSet)
		return nil, NotFound
	}

	app.instrumentation.finishUpdatingRuleSet(replaceRuleSet)

	return &replaceRuleSet, nil
}

func (app App) isPropertyBlacklisted(property rule.Property) bool {
	switch property {
	case rule.PropertyCard:
		return true
	}

	return false
}
