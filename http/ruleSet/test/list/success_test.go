package list

import (
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupSuccessRecorder(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(log, nil, nil, nil,
		func() listRuleSet.ListRuleSet {
			return &successApp{}
		}, nil)

	resource.Routes().ServeHTTP(recorder, r)

	return recorder
}

func Test_HTTP_RuleSet_List_Success(t *testing.T) {

	req, err := http.NewRequest("GET", "/12345/rulesets", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

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
				]
			}
		]`,
		mockRuleSets[0].Id,
	)

	assertJSONEqual(t, "Response body expected to be", expected, body)
}
