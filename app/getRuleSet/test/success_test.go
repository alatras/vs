package test

import (
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/test"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_GetRuleSet_Success(t *testing.T) {
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

	app := getRuleSet.NewGetRuleSet(log, repo)

	fetchedRuleSet, err := app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Id,
	)

	if err != nil {
		t.Errorf("Failed to fetch a RuleSet: %v", err)
		return
	}

	if fetchedRuleSet == nil {
		t.Error("RuleSet should be returned but it was not found")
		return
	}

	test.AssertRuleSet(t, mockRuleSet, *fetchedRuleSet)
}
