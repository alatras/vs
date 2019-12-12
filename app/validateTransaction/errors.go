package validateTransaction

import "fmt"

type kind uint8

const (
	UnexpectedErr kind = iota
	EntityIdNotFoundErr
)

type ValidationError struct {
	kind  kind
	error error
}

func NewError(k kind, e error) ValidationError {
	return ValidationError{
		kind:  k,
		error: e,
	}
}

func (e ValidationError) Error() string {
	switch e.kind {
	case EntityIdNotFoundErr:
		return "invalid entity id in transaction object"
	default:
		return fmt.Sprintf("an unexpected error occurred while validating transaction: %s", e.error)
	}
}

func (e ValidationError) Is(k kind) bool {
	return e.kind == k
}
