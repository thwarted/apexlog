package log

import (
	"fmt"
)

// Printf wraps Infof
func (l *Logger) Printf(msg string, v ...interface{}) {
	NewEntry(l).Infof(msg, v...)
}

// Printf wraps Infof
func (e *Entry) Printf(msg string, v ...interface{}) {
	e.Infof(msg, v...)
}

// Printf wraps Infof
func Printf(msg string, v ...interface{}) {
	Log.Infof(msg, v...)
}

// Println wraps Info
func (l *Logger) Println(v ...interface{}) {
	NewEntry(l).Info(v...)
}

// Println wraps Info
func (e *Entry) Println(v ...interface{}) {
	e.Info(v...)
}

// Println wraps Info
func Println(v ...interface{}) {
	Log.Info(v...)
}

// Sprintlnn => Sprint no newline. This is to get the behavior of how
// fmt.Sprintln where spaces are always added between operands, regardless of
// their type. Instead of vendoring the Sprintln implementation to spare a
// string allocation, we do the simplest thing.
func sprintlnn(args ...interface{}) string {
	// fast path for single, simple string
	if len(args) == 1 {
		if s, ok := args[0].(string); ok {
			return s
		}
	}
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}
