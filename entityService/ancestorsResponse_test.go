package entityService

import (
	"github.com/bitly/go-simplejson"
	"reflect"
	"testing"
)

func Test_EntityService_AncestorsResponse_Success(t *testing.T) {
	responseText := `{
		"parent": {
			"parent": {
				"name": "string222",
				"entityUid": "f04d7b91-acfe-4e08-8e96-67b493192baa"
			},
			"name": "string333",
			"entityUid": "54259bdc-f31c-430d-89e0-d2959f9b013a"
		},
		"name": "string444",
		"entityUid": "558b821e-3549-4a2f-a575-f14815053f2b"
	}`

	expectedEntityIds := []string{
		"558b821e-3549-4a2f-a575-f14815053f2b",
		"54259bdc-f31c-430d-89e0-d2959f9b013a",
		"f04d7b91-acfe-4e08-8e96-67b493192baa",
	}

	json, err := simplejson.NewJson([]byte(responseText))

	if err != nil {
		t.Errorf("failed to decode json: %v", err)
		return
	}

	entityIds, err := entityIdsFromAncestorsResponseJson(json)

	if err != nil {
		t.Errorf("failed to get entity ids from the response: %v", err)
		return
	}

	if !reflect.DeepEqual(entityIds, expectedEntityIds) {
		t.Errorf("expected entity ids to be\n%v\nbut got\n%v", expectedEntityIds, entityIds)
		return
	}
}

func Test_EntityService_AncestorsResponse_Success_NoParents(t *testing.T) {
	responseText := `{
		"name": "string444",
		"entityUid": "558b821e-3549-4a2f-a575-f14815053f2b"
	}`

	expectedEntityIds := []string{
		"558b821e-3549-4a2f-a575-f14815053f2b",
	}

	json, err := simplejson.NewJson([]byte(responseText))

	if err != nil {
		t.Errorf("failed to decode json: %v", err)
		return
	}

	entityIds, err := entityIdsFromAncestorsResponseJson(json)

	if err != nil {
		t.Errorf("failed to get entity ids from the response: %v", err)
		return
	}

	if !reflect.DeepEqual(entityIds, expectedEntityIds) {
		t.Errorf("expected entity ids to be\n%v\nbut got\n%v", expectedEntityIds, entityIds)
		return
	}
}

func Test_EntityService_AncestorsResponse_Fail_NoEntityUid(t *testing.T) {
	responseText := `{
		"parent": {
			"parent": {
				"name": "string222"
			},
			"name": "string333",
			"entityUid": "54259bdc-f31c-430d-89e0-d2959f9b013a"
		},
		"name": "string444",
		"entityUid": "558b821e-3549-4a2f-a575-f14815053f2b"
	}`

	json, err := simplejson.NewJson([]byte(responseText))

	if err != nil {
		t.Errorf("failed to decode json: %v", err)
		return
	}

	_, err = entityIdsFromAncestorsResponseJson(json)

	if err == nil {
		t.Errorf("expected to fail with error but was successful")
		return
	}
}
