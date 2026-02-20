---
id: bv-yfa0
status: open
deps: [bv-n8vq]
links: []
created: 2026-02-20T21:43:08Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Cut over runtime to tk-only sources

## Context / Why

Runtime still contains Beads fallback and Beads-specific loading assumptions. For a true `tk`-only product, runtime source model must standardize on `.tickets/*.md`.

## Scope

### In Scope

- Remove Beads runtime source selection/loading paths from primary execution paths.
- Keep/adjust error messaging to `tk`-first guidance.
- Ensure malformed/missing `.tickets` behavior is deterministic and explicit.

### Out of Scope

- CLI branding rename (`bv` -> `tkv`).
- Release/automation/docs rewrites.

## Assumptions

- Decision-lock ticket defines whether any temporary compatibility path remains.

## Open Questions

- Should any read-only Beads fallback remain behind an explicit compatibility flag?
  - Owner: Adam Push
  - Timing: before merge of this ticket
  - Blocking: yes (unless explicitly set to "no fallback")

## Implementation Spec

- Touch datasource/loader runtime entrypoints and remove Beads-default assumptions.
- Update callsites that expect Beads-path side effects.

Likely files:

- `internal/datasource/*.go`
- `pkg/loader/*.go`
- `cmd/bv/main.go` (runtime load path handling)

## Acceptance Criteria

- Runtime loads from `.tickets` only (or explicit compatibility behavior per locked decision).
- No implicit Beads fallback in primary runtime path.
- Error messages reference ticket-based setup.

## Validation Plan

- `go test ./internal/datasource/...`
- `go test ./pkg/loader/...`
- targeted e2e load-path checks
- `go build ./...`

## Dependencies

- Upstream: `bv-n8vq`
- Downstream: `bv-s1jn`, `bv-kwe0`

## Artifacts

- Runtime cutover notes and validation output in ticket notes.
