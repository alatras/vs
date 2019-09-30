package test

import (
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_ListRuleSet_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}
	for i := range mockRuleSets {
		err = repo.Create(context.TODO(), mockRuleSets[i])
	}

	if err != nil {
		t.Errorf("Failed to create mock rule set: %v", err)
		return
	}

	listApp := listRuleSet.NewListRuleSet(log, repo)

	ruleSets, err := listApp.Execute(
		context.TODO(),
		mockRuleSets[0].EntityId,
	)

	if err != nil {
		t.Errorf("Failed to list RuleSets: %v", err)
		return
	}

	for i := range mockRuleSets {
		AssertRuleSet(t, mockRuleSets[i], ruleSets[i])
	}

}
