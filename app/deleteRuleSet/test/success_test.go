package test

import (
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_DeleteRuleSet_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	err = repo.Create(context.TODO(), mockRuleSet)

	if err != nil {
		t.Errorf("Failed to create mock rule set: %v", err)
		return
	}

	app := deleteRuleSet.NewDeleteRuleSet(log, repo)

	err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Id,
	)

	if err != nil {
		t.Errorf("Failed to delete a RuleSet: %v", err)
		return
	}

	deletedRuleSet, err := repo.GetById(context.TODO(), mockRuleSet.EntityId, mockRuleSet.Id)

	if err != nil {
		t.Errorf("Failed to get a RuleSet: %v", err)
		return
	}

	if deletedRuleSet != nil {
		t.Errorf("Able to fetch a rule set after deletion")
		return
	}
}
