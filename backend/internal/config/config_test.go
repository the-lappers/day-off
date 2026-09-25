package config

import (
	"reflect"
	"testing"
)

func env(vars map[string]string) func(string) string {
	return func(k string) string { return vars[k] }
}

func TestLoadDefaultsPort(t *testing.T) {
	cfg, err := Load(env(map[string]string{"DATABASE_URL": "postgres://localhost/db"}))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Port != "8000" {
		t.Errorf("Port = %q, want 8000", cfg.Port)
	}
	if cfg.CORSOrigins != nil {
		t.Errorf("CORSOrigins = %v, want none", cfg.CORSOrigins)
	}
}

func TestLoadParsesCORSOrigins(t *testing.T) {
	cfg, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://localhost/db",
		"CORS_ORIGINS": " http://localhost:3333, ,https://example.test ",
	}))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	want := []string{"http://localhost:3333", "https://example.test"}
	if !reflect.DeepEqual(cfg.CORSOrigins, want) {
		t.Errorf("CORSOrigins = %v, want %v", cfg.CORSOrigins, want)
	}
}

func TestLoadRejectsBadInput(t *testing.T) {
	cases := map[string]map[string]string{
		"missing DATABASE_URL": {},
		"non-numeric PORT":     {"DATABASE_URL": "postgres://localhost/db", "PORT": "http"},
		"PORT out of range":    {"DATABASE_URL": "postgres://localhost/db", "PORT": "70000"},
	}
	for name, vars := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := Load(env(vars)); err == nil {
				t.Error("Load returned no error")
			}
		})
	}
}
