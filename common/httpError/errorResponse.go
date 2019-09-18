package httpError

import (
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type Response struct {
	HttpStatusCode int     `json:"-"`
	Code int	`json:"code"`
	Details	interface{} `json:"details"`
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var unexpectedErrorMessage = "Unexpected error: if the error persists, please contact an " +
"administrator, quoting the code and timestamp of this error"

var notFoundMessage = "The requested page can't be found. It's likely that the page's URL " +
"is incorrect, or was accessed using an incorrect protocol. For some pages, a strict URL routing is enabled: " +
"you may need to add or remove a trailing slash, to or from the URL."

var malformedParametersMessage = "At least one parameter is invalid. Examine the details " +
"property for more information. Invalid parameters are listed and prefixed accordingly: body for parameters " +
"submitted in the request's body, query for parameters appended to the request's URL, and params for " +
"templated parameters of the request's URL."

func UnexpectedError(details interface{}) render.Renderer {
	 return &Response{
		HttpStatusCode: http.StatusInternalServerError,
		Code: 100,
		Details: details,
		Message: unexpectedErrorMessage,
		Timestamp: time.Now(),
	}
}

func NotFound(details interface{}) render.Renderer {
	return &Response{
		HttpStatusCode: http.StatusNotFound,
		Code: 101,
		Details: details,
		Message: notFoundMessage,
		Timestamp: time.Now(),
	}
}

func MalformedParameters(details interface{}) render.Renderer {
	return &Response {
		HttpStatusCode: http.StatusBadRequest,
		Code: 107,
		Details: details,
		Message: malformedParametersMessage,
		Timestamp: time.Now(),
	}
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}