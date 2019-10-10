package test

import (
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
	"testing"
)

func Test_App_ListRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}

	app := listRuleSet.NewListRuleSet(log, &repo)

	_, err := app.Execute(
		context.TODO(),
		"123",
	)

	if err == nil {
		t.Error("listing RuleSet succeeded but should fail with unexpected error")
	} else if err != listRuleSet.UnexpectedError {
		t.Errorf("listing RuleSet failed but not with unexpected error: %v", err)
	}

}

type stubRepository struct {
}

func (s *stubRepository) Create(ctx context.Context, ruleSet ruleSet.RuleSet) error {
	panic("implement me")
}

func (s *stubRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *stubRepository) ListByEntityId(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	return nil, errors.New("unexpected error")
}

func (s *stubRepository) ListByEntityIds(ctx context.Context, entityIds ...string) ([]ruleSet.RuleSet, error) {
	return nil, errors.New("unexpected error")
}

func (s *stubRepository) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	panic("implement me")
}

func (s *stubRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}
