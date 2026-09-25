package httpserver

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type healthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

// healthHandler serves GET /healthz: 200 when Postgres answers a ping within
// timeout, 503 when it doesn't.
func healthHandler(db Pinger, timeout time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		if err := db.Ping(ctx); err != nil {
			slog.Warn("health check: database unreachable", "err", err)
			writeJSON(w, http.StatusServiceUnavailable, healthResponse{Status: "unavailable", Database: "unreachable"})
			return
		}
		writeJSON(w, http.StatusOK, healthResponse{Status: "ok", Database: "ok"})
	}
}
