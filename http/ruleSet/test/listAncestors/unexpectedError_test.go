package listAncestors

import (
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"bytes"
	"errors"
	"github.com/bitly/go-simplejson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupUnknownErrorRecorder(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(
		log,
		nil,
		nil,
		nil,
		nil,
		func() listAncestorsRuleSet.ListAncestorsRuleSet {
			return &errorApp{error: listAncestorsRuleSet.NewError(listAncestorsRuleSet.UnexpectedErr, errors.New("unexpected"))}
		},
		nil,
		nil,
	)
	resource.Routes().ServeHTTP(recorder, r)

	return recorder
}

func Test_HTTP_RuleSet_ListAncestors_UnexpectedError(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/ancestors", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := setupUnknownErrorRecorder(t, req)

	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("Status code expected to be %d but got %d", http.StatusInternalServerError, status)
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