---
id: bv-ojjg
status: closed
deps: [bv-y047]
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 2
assignee: Adam Push
parent: bv-jzv8
---
# Clean internal user-visible terminology in core packages

## Context / Why

This is Workstream 3 from `docs/tkv-native-cleanup-plan.md`. Mixed bead/ticket terminology in primary execution paths harms maintainability and causes output ambiguity.

## Scope

### In Scope

- Normalize user-visible internal terminology to ticket-native names in core packages.
- Keep changes mechanical and scoped to naming/readability and surfaced output text.

### Out of Scope

- Deep algorithm refactors in analysis/search.
- Env/config namespace migration (Workstream 4).

## Ticket Type & Granularity

- Type: Task
- Granularity target: 1-3 commits scoped to core package naming cleanup.

## Prerequisites & Dev State

- Upstream payload/schema normalization `bv-y047` complete.
- Current contracts matrix and acceptance checks available.

## Assumptions

- No intended behavior change beyond naming/output terminology consistency.

## Open Questions

- None currently; any required semantic behavior change should be split into a follow-up ticket.

## Implementation Guidance

- Touch naming in primary code paths in `pkg/analysis`, `pkg/correlation`, `pkg/export`, `pkg/ui`, `pkg/loader` where terminology leaks to users or maintainers.
- Preserve output semantics except intentional naming normalization.
- Inherit NFR constraints: no performance/security/scale regressions from rename-only changes.

## Acceptance Criteria

- Primary execution paths avoid confusing mixed bead/ticket terminology.
- User-visible text in touched core paths is ticket-native.
- No unintended behavior change in analysis and robot outputs.

## Deployment & State

- Breaking/Non-breaking: expected non-breaking except wording updates.
- Migration requirements: none beyond documentation consistency.
- Safe to deploy independently: yes, after Workstream 2.

## Validation Plan

- `go test ./pkg/analysis/...`
- `go test ./pkg/correlation/...`
- `go test ./pkg/export/...`
- `go test ./pkg/ui/...`
- `go test ./pkg/loader/...`
- `go build ./...`

Expected evidence:

- Package test outputs in PR checklist and ticket notes.
- `tk add-note bv-ojjg "validation: ..."` with non-regression evidence.

## Dependencies

- Upstream: `bv-y047`
- Downstream: `bv-ueib`

## Artifacts

- Naming cleanup notes for touched packages
- Validation evidence in ticket notes

## Notes

**2026-02-22T02:53:25Z**

Completed WS3 terminology cleanup in core user-visible paths. Updated history/insights/TUI status text to issue-native wording (search placeholders, mode labels, empty states, badges, hints), sprint and graph UI copy to issue terminology, correlation explanation/reason strings to issue wording, and loader/correlation error messages to issue-native phrasing. Also updated related unit tests and snapshots/expectations accordingly. Validation passed: go test ./pkg/analysis/... ./pkg/correlation/... ./pkg/export/... ./pkg/ui/... ./pkg/loader/... plus full gates go build ./..., go vet ./..., go test ./... .
