package listDescendantsRuleSet

import "fmt"

type kind uint8

const (
	UnexpectedErr kind = iota
	EntityIdNotFoundErr
	EntityIdFormatIncorrectErr
)

type AppError struct {
	kind  kind
	error error
}

func NewError(k kind, e error) AppError {
	return AppError{
		kind:  k,
		error: e,
	}
}

func (e AppError) Error() string {
	switch e.kind {
	case EntityIdNotFoundErr:
		return "entity with provided id was not found"
	case EntityIdFormatIncorrectErr:
		return "entity id should be a valid UUID"
	default:
		return fmt.Sprintf("an unexpected error occurred while listing descendants rule sets: %s", e.error)
	}
}

func (e AppError) Is(k kind) bool {
	return e.kind == k
}

func (e AppError) HasError() bool {
	return e.error != nil
}
