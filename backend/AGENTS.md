# backend/ — Go API

Only what differs here. Root `AGENTS.md` still applies.

- Go module lives in this folder. Run commands from here: `go test ./...`, `go vet ./...`.
- Layout: `cmd/server` (entry point), `internal/httpserver` (routes, handlers, CORS),
  `internal/config` (env), `internal/database` (pgx pool). `Dockerfile` builds the image
  `docker compose` runs.
- `gofmt` must be clean; `make lint` fails otherwise.
- Serves requests by reading precomputed scores from Postgres. **No ML inference, no
  model loading, no calls into `pipeline/`.**
- Every endpoint must match `docs/api/openapi.yaml`. Spec first, then the handler.
- Health check stays at `/healthz`.
- Port 8000, `CORS_ORIGINS` from env. Never hardcode URLs.
- Database access goes through the migration-defined schema. Don't create tables from Go.
