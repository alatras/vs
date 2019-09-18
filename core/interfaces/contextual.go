package interfaces

import "context"

type Contextual interface {
	SetContext(ctx context.Context)
}