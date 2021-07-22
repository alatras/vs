package ruleSet

import (
	"errors"
	"fmt"
	"validation-service/ruleSet"
)

type RulePayload struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type GetRuleSetResponse struct {
	Id     string        `json:"id"`
	Entity string        `json:"entity"`
	Name   string        `json:"name"`
	Action string        `json:"action"`
	Rules  []RulePayload `json:"rules"`
}

type CreateRuleSetPayload struct {
	Name   string        `json:"name"`
	Action string        `json:"action"`
	Rules  []RulePayload `json:"rules"`
	Tag    string        `json:"tag"`
}

type CreateRuleSetResponse struct {
	CreateRuleSetPayload
	Id     string `json:"id"`
	Entity string `json:"entity"`
}

type UpdateRuleSetPayload struct {
	Name   string        `json:"name"`
	Action string        `json:"action"`
	Rules  []RulePayload `json:"rules"`
	Tag    string        `json:"tag"`
}

type UpdateRuleSetResponse struct {
	UpdateRuleSetPayload
	Id     string `json:"id"`
	Entity string `json:"entity"`
}

func (r RulePayload) Validate() error {
	if r.Key == "" {
		return errors.New("key should be present")
	}

	if r.Operator == "" {
		return errors.New("operator should be present")
	}

	if r.Value == "" {
		return errors.New("value should be present")
	}

	return nil
}

func (payload CreateRuleSetPayload) Validate() error {
	if payload.Name == "" {
		return errors.New("body.name: should be present")
	}

	if payload.Action == "" {
		return errors.New("body.action: should be present")
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

func (payload UpdateRuleSetPayload) Validate() error {
	if payload.Name == "" {
		return errors.New("body.name: should be present")
	}

	if payload.Action == "" {
		return errors.New("body.action: should be present")
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

type ListRuleSetResponse struct {
	ruleSet.RuleSet
}
