package httpClient

import (
	"strconv"
	"time"
	"validation-service/config"
)

func (c Client) CheckEntity(
	token string,
	entityId string,
	traceID string,
) int {
	baseUrl := c.config.EntityService.URL

	c.instrumentation.setMetadata(map[string]interface{}{
		"Entity id":          entityId,
		"Entity Service URL": baseUrl,
	})
	c.instrumentation.setTraceId(traceID)

	url := baseUrl + "/entities/" + entityId

	timeoutInt, err := strconv.Atoi(config.App.EntityService.Timeout)
	if err != nil {
		timeoutInt = 20
	}

	resp, err := c.restyClient.
		SetTimeout(time.Duration(timeoutInt) * time.Second).R().
		SetAuthToken(token).Get(url)

	c.instrumentation.finishHttpRequest(resp.StatusCode(), err)

	if err != nil {
		return 512
	}

	return resp.StatusCode()
}
