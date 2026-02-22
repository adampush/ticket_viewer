---
id: bv-hqio
status: open
deps: [bv-ueib]
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 1
assignee: Adam Push
parent: bv-jzv8
---
# Migrate Go module path and rewrite internal imports to ticket_viewer

## Context / Why

This is Workstream 5 from `docs/tkv-native-cleanup-plan.md`. The module/import namespace must be fully migrated to `github.com/adampush/ticket_viewer` for a true tkv-native identity.

## Scope

### In Scope

- Update `go.mod` module path to `github.com/adampush/ticket_viewer`.
- Rewrite internal Go imports to new module path.
- Ensure build/test/vet passes after migration.

### Out of Scope

- Packaging/release channel updates (Workstream 6).
- Unrelated behavior refactors.

## Ticket Type & Granularity

- Type: Task
- Granularity target: 1-3 commits focused on module/import namespace migration.

## Prerequisites & Dev State

- Upstream env/config migration `bv-ueib` complete.
- Workstream 5 contract defaults locked in plan.

## Assumptions

- Partial migration is invalid; this ticket must land as a complete cutover in one PR.

## Open Questions

- None currently. Any integration-specific fallout must be documented as follow-up tickets.

## Implementation Guidance

- Update module declaration in `go.mod`.
- Rewrite Go imports across `cmd`, `pkg`, `internal`, `tests`, and helper Go scripts.
- Keep runtime behavior unchanged; this is a namespace cutover.
- Inherit NFR constraints: no regressions in determinism, timeout behavior, or test-mode browser guards.

## Acceptance Criteria

- `go.mod` points to `github.com/adampush/ticket_viewer`.
- No Go imports reference `github.com/Dicklesworthstone/beads_viewer`.
- `go build ./...`, `go vet ./...`, and `go test ./...` pass.

## Deployment & State

- Breaking/Non-breaking: breaking for downstream integrations pinned to old module path.
- Migration requirements: release notes include module path cutover note.
- Safe to deploy independently: yes, after Workstream 4.

## Validation Plan

- grep check for `github.com/Dicklesworthstone/beads_viewer` in Go sources
- `go build ./...`
- `go vet ./...`
- `go test ./...`

Expected evidence:

- grep output proving no legacy import references
- full-suite outputs attached in PR checklist
- `tk add-note bv-hqio "validation: ..."` with results

## Dependencies

- Upstream: `bv-ueib`
- Downstream: `bv-4nc1`

## Artifacts

- Module/import migration evidence
- Release note draft snippet for module path change
