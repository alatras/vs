package entityService

import (
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
)

type client struct {
	logger *logger.Logger
	url    string
}

func NewClient(logger *logger.Logger, url string) *client {
	return &client{
		logger: logger.Scoped("EntityServiceClient"),
		url:    url + "/ds-entity-service",
	}
}

func (c *client) Ping() error {
	return nil
}

func (c *client) GetAncestorsOf(entityId string) ([]string, error) {
	errorLog := c.logger.Error.WithField("entityId", entityId)

	resp, err := http.Get(c.url + "/entities/" + entityId + "/ancestors")

	if err != nil {
		errorLog.WithError(err).Error("failed to perform the request")
		return []string{}, RequestError
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errorLog.WithError(err).Error("failed to read the request body")
		return []string{}, RequestError
	}

	json, err := simplejson.NewJson(body)

	if err != nil {
		errorLog.WithError(err).Error("failed to parse the request body")
		return []string{}, err
	}

	entityIds, err := entityIdsFromAncestorsResponseJson(json)

	return entityIds, nil
}

func (c *client) GetDescendantsOf(entityId string) ([]string, error) {
	return []string{}, nil
}
