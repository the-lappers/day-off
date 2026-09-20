<!-- PR title = squash commit on main. Format: type(scope): summary -->
<!-- e.g. feat(frontend): add vibe search box -->
<!-- Base branch: dev (normal work). main only for hotfix/ branches. -->

## What
<!-- 1-2 lines -->

## Why
LAP-
<!-- the Jira key links this PR to the ticket. add #close to transition it, e.g. "LAP-31 #close" -->

## How to test
1.

## Screenshots
<!-- UI changes only, delete otherwise -->

## Checklist
- [ ] `make lint` and `make test` pass locally
- [ ] PR is small (aim < 400 changed lines)
- [ ] New env vars added to `.env.example`
- [ ] Docs / ADR updated if behavior or architecture changed
- [ ] No secrets, scraped data, or Google Places data committed
- [ ] Hotfix into main? Run `make sync` right after merge
