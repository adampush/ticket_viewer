---
id: bv-s1jn
status: closed
deps: [bv-yfa0, bv-n8vq]
links: []
created: 2026-02-20T21:43:08Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Normalize commands to tk across robot and export

## Context / Why

User-facing and agent-facing command snippets still mix `br` semantics. In a `tk`-only product, command output must be consistently valid `tk` syntax.

## Scope

### In Scope

- Normalize robot command helpers and grouped claim commands to `tk` syntax.
- Normalize emit-script output to `tk` syntax.
- Normalize export markdown command snippets to `tk` syntax.
- Remove unsupported Beads-only flags from helper commands.

### Out of Scope

- Product binary rename (`bv` -> `tkv`).
- Module/release plumbing changes.

## Assumptions

- Runtime source behavior from `bv-yfa0` is in place.

## Open Questions

- Should any examples keep historical `br` snippets in non-runtime docs?
  - Owner: Adam Push
  - Timing: before docs ticket starts
  - Blocking: non-blocking for this code ticket

## Implementation Spec

Likely files:

- `pkg/analysis/triage.go`
- `cmd/bv/main.go`
- `pkg/export/markdown.go`
- related tests in `tests/e2e` and `pkg/export`

## Acceptance Criteria

- Robot/emit-script/export command snippets are `tk`-valid.
- No Beads-only command flags in tk command outputs.
- Contract tests assert expected command strings.

## Validation Plan

- `go test ./pkg/analysis/...`
- `go test ./pkg/export/...`
- `go test ./cmd/bv/...`
- targeted e2e robot/script tests

## Dependencies

- Upstream: `bv-yfa0`, `bv-n8vq`
- Downstream: `bv-a0uk`, `bv-kwe0`

## Artifacts

- Command normalization notes and test evidence in ticket notes.

## Notes

**2026-02-20T22:19:27Z**

Normalized command outputs to tk for active runtime surfaces: robot-next claim/show now emit tk commands; emit-script output emits tk claim/show; e2e contract tests updated to tk fixtures and tk command expectations; cmd runtime no-issues message updated to tk create guidance. Export markdown tracker-aware tk commands remain enabled via MarkdownOptions tracker mode and CLI export path continues to pass tracker mode.

**2026-02-20T22:19:36Z**

Validation: go test ./pkg/analysis/... passed; go test ./pkg/export/... passed; go test ./cmd/bv/... passed; go test ./tests/e2e/... -run 'RobotNextContractActionable|RobotNextContractActionableTK|EmitScript' passed; go build ./... passed; go vet ./... passed.
