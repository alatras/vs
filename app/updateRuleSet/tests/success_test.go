package tests

import (
	"context"
	"testing"
	"validation-service/app/updateRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_UpdateRuleSet_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	err = repo.Create(context.TODO(), mockRuleSet)

	if err != nil {
		t.Errorf("Failed to create mock rule set: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := updateRuleSet.NewUpdateRuleSet(log, newRec, repo)

	updatedRuleSet, err := app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockUpdatedRuleSet.Action),
		mockUpdateRules,
		"TEST TAG",
	)

	if err != nil {
		t.Errorf("Failed to update RuleSet: %v", err)
		return
	}

	AssertRuleSet(t, mockUpdatedRuleSet, *updatedRuleSet)
}
