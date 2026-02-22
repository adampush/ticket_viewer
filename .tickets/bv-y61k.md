---
id: bv-y61k
status: open
deps: [bv-4nc1]
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 2
assignee: Adam Push
parent: bv-jzv8
---
# Run end-to-end validation and release migration notes for tkv-native cleanup

## Context / Why

This integration ticket closes `bv-jzv8` by proving cross-stream verification/validation and documenting final migration communications.

## Scope

### In Scope

- Execute full engineering verification gates after all workstreams merge.
- Execute product validation checks for key user flows (CLI, robot, env/config, installer).
- Produce final migration notes and completion evidence for umbrella closure.

### Out of Scope

- New implementation work beyond remediation for validation failures.

## Ticket Type & Granularity

- Type: Task (Integration)
- Granularity target: 1-3 commits for final validation artifacts/docs updates.

## Prerequisites & Dev State

- Upstream workstreams through `bv-4nc1` complete.
- All child-ticket validation evidence available in notes/PRs.

## Assumptions

- Any failures found here are fixed by reopening/splitting targeted follow-up tickets, not by broad scope creep in this ticket.

## Open Questions

- None currently.

## Implementation Guidance

- Run final global verification commands and stream-specific contract checks.
- Validate product user flows from plan's Validation section.
- Consolidate migration communication artifacts and close umbrella if all gates pass.

## Acceptance Criteria

- Full verification suite passes (`go build`, `go vet`, `go test ./...`).
- Product validation checks pass for CLI/robot/env/install flows.
- Final migration notes include breaking-change guidance and before/after examples.
- `bv-jzv8` has complete closure evidence with linked child outcomes.

## Deployment & State

- Breaking/Non-breaking: documentation/integration closure only.
- Migration requirements: publish final migration summary and known constraints.
- Safe to deploy independently: yes (post-stream closure validation).

## Validation Plan

- `go build ./...`
- `go vet ./...`
- `go test ./...`
- `tkv --help`
- `tkv --robot-help`
- `tkv --robot-schema`
- grep checks for active legacy contracts (`BV_*` and old module path) in active codepaths

Expected evidence:

- Consolidated validation table in ticket notes
- links to child ticket notes/PRs
- explicit pass/fail summary for each user flow

## Dependencies

- Upstream: `bv-4nc1`
- Downstream: `bv-jzv8` closure

## Artifacts

- Final migration release-note draft
- umbrella closure note with verification/validation evidence
- follow-up ticket links for any deferred hardening
