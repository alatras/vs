package create

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bytes"
	"github.com/bitly/go-simplejson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupInvalidActionRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(
		log,
		nil,
		func() createRuleSet.CreateRuleSet {
			return &errorApp{error: createRuleSet.InvalidAction}
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

func Test_HTTP_RuleSet_Create_InvalidAction(t *testing.T) {
	requestBody :=
		`{
			"name": "test",
			"action": "TEST",
			"rules": [
				{
					"key": "amount",
					"operator": "==",
					"value": "1000"
				}
			]
		}`

	req, err := http.NewRequest("POST", "/12345/rulesets", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	recorder := setupInvalidActionRecorder(t, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Status code expected to be %d but got %d", http.StatusBadRequest, status)
		return
	}

	body := recorder.Body.String()

	resJson, err := simplejson.NewJson([]byte(body))

	if err != nil {
		t.Errorf("Error while reading response JSON: %s", err)
		return
	}

	errCode := resJson.Get("code").MustInt()
	details := resJson.Get("details").MustString()
	message := resJson.Get("message").MustString()

	expectedErrCode := 107

	if errCode != expectedErrCode {
		t.Errorf("Expected error code %d but got %d", expectedErrCode, errCode)
		return
	}

	expectedDetails := "action should be TAG or BLOCK"

	if details != expectedDetails {
		t.Errorf("Expected details %s but got %s", expectedDetails, details)
		return
	}

	if message != malformedParametersErrorMessage {
		t.Errorf("Expected message %s but got %s", malformedParametersErrorMessage, message)
		return
	}
}
