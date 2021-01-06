package listDescendants

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/listDescendantsRuleSet"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/bitly/go-simplejson"
)

func setupEntityNotFoundErrorRecorder(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(
		log,
		nil,
		nil,
		nil,
		nil,
		nil,
		func() listDescendantsRuleSet.ListDescendantsRuleSet {
			return &errorApp{error: listDescendantsRuleSet.NewError(listDescendantsRuleSet.EntityIdNotFoundErr, errors.New("not found"))}
		},
		nil,
	)
	resource.Routes().ServeHTTP(recorder, r)

	return recorder
}

func Test_HTTP_RuleSet_ListDescendants_EntityNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/descendants", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := setupEntityNotFoundErrorRecorder(t, req)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("Status code expected to be %d but got %d", http.StatusNotFound, status)
	}

	body := recorder.Body.String()

	resJson, err := simplejson.NewJson([]byte(body))

	if err != nil {
		t.Errorf("Error while reading response JSON: %s", err)
		return
	}

	errCode := resJson.Get("code").MustInt()
	message := resJson.Get("message").MustString()

	expectedErrCode := 109

	if errCode != expectedErrCode {
		t.Errorf("Expected error code %d but got %d", expectedErrCode, errCode)
		return
	}

	if message != resourceNotFoundMessage {
		t.Errorf("Expected message %s but got %s", resourceNotFoundMessage, message)
		return
	}
}
