package test

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
	"testing"
)

func Test_App_CreateRuleSet_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	app := createRuleSet.NewCreateRuleSet(log, repo)

	newRuleSet, err := app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockRules,
	)

	if err != nil {
		t.Errorf("Failed to create RuleSet: %v", err)
		return
	}

	AssertRuleSet(t, mockRuleSet, *newRuleSet)
}

func Test_App_CreateRuleSet_Success_Amount_AllOperators(t *testing.T) {
	amountOperators := []string{"<", "<=", "==", "!=", ">=", ">"}

	for _, operator := range amountOperators {
		log := logger.NewStubLogger()
		repo, err := ruleSet.NewStubRepository()

		if err != nil {
			t.Errorf("Failed to init stub repository: %v", err)
			return
		}

		app := createRuleSet.NewCreateRuleSet(log, repo)

		mockRules := []createRuleSet.Rule{
			{
				Key:      "amount",
				Operator: operator,
				Value:    "1000",
			},
		}

		mockRuleSet := ruleSet.New(
			mockRuleSet.EntityId,
			mockRuleSet.Name,
			mockRuleSet.Action,
			[]rule.Metadata{
				{
					Property: "amount",
					Operator: rule.Operator(operator),
					Value:    "1000",
				},
			},
		)

		newRuleSet, err := app.Execute(
			context.TODO(),
			mockRuleSet.EntityId,
			mockRuleSet.Name,
			string(mockRuleSet.Action),
			mockRules,
		)

		if err != nil {
			t.Errorf("Failed to create RuleSet: %v", err)
			return
		}

		AssertRuleSet(t, mockRuleSet, *newRuleSet)
	}
}
