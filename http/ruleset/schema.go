package ruleset

import (
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"errors"
	"fmt"
)

type RulePayload struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type CreateRulesetPayload struct {
	Name   string        `json:"name"`
	Action string        `json:"action"`
	Rules  []RulePayload `json:"rules"`
}

type CreateRulesetResponse struct {
	CreateRulesetPayload
	Id     string `json:"id"`
	Entity string `json:"entity"`
}

func (r RulePayload) Validate() error {
	if r.Value == "" {
		return errors.New("value should be present")
	}

	_, err := rule.NewValidator(rule.Metadata{
		Property: rule.Property(r.Key),
		Operator: rule.Operator(r.Operator),
		Value:    r.Value,
	})

	return err
}

func (payload CreateRulesetPayload) Validate() error {
	if payload.Name == "" {
		return errors.New("body.name: should be present")
	}

	if payload.Action == "" {
		return errors.New("body.action: should be present")
	}

	action := ruleSet.Action(payload.Action)

	if action != ruleSet.Tag && action != ruleSet.Block {
		return errors.New("body.action: should be TAG or BLOCK")
	}

	if len(payload.Rules) == 0 {
		return errors.New("body.rules: at least one rule should be defined")
	}

	for index, rulePayload := range payload.Rules {
		if ruleError := rulePayload.Validate(); ruleError != nil {
			return fmt.Errorf("body.rules.%d: %s", index, ruleError.Error())
		}
	}

	return nil
}
