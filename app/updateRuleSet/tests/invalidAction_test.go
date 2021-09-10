package tests

import (
	"context"
	"testing"
	"validation-service/app/updateRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_UpdateRuleSet_InvalidAction(t *testing.T) {
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
		"INVALID",
		mockUpdateRules,
		"TEST TAG",
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with invalid action error")
	} else if err != updateRuleSet.InvalidAction {
		t.Errorf("RuleSet update failed but not with invalid action error: %v", err)
	}
}
