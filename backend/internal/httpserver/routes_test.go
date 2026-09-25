package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeDB struct{ err error }

func (f fakeDB) Ping(context.Context) error { return f.err }

const allowedOrigin = "http://localhost:3333"

func serve(t *testing.T, db Pinger, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	NewHandler(db, []string{allowedOrigin}).ServeHTTP(rec, req)
	return rec
}

func TestHealthz(t *testing.T) {
	cases := []struct {
		name       string
		db         fakeDB
		wantCode   int
		wantStatus string
		wantDB     string
	}{
		{"database up", fakeDB{}, http.StatusOK, "ok", "ok"},
		{"database down", fakeDB{err: errors.New("connection refused")}, http.StatusServiceUnavailable, "unavailable", "unreachable"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := serve(t, tc.db, httptest.NewRequest(http.MethodGet, "/healthz", nil))
			if rec.Code != tc.wantCode {
				t.Fatalf("code = %d, want %d", rec.Code, tc.wantCode)
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("Content-Type = %q", ct)
			}
			var body healthResponse
			if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			if body.Status != tc.wantStatus || body.Database != tc.wantDB {
				t.Errorf("body = %+v", body)
			}
		})
	}
}

func TestHealthzRejectsOtherMethods(t *testing.T) {
	rec := serve(t, fakeDB{}, httptest.NewRequest(http.MethodPost, "/healthz", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("code = %d, want 405", rec.Code)
	}
}

func TestCORS(t *testing.T) {
	t.Run("allowed origin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		req.Header.Set("Origin", allowedOrigin)
		rec := serve(t, fakeDB{}, req)
		if got := rec.Header().Get("Access-Control-Allow-Origin"); got != allowedOrigin {
			t.Errorf("Access-Control-Allow-Origin = %q", got)
		}
	})
	t.Run("other origin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		req.Header.Set("Origin", "https://evil.test")
		rec := serve(t, fakeDB{}, req)
		if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "" {
			t.Errorf("Access-Control-Allow-Origin = %q, want none", got)
		}
	})
	t.Run("preflight", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/healthz", nil)
		req.Header.Set("Origin", allowedOrigin)
		req.Header.Set("Access-Control-Request-Method", http.MethodGet)
		rec := serve(t, fakeDB{}, req)
		if rec.Code != http.StatusNoContent {
			t.Errorf("code = %d, want 204", rec.Code)
		}
		if rec.Header().Get("Access-Control-Allow-Methods") == "" {
			t.Error("missing Access-Control-Allow-Methods")
		}
	})
}
