package listDescendants

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/listDescendantsRuleSet"
	"validation-service/config"
	"validation-service/http/httpClient"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func setupSuccessRecorder(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
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
		nil,
		nil,
		nil,
		func() listDescendantsRuleSet.ListDescendantsRuleSet {
			return &successApp{}
		},
		nil,
	)

	resource.Routes().ServeHTTP(recorder, r)

	return recorder
}

func Test_HTTP_RuleSet_ListDescendants_Success(t *testing.T) {

	req, err := http.NewRequest("GET", "/12345/rulesets/descendants", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	recorder := setupSuccessRecorder(t, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d but got %d", http.StatusOK, status)
	}

	body := recorder.Body.String()

	expected := fmt.Sprintf(
		`[
			{
				"id": "%s",
				"name": "Test",
				"action": "BLOCK",
				"entity": "12345",
				"rules": [
					{
						"key": "amount",
						"operator": ">=",
						"value": "1000"
					}
				],
				"tag": "TEST TAG"
			}
		]`,
		mockRuleSets[0].Id,
	)

	assertJSONEqual(t, "Response body expected to be", expected, body)
}
