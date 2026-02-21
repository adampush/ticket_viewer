---
id: bv-a0uk
status: closed
deps: [bv-s1jn]
links: []
created: 2026-02-20T21:43:08Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Rename CLI identity from bv to tkv

## Context / Why

After runtime and command semantics are `tk`-native, the user-facing CLI identity must match the fork direction (`tkv` / `ticket_viewer`) to avoid brand and usage confusion.

## Scope

### In Scope

- Rename user-facing CLI identity from `bv` to `tkv`.
- Apply locked decision for command package path (`cmd/bv` keep/rename).
- Update help/version text and runtime self-references.

### Out of Scope

- CI/release automation rewiring (handled in downstream tickets).

## Assumptions

- Decision lock has finalized path policy and alias policy.

## Open Questions

- Keep temporary `bv` alias window?
  - Owner: Adam Push
  - Timing: before merge
  - Blocking: yes

## Implementation Spec

Likely files:

- `cmd/bv/main.go` (or renamed command path)
- `pkg/version/*`
- docs/help text surfaces tightly coupled to binary name

## Acceptance Criteria

- Primary invocation is `tkv` per chosen policy.
- Help and usage text no longer present ambiguous identity.
- Any temporary alias behavior is documented and tested.

## Validation Plan

- `go test ./cmd/...`
- smoke run of renamed invocation
- `go build ./...`

## Dependencies

- Upstream: `bv-s1jn`
- Downstream: `bv-3eud`, `bv-sds5`, `bv-kwe0`

## Artifacts

- Rename behavior notes with alias/deprecation decision evidence.

## Notes

**2026-02-20T22:21:33Z**

Implemented primary CLI identity rename surfaces to tkv while keeping cmd path policy at cmd/bv. Updated usage/help/version/update messaging in cmd/bv/main.go to reference tkv (e.g., Usage: tkv, tkv --update/--rollback, version banner tkv <version>).

**2026-02-20T22:21:40Z**

Validation: go test ./cmd/bv/... passed; go build ./... passed; go vet ./... passed.
