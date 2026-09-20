# AGENTS.md

Day-Off: a one-day trip planner for Bangkok neighbourhoods. Users describe a vibe in
free text (Thai or English); the app ranks venues and builds a time-feasible day route.

Nested `AGENTS.md` files in `frontend/`, `backend/`, `pipeline/` add rules for those
areas. This file holds what applies everywhere. Keep both short.

## Commands

| Command | Does |
|---|---|
| `make setup` | Hooks, deps, `.env`. Run once after cloning. |
| `make dev` | Local Supabase + backend + frontend |
| `make down` | Stop the stack |
| `make pipeline` | Run the batch pipeline once |
| `make lint` | Lint frontend, backend, pipeline |
| `make test` | Test all three |
| `make help` | Everything else |

Never invent a new task runner or script path. Add targets to the `Makefile` instead.
`setup`, `lint`, `test`, `eval` are required CI check names and must not be renamed.

## Layout

```
frontend/   Next.js + TypeScript. User-facing app. Port 3333.
backend/    Go API. Serves requests by reading Postgres. Port 8000.
pipeline/   Python batch job. Scores venues nightly, writes to Postgres.
eval/       Retrieval quality harness (nDCG@5).
infra/      Deployment and compose config.
supabase/   Migrations and local DB config.
docs/       Architecture, ADRs, API spec, playbooks.
```

## Hard rules

- **Go never calls Python, Python never calls Go.** Postgres is the only interface
  between them. No model runs at request time.
- **Schema changes are migration files** in `supabase/migrations/`. Never hand-edit
  tables in the Supabase dashboard, local or hosted.
- **No data in this repo.** No raw reviews, photos, annotations, scraped content, model
  weights, `.csv`/`.jsonl`/`.parquet`. Raw data lives in private storage; fetch and
  scraping code lives in the private `day-off-ingest` repo. CI fails PRs that add them.
- **No secrets.** New env var means a line in `.env.example` in the same PR, with an
  empty value.
- **Thai is the primary UI language**, English second. User-visible strings must work in
  both.
- **Google Places data is request-time only.** Never stored, never used as training data.

## API contract

`docs/api/openapi.yaml` is the source of truth for every endpoint.

Change the API and you change the spec **in the same PR**. Write the spec change first,
then implement. CI blocks a PR that edits Go handlers or routes without touching the
spec; see `docs/playbooks/change-the-api.md`.

## Where to look

| Question | File |
|---|---|
| How does the system fit together? | `docs/ARCHITECTURE.md` |
| Why is it built this way? | `docs/adr/` |
| How do I branch, commit, review, release? | `CONTRIBUTING.md` |
| What are the endpoints? | `docs/api/openapi.yaml` |
| How do I do a recurring chore? | `docs/playbooks/` |
| What bit us before? | `docs/LESSONS_LEARNED.md` |
| How do I run this locally? | `README.md` |

## Conventions

- Branch `type/LAP-<id>-<slug>`, PR title `type(scope): summary`, Jira key `LAP-<id>` in
  the body. Squash into `dev`. Full rules in `CONTRIBUTING.md`.
- Never push to `dev` or `main` directly.
- Requirements are referenced by id (FR-301, NFR-P1). Use them in PR descriptions and
  test names where they apply.
- Update the doc that a change invalidates, in the same PR.
