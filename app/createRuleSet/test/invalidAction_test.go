package test

import (
	"context"
	"testing"
	"validation-service/app/createRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_CreateRuleSet_InvalidAction(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := createRuleSet.NewCreateRuleSet(log, repo)

	_, err = app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		"INVALID",
		mockRules,
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with invalid action error")
	} else if err != createRuleSet.InvalidAction {
		t.Errorf("RuleSet creation failed but not with invalid action error: %v", err)
	}
}
