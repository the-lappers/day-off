# Playbook: write an ADR

An ADR records a decision that is expensive to reverse: a language, a datastore, a schema
shape, a workflow. Not every choice needs one. If someone could reasonably ask "why did
you do it that way?" six months later, write it.

## Steps

1. Next free number in `docs/adr/`, e.g. `ADR-008`.
2. Copy the structure below.
3. Open a PR with the ADR alone, or with the change it justifies.
4. Status is `Proposed` until the team agrees, then `Accepted`. Superseded ADRs keep
   their file and gain a line pointing at the replacement.

## Structure

```
# ADR-00X: <decision in a few words>

Status: Proposed | Accepted | Superseded by ADR-00Y

## Context
What forced the decision. Constraints, requirements by id, what we tried.

## Decision
One sentence. What we are doing.

## Consequences
What this makes easy, what it makes hard, what it costs. Both directions, honestly.

## Alternatives considered
Each option, and the specific reason it lost. "Worse" is not a reason.

## Requirement changes
FRs or NFRs added, modified, or confirmed by this decision.
```

## Rules

- Write consequences you dislike. An ADR that only lists upsides is marketing, and the
  committee reads for trade-off awareness.
- Reference requirements by id (FR-303, NFR-P1) so the decision links to the spec.
- Never edit an accepted ADR to change its decision. Write a new one that supersedes it.
