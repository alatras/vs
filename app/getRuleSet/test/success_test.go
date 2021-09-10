package test

import (
	"context"
	"testing"
	"validation-service/app/getRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_GetRuleSet_Success(t *testing.T) {
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
	app := getRuleSet.NewGetRuleSet(log, newRec, repo)

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

	assertRuleSet(t, mockRuleSet, *fetchedRuleSet)
}
