package tests

import (
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_UpdateRuleSet_NotFound(t *testing.T) {
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
		string(mockUpdatedRuleSet.Action),
		mockUpdateRules,
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with not found error")
	} else if err != updateRuleSet.NotFound {
		t.Errorf("RuleSet update failed but not with not found error: %v", err)
	}
}