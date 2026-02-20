---
id: bv-7k1w
status: closed
deps: [bv-c56w, bv-g3nq]
links: []
created: 2026-02-20T13:03:49Z
type: task
priority: 2
assignee: Adam Push
parent: bv-0kqn
---
# Update docs and verify dual-mode behavior documentation

## Context / Why

Dual-mode support is user-facing. Without synchronized docs, users will run wrong commands and lose trust in robot hints.

## Scope

### In Scope

- Update docs to reflect `beads` + `tk` support and mode precedence.
- Add or update examples for `tk` robot usage.
- Run documentation verification checklist against runtime behavior.

### Out of Scope

- Core loader/datasource implementation.
- New product features beyond documentation and docs validation.
- Functional command-generation changes in exports/robot payloads (owned by `bv-c56w`).

## Assumptions

- Final command-hint behavior from `bv-c56w` is settled.
- `bv --robot-docs` topics are the authoritative machine-readable doc surface.

## Open Questions

- Should README include a dedicated migration guide from Beads-only to dual-mode usage?
  - Owner: maintainer
  - Timing: before closing ticket
  - Blocking: non-blocking (default: concise section, not full migration guide)

## Implementation Spec

- Update top-level README sections for dual-mode tracker support.
- Update help/robot docs wording and examples where Beads-only assumptions exist.
- Validate docs with runtime spot checks from checklist.

Likely files:

- `README.md`
- `cmd/bv/main.go` help/robot-docs text blocks
- `docs/tk-dual-mode-extension-plan.md` (if checklist wording needs final sync)

## Acceptance Criteria

- Documentation clearly states dual-mode support and precedence behavior.
- `tk` examples are executable and accurate.
- Robot docs and help text match implemented behavior.
- Documentation verification checklist is completed with evidence.

## Validation Plan

- `go run ./cmd/bv --help`
- `go run ./cmd/bv --robot-docs guide`
- `go run ./cmd/bv --robot-docs commands`
- `go run ./cmd/bv --robot-next` (in `tk` fixture repo)
- `go run ./cmd/bv --robot-next` (in Beads fixture repo)

Plus standard checks if code/docs text in Go sources changed:

- `go test ./cmd/bv/...`
- `go build ./...`
- `gofmt -l .`

Evidence:

- Captured command outputs/snippets showing docs parity.

## Dependencies

- Upstream:
  - `bv-c56w`
  - `bv-g3nq`
- Downstream: none

## Artifacts

- Updated docs and help text.
- Session notes entry documenting docs parity verification.

## Notes

**2026-02-20T19:58:52Z**

Updated user/operator docs and help text for dual-mode behavior. Changes include: README Tracker Auto-Detection section with tk/beads precedence and command examples; CLI --help wording updated to 'ticket graphs'; robot docs guide data_source/description updated to reflect auto-detected tracker sources and .tickets precedence; robot/export help wording switched from hardcoded br instructions to tracker-aware wording.

**2026-02-20T19:58:58Z**

Validation: go run ./cmd/bv --help executed (updated wording present), go run ./cmd/bv --robot-docs guide executed (updated data_source/description), go run ./cmd/bv --robot-docs commands executed, go run ./cmd/bv --robot-next verified tk output with TICKETS_DIR fixture, go run ./cmd/bv --robot-next verified beads output with BEADS_DIR + empty TICKETS_DIR fixture. Also go test ./cmd/bv/... passed; go build ./... and go vet ./... passed.
