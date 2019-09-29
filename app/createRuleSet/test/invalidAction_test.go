package test

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"testing"
)

func Test_App_CreateRuleSet_InvalidAction(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

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
