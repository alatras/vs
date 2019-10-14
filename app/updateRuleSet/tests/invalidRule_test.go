package tests

import (
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_UpdateRuleSet_InvalidRule_Key(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := updateRuleSet.NewUpdateRuleSet(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidKeyRules,
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}

func Test_App_UpdateRuleSet_InvalidRule_Operation(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := updateRuleSet.NewUpdateRuleSet(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidOperatorRules,
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}

func Test_App_UpdateRuleSet_InvalidRule_Value(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := updateRuleSet.NewUpdateRuleSet(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidValueRules,
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}

func Test_App_UpdateRuleSet_InvalidRule_NoValue(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := updateRuleSet.NewUpdateRuleSet(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidNoValueRules,
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}
