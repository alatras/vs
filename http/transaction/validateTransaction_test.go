package transaction

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"validation-service/app/validateTransaction"
	"validation-service/logger"
	"validation-service/report"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"

	"github.com/bitly/go-simplejson"
)

func setupRecorder(t *testing.T, request *http.Request, report *report.Report, err *validateTransaction.AppError) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()

	log := logger.NewStubLogger()

	successApp := app{
		err: err,
		rep: report,
	}

	resource := NewResource(
		log,
		&successApp,
	)

	resource.Routes().ServeHTTP(recorder, request)

	return recorder
}

func Test_HTTP_ValidateTransaction_Success_Pass(t *testing.T) {
	requestBody :=
		`{
			"transaction": {
				"amount": {
					"value": "100",
					"currencyCode": "EUR"
				},
				"merchant": {
					"organisation": {
						"UUID": "123"
					}
				},
				"customer": {
					"country": "NL"
				}
			}
		}`

	report := report.Report{
		Action:          report.Pass,
		BlockedRuleSets: []ruleSet.RuleSet{},
		TaggedRuleSets:  []ruleSet.RuleSet{},
	}

	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupRecorder(t, req, &report, nil)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d but got %d", http.StatusOK, status)
	}

	body := recorder.Body.String()

	expected := fmt.Sprintf(
		`{
			"action": "PASS",
			"block": [],
			"tags": []
		}`,
	)

	assertJSONEqual(t, "Response body expected to be", expected, body)
}

func Test_HTTP_ValidateTransaction_Success_Block(t *testing.T) {
	requestBody :=
		`{
			"transaction": {
				"amount": {
					"value": "100",
					"currencyCode": "EUR"
				},
				"merchant": {
					"organisation": {
						"UUID": "123"
					}
				},
				"customer": {
					"country": "NL"
				}
			}
		}`

	report := report.Report{
		Action: report.Block,
		BlockedRuleSets: []ruleSet.RuleSet{
			{
				Id:       "234",
				EntityId: "123",
				Action:   "BLOCK",
				Name:     "Test",
				RuleMetadata: []rule.Metadata{
					{
						Property: "amount",
						Operator: ">",
						Value:    "50",
					},
				},
			},
		},
		TaggedRuleSets: []ruleSet.RuleSet{},
	}

	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupRecorder(t, req, &report, nil)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Status code expected to be %d but got %d", http.StatusOK, status)
	}

	body := recorder.Body.String()

	expected := fmt.Sprintf(
		`{
			"action": "BLOCK",
			"block": [
				{
					"id": "234",
					"entity": "123",
					"action": "BLOCK",
					"name": "Test",
					"rules": [
						{
							"key": "amount",
							"operator": ">",
							"value": "50"
						}
					]
				}
			],
			"tags": []
		}`,
	)

	assertJSONEqual(t, "Response body expected to be", expected, body)
}

func Test_HTTP_ValidateTransaction_EntityNotFound(t *testing.T) {
	requestBody :=
		`{
			"transaction": {
				"amount": {
					"value": "100",
					"currencyCode": "EUR"
				},
				"merchant": {
					"organisation": {
						"UUID": "123"
					}
				},
				"customer": {
					"country": "NL"
				}
			}
		}`

	validationError := validateTransaction.NewError(validateTransaction.EntityIdNotFoundErr, errors.New("not found"))

	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupRecorder(t, req, nil, &validationError)

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

	if message != notFoundMessage {
		t.Errorf("Expected message %s but got %s", notFoundMessage, message)
		return
	}
}

func Test_HTTP_ValidateTransaction_EntityIdFormatIncorrect(t *testing.T) {
	requestBody :=
		`{
			"transaction": {
				"amount": {
					"value": "100",
					"currencyCode": "EUR"
				},
				"merchant": {
					"organisation": {
						"UUID": "123"
					}
				},
				"customer": {
					"country": "NL"
				}
			}
		}`

	validationError := validateTransaction.NewError(validateTransaction.EntityIdFormatIncorrectErr, errors.New("incorrect"))

	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupRecorder(t, req, nil, &validationError)

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

func Test_HTTP_ValidateTransaction_UnexpectedError(t *testing.T) {
	requestBody :=
		`{
			"transaction": {
				"amount": {
					"value": "100",
					"currencyCode": "EUR"
				},
				"merchant": {
					"organisation": {
						"UUID": "123"
					}
				},
				"customer": {
					"country": "NL"
				}
			}
		}`

	validationError := validateTransaction.NewError(validateTransaction.UnexpectedErr, errors.New("unexpected"))

	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	recorder := setupRecorder(t, req, nil, &validationError)

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
