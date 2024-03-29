package test

import (
	"context"
	"errors"
	"testing"
	"validation-service/app/deleteRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_DeleteRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := &stubRepository{}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := deleteRuleSet.NewDeleteRuleSet(log, newRec, repo)

	err := app.Execute(
		context.TODO(),
		"12345",
		"12345",
	)

	if err == nil {
		t.Error("RuleSet deletion succeeded but should fail with unexpected error")
	} else if err != deleteRuleSet.UnexpectedError {
		t.Errorf("RuleSet deletion failed but not with unexpected error: %v", err)
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
	panic("implement me")
}

func (s *stubRepository) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	return false, errors.New("unexpected error")
}

func (s *stubRepository) Ping(ctx context.Context) error {
	panic("implement me")
}
