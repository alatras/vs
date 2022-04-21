package delete

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/deleteRuleSet"
	"validation-service/config"
	"validation-service/http/httpClient"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func setupSuccessRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	c := resty.New()
	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()
	responder := httpmock.NewStringResponder(200, "")
	fakeUrl := "/entities/12345"
	httpmock.RegisterResponder("GET", fakeUrl, responder)
	cl := httpClient.NewHttpClient(log, &config.Server{}, &logger.LogRecord{}, c)

	resource := ruleSet.NewResource(
		log,
		cl,
		nil,
		nil,
		func() deleteRuleSet.DeleteRuleSet {
			return &successApp{}
		},
		nil,
		nil,
		nil,
		nil,
	)

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_RuleSet_Delete_Success(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/12345/rulesets/"+mockRuleSet.Id, nil)
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

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
