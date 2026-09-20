# Data boundary: what lives where

One rule: **code that fetches data is private, code that transforms or serves it is public.
Data itself is never in a git repo.**

| Thing | Where | Why |
|---|---|---|
| App code: `frontend/`, `backend/`, `pipeline/`, `eval/`, `infra/` | this repo | It's the method, and it's what we're graded on |
| Fetch adapters: APIs, open data, scrapers, manual loaders | private `day-off-ingest` repo | Source terms and scraping details stay out of public view |
| Raw reviews, photos, annotations, eval gold set | private storage (decision pending) | Licensing, PDPA, and git handles large files badly |
| Vibe store: venues, scores, embeddings | Supabase (dev + prod) | A database, not a file. Reachable only with keys. |
| Secrets | `.env` locally, Vercel and GitHub Secrets for deploys | Never in git, in either repo |

## The handoff

Adapters write a shared raw format. The scorer in `pipeline/` reads only that format and
never knows which source a record came from.
day-off-ingest/adapters/* ──► raw snapshot in storage ──► day-off/pipeline/ ──► Supabase

The raw format is a contract between the two repos. Changing it needs a PR in both.

## Rules that follow from this

- Never commit a data file here. CI fails PRs adding files over 1 MB or `.csv`/`.jsonl`/
  `.parquet`/model weights.
- Tests use fake sample text, never real reviews.
- Never print raw review text or scores in CI logs.
- Google Places responses are request-time only: never stored, never training data (NFR-L2).

## Still open

- Name confirmed as `day-off-ingest`, repo not created yet (Phase 4)
- Where raw snapshots live: Supabase Storage, Hugging Face, or Drive
- Who owns the raw format spec (pipeline owner)
