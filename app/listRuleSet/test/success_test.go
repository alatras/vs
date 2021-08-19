package test

import (
	"context"
	"testing"
	"validation-service/app/listRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_ListRuleSet_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

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

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	listApp := listRuleSet.NewListRuleSet(log, newRec, repo)

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
