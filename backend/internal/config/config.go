// Package config reads the API's settings from environment variables.
package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const defaultPort = "8000"

// Config holds the settings the API needs at startup. See .env.example.
type Config struct {
	Port        string
	DatabaseURL string
	CORSOrigins []string
}

// Load builds a Config from getenv (os.Getenv in production).
func Load(getenv func(string) string) (Config, error) {
	cfg := Config{
		Port:        getenv("PORT"),
		DatabaseURL: getenv("DATABASE_URL"),
		CORSOrigins: splitList(getenv("CORS_ORIGINS")),
	}
	if cfg.Port == "" {
		cfg.Port = defaultPort
	}
	if n, err := strconv.Atoi(cfg.Port); err != nil || n < 1 || n > 65535 {
		return Config{}, fmt.Errorf("PORT must be a number between 1 and 65535, got %q", cfg.Port)
	}
	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}
	return cfg, nil
}

// splitList turns "a, b,,c" into [a b c].
func splitList(s string) []string {
	var out []string
	for _, part := range strings.Split(s, ",") {
		if p := strings.TrimSpace(part); p != "" {
			out = append(out, p)
		}
	}
	return out
}
