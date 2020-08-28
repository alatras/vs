package contextKey

type ContextKey string

const (
	TraceId             ContextKey = "traceId"
	BusinessTransaction ContextKey = "businessTransaction"
)
