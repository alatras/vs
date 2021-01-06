package get

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/getRuleSet"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/bitly/go-simplejson"
)

func setupNotFoundErrorRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	resource := ruleSet.NewResource(
		log,
		nil,
		func() getRuleSet.GetRuleSet {
			return &errorApp{error: getRuleSet.NotFound}
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

func Test_HTTP_RuleSet_Get_NotFoundError(t *testing.T) {
	req, err := http.NewRequest("GET", "/12345/rulesets/"+mockRuleSet.Id, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	recorder := setupNotFoundErrorRecorder(t, req)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("Status code expected to be %d but got %d", http.StatusNotFound, status)
		return
	}

	body := recorder.Body.String()

	resJson, err := simplejson.NewJson([]byte(body))

	if err != nil {
		t.Errorf("Error while reading response JSON: %s", err)
		return
	}

	errCode := resJson.Get("code").MustInt()
	detailsResource := resJson.Get("details").Get("resource").MustString()
	detailsId := resJson.Get("details").Get("id").MustString()
	message := resJson.Get("message").MustString()

	expectedErrCode := 109

	if errCode != expectedErrCode {
		t.Errorf("Expected error code %d but got %d", expectedErrCode, errCode)
		return
	}

	if detailsResource != "ruleSet" {
		t.Errorf("Expected details resource %s but got %s", "ruleSet", detailsResource)
		return
	}

	if detailsId != mockRuleSet.Id {
		t.Errorf("Expected details id %s but got %s", mockRuleSet.Id, detailsId)
		return
	}

	if message != notFoundErrorMessage {
		t.Errorf("Expected message %s but got %s", notFoundErrorMessage, message)
		return
	}
}
