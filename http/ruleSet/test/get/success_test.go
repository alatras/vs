package get

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/getRuleSet"
	"validation-service/http/ruleSet"
	"validation-service/logger"
)

func setupSuccessRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(
		log,
		nil,
		func() getRuleSet.GetRuleSet {
			return &successApp{}
		},
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_RuleSet_Get_Success(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/"+mockRuleSet.Id, nil)

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupSuccessRecorder(t, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d but got %d", http.StatusOK, status)
	}

	body := recorder.Body.String()

	expected := fmt.Sprintf(
		`{
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
			]
		}`,
		mockRuleSet.Id,
	)

	assertJSONEqual(t, "Response body expected to be", expected, body)
}
