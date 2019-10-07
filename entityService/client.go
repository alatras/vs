package entityService

import (
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/bitly/go-simplejson"
	"github.com/sirupsen/logrus"
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
	errorLog := c.logger.Error.WithFields(logrus.Fields{
		"entityId": entityId,
		"method":   "GetAncestorsOf",
	})

	resp, err := http.Get(c.url + "/entities/" + entityId + "/ancestors")

	if err != nil {
		errorLog.WithError(err).Error("failed to perform the request")
		return []string{}, RequestError
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errorLog.WithError(err).Error("failed to read the response body")
		return []string{}, RequestError
	}

	json, err := simplejson.NewJson(body)

	if err != nil {
		errorLog.WithError(err).Error("failed to parse the response body")
		return []string{}, ResponseInvalidError
	}

	if resp.StatusCode == 401 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("request unauthorized")
		return []string{}, UnauthorizedError
	}

	if resp.StatusCode == 404 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("entity was not found")
		return []string{}, EntityNotFound
	}

	if resp.StatusCode != 200 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("response status code is unsuccessful")
		return []string{}, ResponseUnsuccessful
	}

	entityIds, err := entityIdsFromAncestorsResponseJson(json)

	if err != nil {
		errorLog.WithError(err).Error("failed to get entity ids from json")
		return []string{}, ResponseInvalidError
	}

	return entityIds, nil
}

func (c *client) GetDescendantsOf(entityId string) ([]string, error) {
	errorLog := c.logger.Error.WithFields(logrus.Fields{
		"entityId": entityId,
		"method":   "GetDescendantsOf",
	})

	resp, err := http.Get(c.url + "/entities/" + entityId + "/descendants")

	if err != nil {
		errorLog.WithError(err).Error("failed to perform the request")
		return []string{}, RequestError
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errorLog.WithError(err).Error("failed to read the response body")
		return []string{}, RequestError
	}

	json, err := simplejson.NewJson(body)

	if err != nil {
		errorLog.WithError(err).Error("failed to parse the response body")
		return []string{}, ResponseInvalidError
	}

	if resp.StatusCode == 401 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("request unauthorized")
		return []string{}, UnauthorizedError
	}

	if resp.StatusCode == 404 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("entity was not found")
		return []string{}, EntityNotFound
	}

	if resp.StatusCode != 200 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("response status code is unsuccessful")
		return []string{}, ResponseUnsuccessful
	}

	entityIds, err := entityIdsFromDescendantsResponseJson(json)

	if err != nil {
		errorLog.WithError(err).Error("failed to get entity ids from json")
		return []string{}, err
	}

	return entityIds, nil
}
