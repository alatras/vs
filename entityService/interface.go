package entityService

import "errors"

var (
	RequestError         = errors.New("request couldn't be performed")
	UnauthorizedError    = errors.New("request was not authorized")
	EntityNotFound       = errors.New("entity was not found")
	ResponseUnsuccessful = errors.New("response status code is not successful")
	ResponseInvalidError = errors.New("response format is invalid")
)

type EntityService interface {
	Ping() error
	GetAncestorsOf(entityId string) ([]string, error)
	GetDescendantsOf(entityId string) ([]string, error)
}
