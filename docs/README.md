# Documentation index

| File | What it holds |
|---|---|
| `ARCHITECTURE.md` | How the pieces fit: request path, batch path, data boundaries |
| `adr/` | Numbered decision records. Why things are the way they are. |
| `api/openapi.yaml` | The API contract. Source of truth for every endpoint. |
| `playbooks/` | Step-by-step for recurring chores (ADR, migration, API change, docs) |
| `data-boundary.md` | What lives in this repo, the private repo, storage, and the database |
| `LESSONS_LEARNED.md` | Problems that cost real time, and what fixed them |

Outside `docs/`:

| File | What it holds |
|---|---|
| `../README.md` | Prereqs and how to run the stack locally |
| `../CONTRIBUTING.md` | Branching, PRs, reviews, releases |
| `../AGENTS.md` | Instructions for AI coding tools (plus nested files per folder) |

Rule: if a change makes one of these wrong, fix it in the same PR.
