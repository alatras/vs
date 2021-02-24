package logger

import (
	"time"
)

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

// FileLogger for extra log
type FileLogRecord struct {
	Time     time.Time
	TraceID  string
	Name     string // service name
	Hostname string
	PID      int
	Tag      string
	Level    string
	Scope    string
	Metadata Metadata
	Error    LogError
	Fields   []Field
	Message  string
}
