# pipeline/ — Python batch job

Only what differs here. Root `AGENTS.md` still applies.

- Python tooling is `uv` + `ruff`. Run `uv run <cmd>` from this folder.
- This is a **nightly batch job**, not a service. Nothing here is called during a request.
- It reads raw snapshots from private storage (`RAW_STORAGE_URL`) and writes scores to
  Postgres. It does **not** fetch from sources: that lives in the private
  `day-off-ingest` repo.
- Keep fetching, scoring, and writing as separate steps so scoring can rerun on the same
  snapshot without refetching.
- Model weights (WangchanBERTa, CLIP) are downloaded at runtime. Never commit weights.
- Never print raw review text or scores in logs. CI logs are public.
- Every score needs mention count, recency, and confidence alongside it (FR-101 to FR-105).
