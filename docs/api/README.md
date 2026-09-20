# API contract

`openapi.yaml` lives here and is the source of truth for every endpoint.

It does not exist yet. The backend owner writes the first version covering the Tier 0
endpoints: vibe search (UC-01), venue detail (UC-02), trip create and read (UC-03).

Until then the CI tripwire has nothing to compare against. Land the spec first, then the
check in `docs/playbooks/change-the-api.md` starts protecting it.
