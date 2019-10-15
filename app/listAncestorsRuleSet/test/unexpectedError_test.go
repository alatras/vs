package test

import (
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
	"testing"
)

func Test_App_ListAncestorsRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}
	entityService := stubEntityService{}

	app := listAncestorsRuleSet.NewListAncestorsRuleSet(log, &repo, &entityService)

	_, err := app.Execute(
		context.TODO(),
		"123",
	)

	if err == nil {
		t.Error("listing ancestors RuleSet succeeded but should fail with unexpected error")
	} else if err != listAncestorsRuleSet.UnexpectedError {
		t.Errorf("listing ancestors RuleSet failed but not with unexpected error: %v", err)
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

func (s *stubRepository) ListByEntityIds(ctx context.Context, entityIds ...string) ([]ruleSet.RuleSet, error) {
	return nil, errors.New("unexpected error")
}

func (s *stubRepository) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	panic("implement me")
}

func (s *stubRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}

func (s *stubRepository) Ping(ctx context.Context) error {
	panic("implement me")
}
