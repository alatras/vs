package logger

type MDC struct {
	TraceId       string `json:"trace_id"`
	CorrelationId string `json:"correlation_id"`
	AppName       string `json:"app_name"`
	AppVersion    string `json:"app_version"`
	Delay         int    `json:"delay"`
	Metadata      Generic
	Scope         string                 `json:"scope"`
	Duration      int                    `json:"duration"`
	Data          map[string]interface{} `json:"ruleSet"`
}

type Generic interface{}

type Message struct {
	Trace_message string  `json:"trace_message"`
	Trace_details Generic `json:"trace_details"`
}

type OperationMessage struct {
	SOURCE_TARGET_OP_elapsed_time      int               // out_time - in_time
	SOURCE_TARGET_OP_in_time           int               // unix timestamp
	SOURCE_TARGET_OP_out_time          int               // unix timestamp
	SOURCE_TARGET_OP_response_code     string            // HTTP STATUS CODE
	SOURCE_TARGET_OP_error_codestring  string            // error code
	SOURCE_TARGET_OP_error_description string            // thrown exception
	SOURCE_TARGET_OP_in_parameters     map[string]string // sent payload
	SOURCE_TARGET_OP_out_parameters    interface{}       // received payload
	SOURCE_TARGET_OP_URL               string            // URL
}

type LogRecord struct {
	Mdc     MDC
	Message Message
}

type Exception struct {
	ExceptionClass   string      `json:"exception_class"`
	Stacktrace       string      `json:"stacktrace"`
	ExceptionMessage interface{} `json:"exception_message"`
}

func (l *LogRecord) NewRecord() *LogRecord {
	r := new(LogRecord)
	r.Mdc.AppName = AppName
	r.Mdc.AppVersion = AppVersion
	return r
}

func (l *LogRecord) Scoped(scope string) *LogRecord {
	l.Mdc.Scope = scope
	return l
}

func (l *LogRecord) TraceId(traceId string) *LogRecord {
	l.Mdc.TraceId = traceId
	return l
}

func (l *LogRecord) Metadata(metadata Generic) *LogRecord {
	l.Mdc.Metadata = metadata
	return l
}

func (l *LogRecord) MessageObject(
	traceMessage string,
	traceDetails Generic,
) *LogRecord {
	l.Message.Trace_message = traceMessage
	l.Message.Trace_details = traceDetails
	return l
}

func (l *LogRecord) CorrelationId(correlationId string) *LogRecord {
	l.Mdc.CorrelationId = correlationId
	return l
}

func (l *LogRecord) Duration(duration int) *LogRecord {
	l.Mdc.Duration = duration
	return l
}

func (l *LogRecord) Delay(delay int) *LogRecord {
	l.Mdc.Duration = delay
	return l
}

func (l *LogRecord) Data(data map[string]interface{}) *LogRecord {
	l.Mdc.Data = data
	return l
}
