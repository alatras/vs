package get

import (
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/bitly/go-simplejson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupUnexpectedErrorRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(log, nil, func() getRuleSet.GetRuleSet {
		return &errorApp{error: getRuleSet.UnexpectedError}
	})

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_RuleSet_Get_UnexpectedError(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/"+mockRuleSet.Id, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to get request: %v", err)
		return
	}

	recorder := setupUnexpectedErrorRecorder(t, req)

	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("Status code expected to be %d but got %d", http.StatusInternalServerError, status)
		return
	}

	body := recorder.Body.String()

	resJson, err := simplejson.NewJson([]byte(body))

	if err != nil {
		t.Errorf("Error while reading response JSON: %s", err)
		return
	}

	errCode := resJson.Get("code").MustInt()
	message := resJson.Get("message").MustString()

	expectedErrCode := 100

	if errCode != expectedErrCode {
		t.Errorf("Expected error code %d but got %d", expectedErrCode, errCode)
		return
	}

	if message != unexpectedErrorMessage {
		t.Errorf("Expected message %s but got %s", unexpectedErrorMessage, message)
		return
	}
}