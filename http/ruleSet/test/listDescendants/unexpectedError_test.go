package listDescendants

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/listDescendantsRuleSet"
	"validation-service/config"
	"validation-service/http/httpClient"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/bitly/go-simplejson"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func setupUnknownErrorRecorder(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
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
			return &errorApp{error: listDescendantsRuleSet.NewError(listDescendantsRuleSet.UnexpectedErr, errors.New("unexpected"))}
		},
		nil,
	)
	resource.Routes().ServeHTTP(recorder, r)

	return recorder
}

func Test_HTTP_RuleSet_ListDescendants_UnexpectedError(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/descendants", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

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
