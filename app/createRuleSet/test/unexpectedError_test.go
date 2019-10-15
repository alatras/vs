package test

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
	"testing"
)

func Test_App_CreateRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}

	app := createRuleSet.NewCreateRuleSet(log, &repo)

	_, err := app.Execute(
		context.TODO(),
		mockRuleSet.EntityId,
		mockRuleSet.Name,
		string(mockRuleSet.Action),
		mockRules,
	)

	if err == nil {
		t.Error("RuleSet creation succeeded but should fail with unexpected error")
	} else if err != createRuleSet.UnexpectedError {
		t.Errorf("RuleSet creation failed but not with unexpected error: %v", err)
	}
}

type stubRepository struct {
}

func (s *stubRepository) Create(ctx context.Context, ruleSet ruleSet.RuleSet) error {
	return errors.New("unexpected error")
}

func (s *stubRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *stubRepository) ListByEntityIds(ctx context.Context, entityIds ...string) ([]ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *stubRepository) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	panic("implement me")
}

func (s *stubRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}
