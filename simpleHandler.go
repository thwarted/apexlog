// Package text implements a development-friendly textual handler.
package log

import (
	"fmt"
	"io"
	"sort"
	"sync"
	"time"
)

// start time.
var shstart = time.Now()

// Strings mapping.
var Strings = [...]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	FatalLevel: "FATAL",
}

// field used for sorting.
type shfield struct {
	Name  string
	Value interface{}
}

// by sorts projects by call count.
type byName []shfield

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Handler implementation.
type SimpleHandler struct {
	mu     sync.Mutex
	Writer io.Writer
}

// New handler.
func NewSimple(w io.Writer) Handler {
	return &SimpleHandler{
		Writer: w,
	}
}

// HandleLog implements log.Handler.
func (h *SimpleHandler) HandleLog(e *Entry) error {
	level := Strings[e.Level]

	var fields []shfield

	for k, v := range e.Fields {
		fields = append(fields, shfield{k, v})
	}

	sort.Sort(byName(fields))

	h.mu.Lock()
	defer h.mu.Unlock()

	ts := time.Since(shstart) / time.Second
	fmt.Fprintf(h.Writer, "%6s[%04d] %-25s", level, ts, e.Message)

	for _, f := range fields {
		fmt.Fprintf(h.Writer, " %s=%v", f.Name, f.Value)
	}

	fmt.Fprintln(h.Writer)

	return nil
}
