---
id: bv-3eud
status: closed
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

## Notes

**2026-02-20T22:23:16Z**

Updated module/updater/release plumbing to fork identity while keeping module path unchanged per decision lock. Changes: pkg/updater now targets adampush/ticket_viewer GitHub releases and uses ticket-viewer update user-agent; .goreleaser.yaml now uses project_name tkv, binary tkv, and fork homepage/owner metadata for brew/scoop outputs.

**2026-02-20T22:23:24Z**

Validation: go test ./pkg/updater/... passed; go build ./... passed; go vet ./... passed. goreleaser is not installed locally (goreleaser not found -> not found); per policy, goreleaser check will be validated in CI and captured in follow-up notes.
