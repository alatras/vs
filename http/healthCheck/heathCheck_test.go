package healthCheck

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/logger"
	"validation-service/ruleSet"
)

func setupSuccessRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubHealthCheckLogger()

	resource := NewResource(log, &StubRepositorySuccess{})

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func setupUnexpectedErrorRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubHealthCheckLogger()

	resource := NewResource(log, &StubRepositoryError{})

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_HealthCheck_Get_Success(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupSuccessRecorder(t, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d but got %d", http.StatusOK, status)
	}
}

func Test_HTTP_HealthCheck_Get_Failure(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupUnexpectedErrorRecorder(t, req)

	if status := recorder.Code; status != http.StatusServiceUnavailable {
		t.Errorf("Status code expected to be %d but got %d", http.StatusServiceUnavailable, status)
	}
}

type StubRepositorySuccess struct {
}

func (s *StubRepositorySuccess) Create(ctx context.Context, ruleSet ruleSet.RuleSet) error {
	panic("implement me")
}

func (s *StubRepositorySuccess) GetById(ctx context.Context, entityId string, ruleSetId string) (*ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepositorySuccess) ListByEntityId(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepositorySuccess) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	panic("implement me")
}

func (s *StubRepositorySuccess) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}

func (s *StubRepositorySuccess) Ping(ctx context.Context) error {
	return nil
}

func (s *StubRepositorySuccess) ListByEntityIds(ctx context.Context, entityIds ...string) ([]ruleSet.RuleSet, error) {
	panic("implement me")
}

type StubRepositoryError struct {
}

func (s *StubRepositoryError) ListByEntityIds(ctx context.Context, entityIds ...string) ([]ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepositoryError) Create(ctx context.Context, ruleSet ruleSet.RuleSet) error {
	panic("implement me")
}

func (s *StubRepositoryError) GetById(ctx context.Context, entityId string, ruleSetId string) (*ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepositoryError) ListByEntityId(ctx context.Context, entityId string) ([]ruleSet.RuleSet, error) {
	panic("implement me")
}

func (s *StubRepositoryError) Replace(ctx context.Context, entityId string, ruleSet ruleSet.RuleSet) (bool, error) {
	panic("implement me")
}

func (s *StubRepositoryError) Delete(ctx context.Context, entityId string, ruleSetIds ...string) (bool, error) {
	panic("implement me")
}

func (s *StubRepositoryError) Ping(ctx context.Context) error {
	return errors.New("cannot connect to mongo")
}
