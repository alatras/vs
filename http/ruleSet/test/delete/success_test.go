package delete

import (
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupSuccessRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(log, nil, nil, func() deleteRuleSet.DeleteRuleSet {
		return &successApp{}
	}, nil)

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_RuleSet_Delete_Success(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/12345/rulesets/"+mockRuleSet.Id, nil)

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	recorder := setupSuccessRecorder(t, req)

	if status := recorder.Code; status != http.StatusNoContent {
		t.Errorf("Status code expected to be %d but got %d", http.StatusNoContent, status)
		return
	}

	bodySize := len(recorder.Body.String())

	if bodySize != 0 {
		t.Errorf("Response should have no body but got body of size %d", bodySize)
		return
	}
}
