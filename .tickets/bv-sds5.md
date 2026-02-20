---
id: bv-sds5
status: open
deps: [bv-a0uk, bv-8a55]
links: []
created: 2026-02-20T21:43:08Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Update docs and robot docs for tkv

## Context / Why

Docs/help are part of product behavior. After rename/cutover, stale wording or commands causes user/operator errors.

## Scope

### In Scope

- Update README, AGENTS docs, robot docs/help, tutorial/help content.
- Align examples and command text to `tkv`/`tk` semantics.
- Preserve explicit historical notes only where intentional.

### Out of Scope

- New runtime features.

## Assumptions

- CLI/automation rename tickets have settled command/path names.

## Open Questions

- Level of historical references retained in docs.
  - Owner: Adam Push
  - Timing: before merge
  - Blocking: non-blocking if allowlist policy is documented

## Implementation Spec

Likely files:

- `README.md`
- `AGENTS.md`
- `cmd/bv/main.go` help/robot-docs text blocks (or renamed cmd path)
- `pkg/ui/tutorial*`
- `pkg/agents/*` blurb/template text

## Acceptance Criteria

- User-facing docs/help text matches runtime behavior.
- Robot docs are consistent with renamed identity and command syntax.
- Historical references are intentional and allowlisted.

## Validation Plan

- `go run ./cmd/... --help`
- `go run ./cmd/... --robot-docs guide`
- `go run ./cmd/... --robot-docs commands`
- targeted docs parity spot checks

## Dependencies

- Upstream: `bv-a0uk`, `bv-8a55`
- Downstream: `bv-kwe0`

## Artifacts

- Docs parity checklist evidence and allowlist notes.
