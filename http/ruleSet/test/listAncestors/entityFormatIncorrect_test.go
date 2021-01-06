package listAncestors

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/listAncestorsRuleSet"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/bitly/go-simplejson"
)

func setupEntityFormatIncorrectErrorRecorder(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(
		log,
		nil,
		nil,
		nil,
		nil,
		func() listAncestorsRuleSet.ListAncestorsRuleSet {
			return &errorApp{error: listAncestorsRuleSet.NewError(listAncestorsRuleSet.EntityIdFormatIncorrectErr, errors.New("incorrect"))}
		},
		nil,
		nil,
	)
	resource.Routes().ServeHTTP(recorder, r)

	return recorder
}

func Test_HTTP_RuleSet_ListAncestors_EntityIdFormatIncorrect(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/ancestors", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := setupEntityFormatIncorrectErrorRecorder(t, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Status code expected to be %d but got %d", http.StatusBadRequest, status)
	}

	body := recorder.Body.String()

	resJson, err := simplejson.NewJson([]byte(body))

	if err != nil {
		t.Errorf("Error while reading response JSON: %s", err)
		return
	}

	errCode := resJson.Get("code").MustInt()
	message := resJson.Get("message").MustString()

	expectedErrCode := 107

	if errCode != expectedErrCode {
		t.Errorf("Expected error code %d but got %d", expectedErrCode, errCode)
		return
	}

	if message != malformedParametersMessage {
		t.Errorf("Expected message %s but got %s", malformedParametersMessage, message)
		return
	}
}
