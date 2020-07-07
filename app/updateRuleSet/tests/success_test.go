package tests

import (
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
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

	app := updateRuleSet.NewUpdateRuleSet(log, repo)

	updatedRuleSet, err := app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockUpdatedRuleSet.Action),
		mockUpdateRules,
	)

	if err != nil {
		t.Errorf("Failed to update RuleSet: %v", err)
		return
	}

	AssertRuleSet(t, mockUpdatedRuleSet, *updatedRuleSet)
}
