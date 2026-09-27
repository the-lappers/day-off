// Package httpserver wires HTTP routes and middleware. Every route here must
// match docs/api/openapi.yaml.
package httpserver

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// Pinger is the slice of the database pool the handlers need.
type Pinger interface {
	Ping(ctx context.Context) error
}

const healthTimeout = 2 * time.Second

// NewHandler returns the API's root handler.
func NewHandler(db Pinger, corsOrigins []string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler(db, healthTimeout))
	return withCORS(corsOrigins, mux)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("write response", "err", err)
	}
}
