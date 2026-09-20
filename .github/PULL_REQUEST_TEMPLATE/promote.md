<!-- Release PR: dev -> main. Merge with "Create a merge commit". Never squash. -->

## Promote dev to main

Target: <!-- demo-proposal / sprint-3 / demo-final -->

## Included tickets
- LAP-

## Before merge
- [ ] Staging (dev) tested end to end by at least 2 people
- [ ] New migrations listed below and tested on dev DB
- [ ] New env vars added to Vercel **Production** + backend prod host
- [ ] Manual `pg_dump` of prod DB taken
- [ ] Not within 48h of a demo (or team agreed)

## Migrations
-

## After merge
- [ ] Prod URL smoke-tested
- [ ] Tag created: `git tag <name> origin/main && git push origin <name>`
