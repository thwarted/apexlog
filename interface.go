package log

// Interface represents the API of both Logger and Entry.
type Interface interface {
	WithFields(fields Fielder) *Entry
	WithField(key string, value interface{}) *Entry
	WithError(err error) *Entry
	WithMemStats() *Entry
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Println(v ...interface{})
	Debugf(msg string, v ...interface{})
	Infof(msg string, v ...interface{})
	Warnf(msg string, v ...interface{})
	Errorf(msg string, v ...interface{})
	Fatalf(msg string, v ...interface{})
	Printf(msg string, v ...interface{})
	Trace(msg string) *Entry
}

// Any values that are dynamically determined at log time
// using a function call should be cast to this type.
type Fn func() interface{}
