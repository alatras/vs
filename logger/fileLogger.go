package logger

// Field for one filed in log
type Field struct {
	Key   string
	Value interface{}
}

// ErrorDetails of error in log
type ErrorDetails struct {
	ReasonCode   string
	ReasonStatus string
}

// Error for log
type LogError struct {
	Message string
	Name    string
	Stack   string
	Code    int
	Details ErrorDetails
}
