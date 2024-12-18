package test

import (
	"context"
	"errors"
	"testing"
	"validation-service/app/getRuleSet"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func Test_App_GetRuleSet_UnexpectedError(t *testing.T) {
	log := logger.NewStubLogger()
	repo := &stubRepository{}

	var rec *logger.LogRecord
	newRec := rec.NewRecord()
	app := getRuleSet.NewGetRuleSet(log, newRec, repo)

	_, err := app.Execute(
		context.TODO(),
		"12345",
		"12345",
	)

	if err == nil {
		t.Error("RuleSet fetch succeeded but should fail with unexpected error")
	} else if err != getRuleSet.UnexpectedError {
		t.Errorf("RuleSet fetch failed but not with unexpected error: %v", err)
	}
}

type stubRepository struct {
}

func (s *stubRepository) Create(ctx context.Context, ruleSet ruleSet.RuleSet) error {
	panic("implement me")
}

func (s *stubRepository) GetById(ctx context.Context, entityId string, ruleSetId string) (*ruleSet.RuleSet, error) {
	return nil, errors.New("unexpected error")
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

func (s *stubRepository) Ping(ctx context.Context) error {
	panic("implement me")
}
