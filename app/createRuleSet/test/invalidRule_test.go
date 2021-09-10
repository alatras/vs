package test

import (
	"context"
	"testing"
	"validation-service/app/createRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_CreateRuleSet_InvalidRule_Key(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := createRuleSet.NewCreateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidKeyRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}

func Test_App_CreateRuleSet_InvalidRule_Operator(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := createRuleSet.NewCreateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidOperatorRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}

func Test_App_CreateRuleSet_InvalidRule_Value(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := createRuleSet.NewCreateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidValueRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}

func Test_App_CreateRuleSet_InvalidRule_NoValue(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := createRuleSet.NewCreateRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockInvalidNoValueRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid rule error")
	} else if err != createRuleSet.InvalidRule {
		t.Errorf("RuleSet creation failed but not with invalid rule error: %v", err)
	}
}
