package update

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"validation-service/app/updateRuleSet"
	"validation-service/config"
	"validation-service/http/httpClient"
	"validation-service/http/ruleSet"
	"validation-service/logger"

	"github.com/bitly/go-simplejson"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func setupNotFoundErrorRecorder(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
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
		nil,
		func() updateRuleSet.UpdateRuleSet {
			return &errorApp{error: updateRuleSet.NotFound}
		},
	)

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_RuleSet_Update_NotFoundError(t *testing.T) {
	requestBody :=
		`{
			"name": "Test",
			"action": "BLOCK",
			"rules": [
				{
					"key": "amount",
					"operator": ">=",
					"value": "1000"
				}
			]
		}`

	req, err := http.NewRequest("PUT", "/12345/rulesets/"+mockRuleSet.Id, bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

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
