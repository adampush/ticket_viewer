---
id: bv-kwe0
status: open
deps: [bv-yfa0, bv-s1jn, bv-a0uk, bv-3eud, bv-8a55, bv-sds5]
links: []
created: 2026-02-20T21:43:09Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Run compatibility sweep and finalize cutover

## Context / Why

Cutover touches many surfaces. Final integration sweep is required to verify no stale identity strings, broken paths, or inconsistent policies remain.

## Scope

### In Scope

- Run full build/test/vet/format checks.
- Run targeted package checks from cutover plan.
- Run text-audit regex checks and resolve/allowlist matches.
- Confirm release/install/nix/tooling checks and evidence.

### Out of Scope

- New implementation features unrelated to resolving cutover defects.

## Assumptions

- All upstream implementation tickets are merged first.

## Open Questions

- Any unresolved audit match is blocking unless allowlisted with rationale.
  - Owner: Adam Push
  - Timing: before closure
  - Blocking: yes

## Implementation Spec

- Execute validation gates defined in `docs/tkv-fork-rename-cutover-plan.md`.
- Produce tracked allowlist artifact for intentional historical references.
- File follow-up tickets for any deferred non-blocking cleanup.

## Acceptance Criteria

- Validation gates pass (or approved CI/tool fallback evidence recorded).
- Text-audit results are either fixed or allowlisted with rationale.
- Final release/install/packaging surfaces are coherent.

## Validation Plan

- `go build ./...`
- `go vet ./...`
- `go test ./...`
- targeted package checks from plan
- text-audit `rg` command set from plan

## Dependencies

- Upstream: `bv-yfa0`, `bv-s1jn`, `bv-a0uk`, `bv-3eud`, `bv-8a55`, `bv-sds5`
- Downstream: none

## Artifacts

- Final validation report in ticket notes.
- Text-audit allowlist artifact with per-match rationale.
- Session handoff summary.
