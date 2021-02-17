package logger

import (
	"time"
)

type Field struct {
	Key   string
	Value interface{}
}

type FileLogger struct {
	TraceID   string
	Scope     string
	Metadata  Metadata
	Error     error
	Fields    []Field
	Message   string
	CreatedAt time.Time
	StartedAt time.Time
}
