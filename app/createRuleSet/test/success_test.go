package test

import (
	"context"
	"testing"
	"validation-service/app/createRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

func Test_App_CreateRuleSet_Success(t *testing.T) {
	log := logger.NewStubLogger()
	repo, err := ruleSet.NewStubRepository(nil)

	if err != nil {
		t.Errorf("Failed to init stub repository: %v", err)
		return
	}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := createRuleSet.NewCreateRuleSet(log, newRec, repo)

	newRuleSet, err := app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockRules,
		"TEST TAG",
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
		repo, err := ruleSet.NewStubRepository(nil)

		if err != nil {
			t.Errorf("Failed to init stub repository: %v", err)
			return
		}

		var rec *logger.LogRecord
		newRec := rec.NewRecord()
		app := createRuleSet.NewCreateRuleSet(log, newRec, repo)

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
			"TEST TAG",
		)

		newRuleSet, err := app.Execute(
			context.TODO(),
			mockRuleSet.EntityId,
			mockRuleSet.Name,
			string(mockRuleSet.Action),
			mockRules,
			"TEST TAG",
		)

		if err != nil {
			t.Errorf("Failed to create RuleSet: %v", err)
			return
		}

		AssertRuleSet(t, mockRuleSet, *newRuleSet)
	}
}
