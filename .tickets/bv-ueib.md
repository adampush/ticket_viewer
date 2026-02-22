---
id: bv-ueib
status: closed
deps: [bv-ojjg]
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 1
assignee: Adam Push
parent: bv-jzv8
---
# Migrate active env/config namespace from BV_* to TKV_*

## Context / Why

This is Workstream 4 from `docs/tkv-native-cleanup-plan.md`. Active env/config namespace must be `TKV_*` to match product identity and eliminate contract ambiguity.

## Scope

### In Scope

- Replace active `BV_*` env lookups with `TKV_*` in runtime codepaths.
- Update docs and tests to align with `TKV_*` contract where applicable.
- Preserve locked default behavior: legacy `BV_*` ignored in active codepaths.

### Out of Scope

- Module path migration (Workstream 5).
- Packaging/distribution channel updates (Workstream 6).

## Ticket Type & Granularity

- Type: Task
- Granularity target: 1-3 commits focused on env/config contract migration.

## Prerequisites & Dev State

- Upstream terminology cleanup `bv-ojjg` complete.
- Contract matrix and env behavior defaults locked in plan.

## Assumptions

- Active codepaths are the migration gate; test fixtures/historical notes may still reference legacy names if explicitly excluded by plan gate.

## Open Questions

- Should any targeted warning be emitted when legacy `BV_*` vars are set?
  - Owner: Adam Push
  - Timing: before merge
  - Blocking: no (default remains "ignore legacy vars")

## Implementation Guidance

- Update env lookups across `cmd/bv`, `pkg/ui`, `pkg/export`, `pkg/agents`, `internal/datasource`.
- Update docs surfaces (`README.md`, `AGENTS.md`) for active env namespace.
- Update tests that intentionally verify active env behavior.
- Inherit NFR constraints: no behavioral regressions beyond naming contract change.

## Acceptance Criteria

- Active codepaths use `TKV_*` env namespace.
- No active runtime path depends on `BV_*` values.
- Docs for active env contract are `TKV_*`.

## Deployment & State

- Breaking/Non-breaking: breaking for users depending on legacy `BV_*` active behavior.
- Migration requirements: release notes include env migration examples.
- Safe to deploy independently: yes, after Workstream 3.

## Validation Plan

- grep check for `\bBV_[A-Z0-9_]+\b` in active code/docs (excluding allowed non-active contexts per plan)
- `go test ./cmd/bv/...`
- `go test ./pkg/ui/...`
- `go test ./pkg/export/...`
- `go test ./pkg/agents/...`
- `go test ./internal/datasource/...`
- `go build ./...`

Expected evidence:

- grep logs + targeted test outputs in PR checklist.
- `tk add-note bv-ueib "validation: ..."` with pass signals.

## Dependencies

- Upstream: `bv-ojjg`
- Downstream: `bv-hqio`

## Artifacts

- Env migration note block (old->new examples)
- Validation logs and checklist evidence

## Notes

**2026-02-22T03:07:17Z**

Migrated active env namespace from BV_* to TKV_* across runtime codepaths (cmd/ui/export/agents/internal/search/analysis/watcher/hooks), updated docs (README.md, AGENTS.md), and aligned affected tests including e2e env contracts. Validation: go test ./cmd/bv/... ./pkg/ui/... ./pkg/export/... ./pkg/agents/... ./internal/datasource/...; go test ./...; go build ./...; go vet ./...; grep check: rg -n '\bBV_[A-Z0-9_]+\b' cmd pkg internal README.md AGENTS.md --glob '!**/*_test.go' (no matches).
