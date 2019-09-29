package create

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const (
	errorColorFormat   = "\033[1;31m%s\033[0m"
	successColorFormat = "\033[0;36m%s\033[0m"
)

func assertJSONEqual(t *testing.T, message, expected, actual string) {
	var jsonExpected, jsonActual interface{}

	if err := json.Unmarshal([]byte(expected), &jsonExpected); err != nil {
		t.Errorf("Failed to decode expected JSON: %v", err)
		return
	}

	if err := json.Unmarshal([]byte(actual), &jsonActual); err != nil {
		t.Errorf("Failed to decode actual JSON: %v", err)
		return
	}

	equal := reflect.DeepEqual(jsonExpected, jsonActual)

	if equal {
		return
	}

	expectedFormattedJSON, err := json.Marshal(jsonExpected)

	if err != nil {
		t.Errorf("Failed to encode: %v", err)
		return
	}

	actualFormattedJSON, err := json.Marshal(jsonActual)

	if err != nil {
		t.Errorf("Failed to encode: %v", err)
		return
	}

	expectedColored := fmt.Sprintf(successColorFormat, expectedFormattedJSON)
	actualColored := fmt.Sprintf(errorColorFormat, actualFormattedJSON)

	t.Errorf("%s:\n%s\nBut got:\n%s", message, expectedColored, actualColored)
}
