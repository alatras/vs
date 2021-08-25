package test

import (
	"context"
	"testing"
	"validation-service/app/getRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_GetRuleSet_NotFound(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := getRuleSet.NewGetRuleSet(log, newRec, repo)

	_, err = app.Execute(
		context.TODO(),
		"12345",
		"12345",
	)

	if err == nil {
		t.Error("RuleSet fetch succeeded but should fail with not found error")
	} else if err != getRuleSet.NotFound {
		t.Errorf("RuleSet fetch failed but not with not found error: %v", err)
	}
}
