---
id: bv-g3nq
status: closed
deps: [bv-92o5, bv-v3f4, bv-c56w]
links: []
created: 2026-02-20T13:03:47Z
type: task
priority: 1
assignee: Adam Push
parent: bv-0kqn
---
# Add dual-mode tests for detection, parsing, and contracts

## Context / Why

Dual-mode support touches source selection, parsing, and robot contracts. We need explicit coverage to prevent regressions and to prove deterministic behavior.

## Scope

### In Scope

- Add/update tests for:
  - tracker detection precedence
  - `tk` parsing and mapping edge cases
  - unknown `tk` status handling (malformed -> skip)
  - robot command contract behavior by mode
  - export actionable command snippets behavior by mode
  - mixed-repo deterministic selection
  - helper command validity in `tk` mode (no Beads-only flags)
  - `tk`-only end-to-end load path with no `.beads` directory
  - mixed repo where `.tickets` exists but all ticket files are malformed (expect clear error, no Beads fallback)
  - datasource metadata path correctness (selected source/tracker + mixed-source diagnostics)
  - compatibility coverage for existing `datasource.LoadIssues(...)` callers
  - backward-compat export API behavior for existing callsites
  - neither-source error behavior (no `.tickets`, no `.beads`)

### Out of Scope

- Implementing core detection/loader/command logic itself.
- Documentation wording changes (`bv-7k1w`).

## Assumptions

- Existing test utilities can create temporary repo fixtures for `.tickets` and `.beads`.
- Mode-conditional assertions are acceptable where behavior intentionally differs.

## Open Questions

- Should mixed-mode warning text be assertion-stable or loosely matched to avoid brittle tests?
  - Owner: maintainer
  - Timing: before test finalization
  - Blocking: non-blocking (default: assert key substring only)

## Implementation Spec

- Add focused fixture-driven tests in datasource/loader packages.
- Update robot contract tests to assert mode-specific command hints.
- Add mixed `.tickets` + `.beads` fixture test proving `tk` precedence.
- Ensure legacy Beads tests remain green.

Likely files:

- `internal/datasource/*_test.go`
- `pkg/loader/*_test.go`
- `cmd/bv/main_robot_test.go` and/or `tests/e2e/robot_contract_test.go`
- `pkg/export/markdown_test.go`

## Acceptance Criteria

- Tests explicitly validate dual-mode detection and fallback behavior.
- Tests cover malformed and partial `tk` ticket inputs without crash.
- Robot contract tests pass in both modes with expected command hints.
- `tk`-mode helper commands are asserted to use valid `tk` syntax (`tk ready`, `tk blocked`, `tk show`, `tk start`) without Beads-only `--json` flags.
- `tk`-only e2e test confirms robot commands succeed without `.beads` present.
- Mixed malformed-`tk` + valid Beads fixture confirms deterministic error (no silent fallback).
- Detailed datasource metadata is asserted in mixed-source and single-source fixtures.
- Existing `GenerateMarkdown(...)` behavior is covered by compatibility tests when new options path is introduced.
- Mixed-source diagnostic note assertions are scoped to payloads that expose usage hints/diagnostics.
- No regressions in touched package test suites.

## Validation Plan

- `go test ./internal/datasource/...`
- `go test ./pkg/loader/...`
- `go test ./cmd/bv/...`
- `go test ./pkg/export/...`
- `go test ./tests/e2e/... -run Robot`
- `go build ./...`
- `go vet ./...`
- `gofmt -l .`

Evidence:

- CI/local test logs showing pass status for new and legacy assertions.

## Dependencies

- Upstream:
  - `bv-92o5`
  - `bv-v3f4`
  - `bv-c56w`
- Downstream:
  - `bv-7k1w`

## Artifacts

- New fixtures and tests for dual-mode scenarios.
- Session notes entry with validation command outputs and summary.

## Notes

**2026-02-20T19:56:09Z**

Expanded and validated dual-mode test coverage across datasource, loader, export, cmd, and e2e layers. Added tests for: tk-only robot-next command outputs, tk emit-script command outputs, mixed .tickets + .beads usage hint diagnostics, tickets precedence over beads sources, nested-directory tickets discovery, TICKETS_DIR override, tk-only loading without .beads, and malformed tickets no-fallback behavior.

**2026-02-20T19:56:15Z**

Validation: go test ./internal/datasource/... passed; go test ./pkg/loader/... -run Tickets passed; go test ./pkg/export/... passed; go test ./cmd/bv/... passed; go test ./tests/e2e/... -run 'RobotNextContractActionable|RobotNextContractActionableTK|RobotTriage_MixedSourcesAddsUsageHint|EmitScript' passed; go build ./... passed; go vet ./... passed.
