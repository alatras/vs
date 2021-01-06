package tests

import (
	"context"
	"errors"
	"testing"
	"validation-service/app/updateRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_UpdateRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := stubRepository{}

	app := updateRuleSet.NewUpdateRuleSet(log, &repo)

	_, err := app.Execute(
		context.TODO(),
		mockUpdatedRuleSet.EntityId,
		mockUpdatedRuleSet.Id,
		mockUpdatedRuleSet.Name,
		string(mockUpdatedRuleSet.Action),
		mockUpdateRules,
	)

	if err == nil {
		t.Error("RuleSet update succeeded but should fail with unexpected error")
	} else if err != updateRuleSet.UnexpectedError {
		t.Errorf("RuleSet update failed but not with unexpected error: %v", err)
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
	panic("implement me")
}

func (s *stubRepository) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	return true, errors.New("unexpected error")
}

func (s *stubRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}

func (s *stubRepository) Ping(ctx context.Context) error {
	panic("implement me")
}
