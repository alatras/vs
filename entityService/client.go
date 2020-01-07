package entityService

import (
	"bitbucket.verifone.com/validation-service/logger"
	"crypto/tls"
	"github.com/bitly/go-simplejson"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type client struct {
	logger     *logger.Logger
	httpClient *http.Client
	url        string
	jwtToken   string
}

func NewClient(logger *logger.Logger, url string, jwtToken string, skipCertificateVerification bool) *client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipCertificateVerification,
		},
	}

	httpClient := &http.Client{Transport: tr}

	return &client{
		logger:     logger.Scoped("EntityServiceClient"),
		httpClient: httpClient,
		url:        url + "/ds-entity-service",
		jwtToken:   jwtToken,
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

	req, err := http.NewRequest("GET", c.url+"/entities/"+entityId+"/ancestors", nil)

	if err != nil {
		errorLog.WithError(err).Error("failed to create the request")
		return []string{}, RequestError
	}

	req.Header.Set("Authorization", "Bearer "+c.jwtToken)

	resp, err := c.httpClient.Do(req)

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

	if resp.StatusCode == 400 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("entity id is not correct")
		return []string{}, EntityIdFormatIncorrect
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

	req, err := http.NewRequest("GET", c.url+"/entities/"+entityId+"/descendants", nil)

	if err != nil {
		errorLog.WithError(err).Error("failed to create the request")
		return []string{}, RequestError
	}

	req.Header.Set("Authorization", "Bearer "+c.jwtToken)

	resp, err := c.httpClient.Do(req)

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

	if resp.StatusCode == 400 {
		errorLog.WithField("errorDetails", json.MustMap()).Error("entity id is not correct")
		return []string{}, EntityIdFormatIncorrect
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
