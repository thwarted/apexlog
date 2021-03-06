package log

import (
	"os"
)

// singletons ftw?
var Log Interface = &Logger{
	Handler: NewSimple(os.Stderr),
	Level:   InfoLevel,
}

// SetHandler sets the handler. This is not thread-safe.
func SetHandler(h Handler) {
	if logger, ok := Log.(*Logger); ok {
		logger.Handler = h
	}
}

// SetLevel sets the log level. This is not thread-safe.
func SetLevel(l Level) {
	if logger, ok := Log.(*Logger); ok {
		logger.Level = l
	}
}

// WithFields returns a new entry with `fields` set.
func WithFields(fields Fielder) *Entry {
	return Log.WithFields(fields)
}

// WithField returns a new entry with the `key` and `value` set.
func WithField(key string, value interface{}) *Entry {
	return Log.WithField(key, value)
}

// WithError returns a new entry with the "error" set to `err`.
func WithError(err error) *Entry {
	return Log.WithError(err)
}

// WithError returns a new entry with the "error" set to `err`.
func WithMemStats() *Entry {
	return Log.WithMemStats()
}

// Debug level message.
func Debug(v ...interface{}) {
	Log.Debug(v...)
}

// Info level message.
func Info(v ...interface{}) {
	Log.Info(v...)
}

// Warn level message.
func Warn(v ...interface{}) {
	Log.Warn(v...)
}

// Error level message.
func Error(v ...interface{}) {
	Log.Error(v...)
}

// Fatal level message, followed by an exit.
func Fatal(v ...interface{}) {
	Log.Fatal(v...)
}

// Debugf level formatted message.
func Debugf(msg string, v ...interface{}) {
	Log.Debugf(msg, v...)
}

// Infof level formatted message.
func Infof(msg string, v ...interface{}) {
	Log.Infof(msg, v...)
}

// Warnf level formatted message.
func Warnf(msg string, v ...interface{}) {
	Log.Warnf(msg, v...)
}

// Errorf level formatted message.
func Errorf(msg string, v ...interface{}) {
	Log.Errorf(msg, v...)
}

// Fatalf level formatted message, followed by an exit.
func Fatalf(msg string, v ...interface{}) {
	Log.Fatalf(msg, v...)
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func Trace(msg string) *Entry {
	return Log.Trace(msg)
}
