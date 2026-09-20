# Playbook: change the API

`docs/api/openapi.yaml` is the contract between the Go API and the frontend. It is also
what AI tools read to generate correct client and handler code.

## Rule

**Spec first, then code, in the same PR.**

1. Edit `docs/api/openapi.yaml`: path, request shape, response shape, error cases.
2. Commit that alone, so the diff is reviewable on its own.
3. Implement the handler in `backend/`.
4. Update frontend types and calls to match.

For a feature spanning backend and frontend, merge the spec change as **its own PR**
first. Both people then build against it in parallel, the backend with fake data and the
frontend with a mock, and neither waits for the other.

## The CI tripwire

CI fails a PR that changes Go handler or route files without touching the spec:

```
API code changed but docs/api/openapi.yaml did not
```

It checks that the spec **moved**, not that it is correct. Only review catches a spec
that disagrees with the code, so reviewers read the spec diff first.

CI also parses the spec, so a broken YAML file fails fast.

### Escape hatch

A refactor that genuinely doesn't change the contract (renaming an internal function,
moving a file) can pass by adding the label `spec-exempt` to the PR and saying why in the
body. The label shows in the PR list, so overuse is visible.

## Scope of the spec

Tier 0 endpoints only, to start: vibe search, venue detail, trip create and read. Add an
endpoint when its ticket starts, not before. Don't spec features that aren't being built.
