// Package discard implements a no-op handler useful for benchmarks and tests.
package discard

import (
	log "github.com/thwarted/apexlog"
)

// Default handler.
var Default = New()

// Handler implementation.
type Handler struct{}

// New handler.
func New() *Handler {
	return &Handler{}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	return nil
}
