package ruleset

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func assertJSONEqual(a, b string) (bool, error) {
	var jsonA, jsonB interface{}

	if err := json.Unmarshal([]byte(a), &jsonA); err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(b), &jsonB); err != nil {
		return false, err
	}

	return reflect.DeepEqual(jsonB, jsonA), nil
}

func Test_RuleSet_CreateSuccess(t *testing.T) {
	body := `{
		"name": "test",
		"action": "BLOCK",
		"rules": [
			{
				"key": "amount",
				"operator": "==",
				"value": "1000"
			}
		]
	}`

	req, err := http.NewRequest("POST", "/12345/rulesets", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	log := logger.NewStubLogger()
	ruleSetRepo, err := ruleSet.NewStubRepository()

	if err != nil {
		t.Errorf("Failed to create rule set repository: %v", err)
	}

	resource := NewResource(log, ruleSetRepo)

	resource.Routes().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d.\n Got %d", http.StatusOK, status)
	}

	actual := rr.Body.String()

	resJson, err := simplejson.NewJson([]byte(actual))

	if err != nil {
		t.Errorf("Error while reading response JSON: %s", err)
	}

	createdId := resJson.Get("id").MustString()

	expected := fmt.Sprintf(`{
		"id": "%s",
		"name": "test",
		"action": "BLOCK",
		"entity": "12345",
		"rules": [
			{
				"key": "amount",
				"operator": "==",
				"value": "1000"
			}
		]
	}`, createdId)

	equal, err := assertJSONEqual(actual, expected)

	if err != nil {
		t.Errorf("Body failed to decode: %v", err)
	}

	if !equal {
		t.Errorf("Body expected to be %s.\n Got %s", expected, actual)
	}
}
