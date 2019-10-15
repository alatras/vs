package entityService

import (
	"errors"
	"github.com/bitly/go-simplejson"
)

func entityIdsFromDescendantsResponseJson(json *simplejson.Json) ([]string, error) {
	var entityIds []string

	entityId, err := json.Get("entityUid").String()

	if err != nil || entityId == "" {
		return []string{}, errors.New("entityUid is not present")
	}

	entityIds = append(entityIds, entityId)

	childrenJson, exists := json.CheckGet("children")

	if exists {
		children, err := childrenJson.Array()

		if err != nil {
			return []string{}, errors.New("children is not array")
		}

		childrenCount := len(children)

		childIndex := 0

		for childIndex = 0; childIndex < childrenCount; childIndex++ {
			childrenEntityIds, err := entityIdsFromDescendantsResponseJson(childrenJson.GetIndex(childIndex))

			if err != nil {
				return []string{}, err
			}

			entityIds = append(entityIds, childrenEntityIds...)
		}
	}

	return entityIds, nil
}
