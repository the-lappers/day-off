package database

import (
	"context"
	"testing"
)

func TestOpenRejectsInvalidURL(t *testing.T) {
	if _, err := Open(context.Background(), "postgres://%zz"); err == nil {
		t.Fatal("Open returned no error for an invalid URL")
	}
}

func TestOpenDoesNotConnect(t *testing.T) {
	// Nothing listens on port 1. Open must still succeed because the pool
	// connects lazily, which lets the API start while Postgres is down.
	pool, err := Open(context.Background(), "postgres://user:pass@127.0.0.1:1/db")
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	pool.Close()
}
