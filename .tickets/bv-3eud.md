---
id: bv-3eud
status: open
deps: [bv-a0uk, bv-n8vq]
links: []
created: 2026-02-20T21:43:08Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Update module, updater, and release plumbing

## Context / Why

Identity cutover is incomplete until module/updater/release metadata point to the fork identity. This is high blast-radius and must be isolated.

## Scope

### In Scope

- Apply module path policy from decision lock.
- Update updater endpoints/release metadata to fork identity.
- Update goreleaser/build metadata and ldflags path references.

### Out of Scope

- CI workflow and installer script path rewrites (handled in automation ticket).

## Assumptions

- CLI rename ticket has landed or is stable enough for plumbing updates.

## Open Questions

- Exact release channel strategy for fork (stable vs prerelease mapping).
  - Owner: Adam Push
  - Timing: before merge
  - Blocking: yes

## Implementation Spec

Likely files:

- `go.mod`
- `.goreleaser.yaml`
- `pkg/updater/*`
- `pkg/version/*`

## Acceptance Criteria

- Module/release/updater metadata all point to fork identity.
- ldflags version injection remains correct.
- `goreleaser check` passes (or CI equivalent with evidence).

## Validation Plan

- `go test ./pkg/updater/...`
- `go build ./...`
- `goreleaser check` (or CI fallback evidence)

## Dependencies

- Upstream: `bv-a0uk`, `bv-n8vq`
- Downstream: `bv-8a55`, `bv-kwe0`

## Artifacts

- Release plumbing verification notes and command outputs.
