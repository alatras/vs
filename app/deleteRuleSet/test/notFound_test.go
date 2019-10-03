package test

import (
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_DeleteRuleSet_NotFound(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := deleteRuleSet.NewDeleteRuleSet(log, repo)

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