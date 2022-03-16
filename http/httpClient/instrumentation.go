package httpClient

import (
	"time"
	"validation-service/logger"
)

type metadata = logger.Metadata

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
	record    *logger.LogRecord
}

func newInstrumentation(log *logger.Logger, record *logger.LogRecord) *instrumentation {
	return &instrumentation{
		logger: log,
		record: record.NewRecord().Scoped("CheckEntity"),
	}
}

func (i *instrumentation) setMetadata(metadata metadata) {
	i.record = i.record.Metadata(metadata)
}

func (i *instrumentation) setTraceId(id string) {
	i.record.Mdc.TraceId = id
}

func (i *instrumentation) finishHttpRequest(status int, err error) {
	duration := time.Since(i.startedAt)

	i.record.Duration(int(duration)).Data(
		map[string]interface{}{"status": status, "error": err},
	).MessageObject("Finished HTTP request", "")
	i.doLog("")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
