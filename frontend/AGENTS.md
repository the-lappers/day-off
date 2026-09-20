<!-- BEGIN:nextjs-agent-rules -->

# This is NOT the Next.js you know

This version has breaking changes — APIs, conventions, and file structure may all differ from your training data. Read the relevant guide in `node_modules/next/dist/docs/` (resolved from this file's directory; in monorepos the `next` package may not be visible from the repo root) before writing any code. Heed deprecation notices.

This block is written and re-added by `next dev` — verify at `node_modules/next/dist/server/lib/generate-agent-files.js`. Removing it from a diff only re-creates the uncommitted change; committing it with your work keeps the tree clean.

<!-- END:nextjs-agent-rules -->
# frontend/ — Next.js app

Only what differs here. Root `AGENTS.md` still applies.

- Next.js App Router with TypeScript. Own `package.json` and lockfile in this folder.
- `npm run lint` must pass; the root `make lint` calls it.
- API types come from `docs/api/openapi.yaml`. Don't hand-write a response shape that
  isn't in the spec.
- Backend base URL comes from `NEXT_PUBLIC_API_URL`. Never hardcode `localhost`.
- UI text is Thai first, English second (NFR-L1).
- Must render correctly from 360px width up (NFR-O1).
- Show venue scores as tags and numbers only. Never surface the source review text or
  photos as evidence for a score (ADR-006).
