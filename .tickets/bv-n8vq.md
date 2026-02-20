---
id: bv-n8vq
status: closed
deps: []
links: []
created: 2026-02-20T21:43:08Z
type: task
priority: 1
assignee: Adam Push
parent: bv-b3gx
---
# Lock pre-cutover rename decisions

## Context / Why

The cutover has multiple policy forks (path/module/env/aliasing). Without explicit decisions first, implementation will drift and produce contradictory behavior.

## Scope

### In Scope

- Lock and document decisions for:
  - `cmd/bv` vs `cmd/tkv` path policy
  - module path rename timing
  - flag/env alias strategy and removal timeline
  - `bead-*` correlation naming policy
  - updater/release channel/repo policy
  - nix attr naming policy (`.#bv` vs `.#tkv`)
  - workspace/config namespace policy (`.bv/*`, `~/.config/bv`)

### Out of Scope

- Code implementation of these decisions.

## Assumptions

- Maintainer/owner can provide final policy choices in this ticket.

## Open Questions

- Each policy bullet above is initially open.
  - Owner: Adam Push
  - Timing: before starting dependent implementation tickets
  - Blocking: yes

## Implementation Spec

- Create a decision record in this ticket notes (or linked markdown) with chosen option + rationale + any temporary alias sunset date.
- Link decisions to impacted tickets.

## Acceptance Criteria

- All required decision gates are resolved and recorded.
- Dependent tickets reference these locked decisions.

## Validation Plan

- Manual review checklist confirming each gate has a resolved value and owner rationale.

## Dependencies

- Upstream: none
- Downstream: `bv-yfa0`, `bv-s1jn`, `bv-3eud`

## Artifacts

- Decision log artifact (ticket note or markdown) with per-decision rationale.

## Notes

**2026-02-20T22:02:06Z**

Locked all pre-cutover decisions and recorded them in docs/tkv-cutover-decisions.md: cmd path stays cmd/bv; module rename deferred; BV_* envs retained; bead-* flag naming retained for now; updater/release target set to adampush/ticket_viewer; nix adds .#tkv with temporary .#bv alias; tooling fallback to CI evidence when goreleaser/nix absent; .bv and ~/.config/bv namespaces retained in this cutover.
