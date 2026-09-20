# Playbook: document a change

Docs are updated in the PR that makes them wrong. A stale doc is worse than a missing
one, because both people and AI tools read it as true.

## Which file

| You changed | Update |
|---|---|
| An endpoint | `docs/api/openapi.yaml` (CI enforces this) |
| How components talk, or a new component | `docs/ARCHITECTURE.md` |
| A decision that is expensive to reverse | New ADR in `docs/adr/` |
| A command, a folder's purpose, a hard rule | `AGENTS.md`, or the nested one for that folder |
| Local setup steps or prereqs | `README.md` |
| Branching, review, or release process | `CONTRIBUTING.md` |
| A problem that cost over an hour | `docs/LESSONS_LEARNED.md` |
| A chore you have now explained twice | New file in `docs/playbooks/` |

## Rules

- Same PR, not "later". Later never arrives during a sprint.
- Keep `AGENTS.md` short. Agents truncate long instruction files, so detail goes in
  `docs/` and `AGENTS.md` links to it.
- Nested `AGENTS.md` files hold only what differs from the root. Never repeat the root.
- Requirements are referenced by id (FR-401, NFR-D2) rather than reworded.
