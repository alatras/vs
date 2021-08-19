package tests

import (
	"context"
	"testing"
	"validation-service/app/updateRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_UpdateRuleSet_InvalidRule_Key(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := updateRuleSet.NewUpdateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidKeyRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}

func Test_App_UpdateRuleSet_InvalidRule_Operation(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := updateRuleSet.NewUpdateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidOperatorRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}

func Test_App_UpdateRuleSet_InvalidRule_Value(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := updateRuleSet.NewUpdateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidValueRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}

func Test_App_UpdateRuleSet_InvalidRule_NoValue(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := updateRuleSet.NewUpdateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidNoValueRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid rule error")
	} else if err != updateRuleSet.InvalidRule {
		t.Errorf("RuleSet update failed but not with invalid rule error: %v", err)
	}
}
