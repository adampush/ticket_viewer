---
id: bv-wcue
status: open
deps: []
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 1
assignee: Adam Push
parent: bv-jzv8
---
# Rename public CLI and robot surface to ticket-native terminology

## Context / Why

This is Workstream 1 from `docs/tkv-native-cleanup-plan.md`. Public CLI/robot surfaces still expose bead-era naming, which creates immediate confusion for users and agents.

## Scope

### In Scope

- Rename public CLI/help/robot-help wording to ticket-native active terminology.
- Remove bead-era active flag/command terms from help output.
- Update help/documentation snapshots that assert CLI surface text.

### Out of Scope

- Robot payload key/schema renames (Workstream 2).
- Env/config namespace migration (Workstream 4).

## Ticket Type & Granularity

- Type: Task
- Granularity target: 1-3 commits focused on `cmd/bv` public surface and related help tests.

## Prerequisites & Dev State

- Approved plan: `docs/tkv-native-cleanup-plan.md` (PASS under dual-lens review).
- Parent umbrella: `bv-jzv8`.
- Work branch must include current mainline merges.

## Assumptions

- Legacy bead-era names in active public help are breaking changes and should not be retained as compatibility aliases.

## Open Questions

- None currently. If discovered, add via ticket notes and mark blocking status explicitly.

## Implementation Guidance

- Update user-facing CLI and robot-help text in `cmd/bv/main.go` and related help emitters.
- Ensure examples and hints in active help surfaces are `tkv`/ticket-native.
- Inherit NFR constraints: no startup-performance regression and no behavioral changes outside naming.

Likely files:

- `cmd/bv/main.go`
- help/docs snapshot tests under `cmd/bv` and `tests/e2e` as needed

## Acceptance Criteria

- `tkv --help` contains no active bead-era terminology (except explicitly historical notes).
- `tkv --robot-help` contains no active bead-era terminology.
- No bead-era active flags remain in the CLI surface.

## Deployment & State

- Breaking/Non-breaking: breaking for users relying on old naming in active help/flags.
- Migration requirements: document before/after examples in release notes.
- Safe to deploy independently: yes (before Workstream 2).

## Validation Plan

- `tkv --help`
- `tkv --robot-help`
- `tkv --robot-docs all`
- `go test ./cmd/bv/...`
- `go test ./tests/e2e/... -run 'Robot|Help|Docs'` (if relevant tests exist)
- `go build ./...`

Expected evidence:

- Command outputs/snapshots attached in PR checklist.
- `tk add-note bv-wcue "validation: ..."` with pass signals.

## Dependencies

- Upstream: none
- Downstream: `bv-y047`

## Artifacts

- Updated help/docs snapshots
- PR checklist evidence for surface checks
- Ticket notes with validation outputs
