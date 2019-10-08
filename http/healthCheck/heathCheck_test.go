package healthCheck

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupSuccessRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := NewResource(log, &StubRepositorySuccess{}, &StubEntityServiceSuccess{})

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func setupUnexpectedErrorRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := NewResource(log, &StubRepositoryError{}, &StubEntityServiceSuccess{})

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func setupSuccessRecorderEntityService(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := NewResource(log, &StubRepositorySuccess{}, &StubEntityServiceSuccess{})

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func setupUnexpectedErrorRecorderEntityService(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := NewResource(log, &StubRepositorySuccess{}, &StubEntityServiceError{})

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

	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("Status code expected to be %d but got %d", http.StatusInternalServerError, status)
	}
}

func Test_HTTP_HealthCheck_Get_Success_EntityService(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupSuccessRecorderEntityService(t, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d but got %d", http.StatusOK, status)
	}
}

func Test_HTTP_HealthCheck_Get_Failure_EntityService(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupUnexpectedErrorRecorderEntityService(t, req)

	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("Status code expected to be %d but got %d", http.StatusInternalServerError, status)
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

type StubEntityServiceSuccess struct {
}

func (e *StubEntityServiceSuccess) Ping() error {
	return nil
}

func (e *StubEntityServiceSuccess) GetAncestorsOf(entityId string) ([]string, error) {
	panic("implement me")
}

func (e *StubEntityServiceSuccess) GetDescendantsOf(entityId string) ([]string, error) {
	panic("implement me")
}

type StubEntityServiceError struct {
}

func (e *StubEntityServiceError) Ping() error {
	return errors.New("service is down")
}

func (e *StubEntityServiceError) GetAncestorsOf(entityId string) ([]string, error) {
	panic("implement me")
}

func (e *StubEntityServiceError) GetDescendantsOf(entityId string) ([]string, error) {
	panic("implement me")
}
