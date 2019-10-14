package tests

import (
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
	"testing"
)

func Test_App_UpdateRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := StubRepository{}

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

type StubRepository struct {
}

func (s *StubRepository) Create(ctx context.Context, ruleSet ruleSet.RuleSet) error {
	panic("implement me")
}

func (s *StubRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepository) ListByEntityId(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepository) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	return true, errors.New("unexpected error")
}

func (s *StubRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}
