package log

import (
	"fmt"
	"os"
	"time"
)

// assert interface compliance.
var _ Interface = (*Entry)(nil)

// Entry represents a single log entry.
type Entry struct {
	Logger    *Logger   `json:"-"`
	Fields    Fields    `json:"fields"`
	Level     Level     `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	start     time.Time
	fields    []Fields
}

// NewEntry returns a new entry for `log`.
func NewEntry(log *Logger) *Entry {
	return &Entry{
		Logger: log,
	}
}

// WithFields returns a new entry with `fields` set.
func (e *Entry) WithFields(fields Fielder) *Entry {
	f := []Fields{}
	f = append(f, e.fields...)
	f = append(f, fields.Fields())
	return &Entry{
		Logger: e.Logger,
		fields: f,
	}
}

// WithField returns a new entry with the `key` and `value` set.
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(Fields{key: value})
}

// WithError returns a new entry with the "error" set to `err`.
func (e *Entry) WithError(err error) *Entry {
	if err == nil {
		return e.WithField("error", err)
	}
	return e.WithField("error", err.Error())
}

func (e *Entry) WithMemStats() *Entry {
	return e.WithField("memstats", Fn(getMemStats))
}

// Debug level message.
func (e *Entry) Debug(v ...interface{}) {
	e.Logger.log(DebugLevel, e, sprintlnn(v...))
}

// Info level message.
func (e *Entry) Info(v ...interface{}) {
	e.Logger.log(InfoLevel, e, sprintlnn(v...))
}

// Warn level message.
func (e *Entry) Warn(v ...interface{}) {
	e.Logger.log(WarnLevel, e, sprintlnn(v...))
}

// Error level message.
func (e *Entry) Error(v ...interface{}) {
	e.Logger.log(ErrorLevel, e, sprintlnn(v...))
}

// Fatal level message, followed by an exit.
func (e *Entry) Fatal(v ...interface{}) {
	e.Logger.log(FatalLevel, e, sprintlnn(v...))
	os.Exit(1)
}

// Debugf level formatted message.
func (e *Entry) Debugf(msg string, v ...interface{}) {
	e.Debug(fmt.Sprintf(msg, v...))
}

// Infof level formatted message.
func (e *Entry) Infof(msg string, v ...interface{}) {
	e.Info(fmt.Sprintf(msg, v...))
}

// Warnf level formatted message.
func (e *Entry) Warnf(msg string, v ...interface{}) {
	e.Warn(fmt.Sprintf(msg, v...))
}

// Errorf level formatted message.
func (e *Entry) Errorf(msg string, v ...interface{}) {
	e.Error(fmt.Sprintf(msg, v...))
}

// Fatalf level formatted message, followed by an exit.
func (e *Entry) Fatalf(msg string, v ...interface{}) {
	e.Fatal(fmt.Sprintf(msg, v...))
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func (e *Entry) Trace(msg string) *Entry {
	e.Info(msg)
	v := e.WithFields(e.Fields)
	v.Message = msg
	v.start = time.Now()
	return v
}

// Stop should be used with Trace, to fire off the completion message. When
// an `err` is passed the "error" field is set, and the log level is error.
func (e *Entry) Stop(err *error) {
	if err == nil || *err == nil {
		e.WithField("duration", time.Since(e.start)).Info(e.Message)
	} else {
		e.WithField("duration", time.Since(e.start)).WithError(*err).Error(e.Message)
	}
}

// mergedFields returns the fields list collapsed into a single map and
// resolves all Fn types.
func (e *Entry) mergedFields() Fields {
	f := Fields{}

	for _, fields := range e.fields {
		for k, v := range fields {
			switch v := v.(type) {
			case error:
				f[k] = v.Error()
			case Fn:
				f[k] = v()
			default:
				f[k] = v
			}
		}
	}

	return f
}

// finalize returns a copy of the Entry with Fields merged.
func (e *Entry) finalize(level Level, msg string) *Entry {
	return &Entry{
		Logger:    e.Logger,
		Fields:    e.mergedFields(),
		Level:     level,
		Message:   msg,
		Timestamp: time.Now(),
	}
}
