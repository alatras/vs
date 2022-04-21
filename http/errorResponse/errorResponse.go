package errorResponse

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Response struct {
	HttpStatusCode int         `json:"-"`
	Code           int         `json:"code"`
	Details        interface{} `json:"details"`
	Message        string      `json:"message"`
	Timestamp      time.Time   `json:"timestamp"`
}

const (
	unexpectedErrorMessage = "Unexpected error: if the error persists, please contact an " +
		"administrator, quoting the code and timestamp of this error"

	notFoundMessage = "The requested page can't be found. It's likely that the page's URL " +
		"is incorrect, or was accessed using an incorrect protocol. For some pages, a strict URL routing is enabled: " +
		"you may need to add or remove a trailing slash, to or from the URL."

	resourceNotFoundMessage = "The requested resource, or one of its sub-resources, can't be " +
		"found. If the submitted query is valid, this error is likely to be caused by a problem with a nested " +
		"resource that has been deleted or modified. Check the details property for additional insights."

	malformedParametersMessage = "At least one parameter is invalid. Examine the details " +
		"property for more information. Invalid parameters are listed and prefixed accordingly: body for parameters " +
		"submitted in the request's body, query for parameters appended to the request's URL, and params for " +
		"templated parameters of the request's URL."

	forbiddenMessage = "The resource is forbidden. Examine the details property for more information. "

	httpClientMessage = "The requested cannot be processed. Error in calling external service. Examine the details. "

	unauthorizedRequestMessage = "The requested is not authorized. " +
		"Please make sure an auth token is in the request header. Check its validity. "

	EntityServiceUnauthorizeMessage = "Entity Service did not authorize the request. "

	EntityServiceForbidMessage = "Entity Service forbids accessing this resource. "

	EntityServiceUnexpectedErrorMessage = "Entity Service server error. "

	EntityServiceMalformedEntiresMessage = "Malformed data by Entity Service. Check Entity ID and token. "

	EntityServiceNotFoundMessage = "Resource is not found by Entity Service. "

	EntityServiceUnexpectedMessage = "Unexpected error by Entity Service. "

	UnexpectedErrorWithResty = "Unexpected error using Resty to make HTTP call. "
)

func UnexpectedError(details interface{}) *Response {
	return &Response{
		HttpStatusCode: http.StatusInternalServerError,
		Code:           100,
		Details:        details,
		Message:        unexpectedErrorMessage,
		Timestamp:      time.Now(),
	}
}

func NotFound(details interface{}) *Response {
	return &Response{
		HttpStatusCode: http.StatusNotFound,
		Code:           101,
		Details:        details,
		Message:        notFoundMessage,
		Timestamp:      time.Now(),
	}
}

func ResourceNotFound(resource string, id string) *Response {
	return &Response{
		HttpStatusCode: http.StatusNotFound,
		Code:           109,
		Details: map[string]string{
			"resource": resource,
			"id":       id,
		},
		Message:   resourceNotFoundMessage,
		Timestamp: time.Now(),
	}
}

func MalformedParameters(details interface{}) *Response {
	return &Response{
		HttpStatusCode: http.StatusBadRequest,
		Code:           107,
		Details:        details,
		Message:        malformedParametersMessage,
		Timestamp:      time.Now(),
	}
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}

func HttpClientError(id string, description string) *Response {
	return &Response{
		HttpStatusCode: http.StatusInternalServerError,
		Code:           108,
		Details: map[string]string{
			"id":          id,
			"description": description,
		},
		Message:   httpClientMessage,
		Timestamp: time.Now(),
	}
}

func UnauthorizedRequest(id string, description string) *Response {
	return &Response{
		HttpStatusCode: http.StatusUnauthorized,
		Code:           110,
		Details: map[string]string{
			"id":          id,
			"description": description,
		},
		Message:   unauthorizedRequestMessage,
		Timestamp: time.Now(),
	}
}

func ForbiddenRequest(id string, description string) *Response {
	return &Response{
		HttpStatusCode: http.StatusForbidden,
		Code:           111,
		Details: map[string]string{
			"id":          id,
			"description": description,
		},
		Message:   forbiddenMessage,
		Timestamp: time.Now(),
	}
}
