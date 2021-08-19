package test

import (
	"context"
	"testing"
	"validation-service/app/deleteRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_DeleteRuleSet_NotFound(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := deleteRuleSet.NewDeleteRuleSet(log, newRec, repo)

	err = app.Execute(
		context.TODO(),
		"12345",
		"12345",
	)

	if err == nil {
		t.Error("RuleSet deletion succeeded but should fail with not found error")
	} else if err != deleteRuleSet.NotFound {
		t.Errorf("RuleSet deletion failed but not with not found error: %v", err)
	}
}
