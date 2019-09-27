package test

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_CreateRuleSet_InvalidRule_Key(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := createRuleSet.NewCreateRuleset(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidKeyRules,
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}

func Test_App_CreateRuleSet_InvalidRule_Operator(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := createRuleSet.NewCreateRuleset(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidOperatorRules,
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}

func Test_App_CreateRuleSet_InvalidRule_Value(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := createRuleSet.NewCreateRuleset(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidValueRules,
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}

func Test_App_CreateRuleSet_InvalidRule_NoValue(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := createRuleSet.NewCreateRuleset(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidNoValueRules,
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}