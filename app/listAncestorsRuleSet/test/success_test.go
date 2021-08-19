package test

import (
	"context"
	"testing"
	"validation-service/app/listAncestorsRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_ListAncestorsRuleSet_Success(t *testing.T) {
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
	app := listAncestorsRuleSet.NewListAncestorsRuleSet(log, newRec, repo)

	ruleSets, error := app.Execute(
		context.TODO(),
		[]string{mockRuleSets[0].EntityId},
	)

	if error.HasError() {
		t.Errorf("Failed to list RuleSets: %v", err)
		return
	}

	for i := range mockRuleSets {
		AssertRuleSet(t, mockRuleSets[i], ruleSets[i])
	}

}
