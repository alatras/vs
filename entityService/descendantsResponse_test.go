package entityService

import (
	"github.com/bitly/go-simplejson"
	"reflect"
	"testing"
)

func Test_DescendantsResponse_Success(t *testing.T) {
	responseText := `{
		"children": [
			{
				"entityUid": "rootChild1",
				"name": "rootChild1",
				"children": [
					{
						"entityUid": "rootChild1Child1",
						"name": "rootChild1Child1"
					},
					{
						"entityUid": "rootChild1Child2",
						"name": "rootChild1Child2"
					}
				]
			},
			{
				"entityUid": "rootChild2",
				"name": "rootChild2",
				"children": [
					{
						"entityUid": "rootChild2Child1",
						"name": "rootChild2Child1"
					},
					{
						"entityUid": "rootChild2Child2",
						"name": "rootChild2Child2"
					}
				]
			}
		],
		"entityUid": "root",
		"name": "root"
	}`

	expectedEntityIds := []string{
		"root",
		"rootChild1",
		"rootChild1Child1",
		"rootChild1Child2",
		"rootChild2",
		"rootChild2Child1",
		"rootChild2Child2",
	}

	json, err := simplejson.NewJson([]byte(responseText))

	if err != nil {
		t.Errorf("failed to decode json: %v", err)
		return
	}

	entityIds, err := entityIdsFromDescendantsResponseJson(json)

	if err != nil {
		t.Errorf("failed to get entity ids from the response: %v", err)
		return
	}

	if !reflect.DeepEqual(entityIds, expectedEntityIds) {
		t.Errorf("expected entity ids to be\n%v\nbut got\n%v", expectedEntityIds, entityIds)
		return
	}
}

func Test_DescendantsResponse_Success_NoChildren(t *testing.T) {
	responseText := `{	
		"entityUid": "root",
		"name": "root"
	}`

	expectedEntityIds := []string{
		"root",
	}

	json, err := simplejson.NewJson([]byte(responseText))

	if err != nil {
		t.Errorf("failed to decode json: %v", err)
		return
	}

	entityIds, err := entityIdsFromDescendantsResponseJson(json)

	if err != nil {
		t.Errorf("failed to get entity ids from the response: %v", err)
		return
	}

	if !reflect.DeepEqual(entityIds, expectedEntityIds) {
		t.Errorf("expected entity ids to be\n%v\nbut got\n%v", expectedEntityIds, entityIds)
		return
	}
}

func Test_EntityService_DescendantsResponse_Fail_NoentityUid(t *testing.T) {
	responseText := `{
		"children": [
			{
				"entityUid": "rootChild1",
				"name": "rootChild1",
				"children": [
					{
						"entityUid": "rootChild1Child1",
						"name": "rootChild1Child1"
					},
					{
						"name": "rootChild1Child2"
					}
				]
			},
			{
				"entityUid": "rootChild2",
				"name": "rootChild2",
				"children": [
					{
						"entityUid": "rootChild2Child1",
						"name": "rootChild2Child1"
					},
					{
						"entityUid": "rootChild2Child2",
						"name": "rootChild2Child2"
					}
				]
			}
		],
		"entityUid": "root",
		"name": "root"
	}`

	json, err := simplejson.NewJson([]byte(responseText))

	if err != nil {
		t.Errorf("failed to decode json: %v", err)
		return
	}

	_, err = entityIdsFromDescendantsResponseJson(json)

	if err == nil {
		t.Errorf("expected to fail with error but was successful")
		return
	}
}
