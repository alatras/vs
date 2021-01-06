package tests

import (
	"reflect"
	"testing"
	"validation-service/ruleSet"
)

func AssertRuleSet(t *testing.T, expected, actual ruleSet.RuleSet) {
	if actual.EntityId != expected.EntityId {
		t.Errorf("Expected entity id to be %s but got %s", expected.EntityId, actual.EntityId)
		return
	}

	if actual.Name != expected.Name {
		t.Errorf("Expected name to be %s but got %s", expected.Name, actual.Name)
		return
	}

	if actual.Action != expected.Action {
		t.Errorf("Expected action to be %s but got %s", expected.Action, actual.Action)
		return
	}

	if !reflect.DeepEqual(actual.RuleMetadata, expected.RuleMetadata) {
		t.Errorf("Expected rules to be %v but got %v", expected.RuleMetadata, actual.RuleMetadata)
		return
	}
}
