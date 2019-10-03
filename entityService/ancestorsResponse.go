package entityService

import (
	"errors"
	"github.com/bitly/go-simplejson"
)

func entityIdsFromAncestorsResponseJson(json *simplejson.Json) ([]string, error) {
	var entityIds []string

	currentJson := json

	for currentJson != nil {
		entityId, err := currentJson.Get("entityUid").String()

		if err != nil || entityId == "" {
			return []string{}, errors.New("entityUid is not present")
		}

		entityIds = append(entityIds, entityId)

		parentJson, exists := currentJson.CheckGet("parent")

		if exists {
			currentJson = parentJson
		} else {
			currentJson = nil
		}
	}

	return entityIds, nil
}
