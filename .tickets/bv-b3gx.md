---
id: bv-b3gx
status: closed
deps: []
links: []
created: 2026-02-20T21:42:49Z
type: epic
priority: 1
assignee: Adam Push
---
# Execute tkv rename and tk-only cutover

## Context / Why

We now control a fork (`adampush/ticket_viewer`) and want to move from a Beads-branded codebase to a `tk`-only product with clear identity (`ticket_viewer` / `tkv`). This epic coordinates the full cutover with minimal regression risk.

## Scope

### In Scope

- Lock policy decisions required for safe rename/cutover.
- Remove remaining Beads runtime behavior and normalize to `tk` workflows.
- Rename CLI/product identity and align module/release/automation/docs surfaces.
- Run compatibility sweep and finalize with evidence.

### Out of Scope

- New non-cutover features.
- Analytics algorithm redesign.
- Historical Beads data migration tooling beyond explicit rename/cutover needs.

## Assumptions

- This work lands on the fork and does not need upstream acceptance.
- Temporary aliases are allowed only if explicitly approved in decision-lock ticket.

## Open Questions

- None at epic level; all decision items are owned by `bv-n8vq`.

## Implementation Spec

- Execute child tickets in dependency order.
- Require validation evidence in each child before closure.
- Preserve one active implementation ticket at a time.

## Acceptance Criteria

- All child tickets are closed with validation evidence.
- Product behavior and docs are consistent with `tk`-only `tkv` identity.
- Final compatibility sweep is green and recorded.

## Validation Plan

- Aggregate all child validation artifacts.
- Re-run final gates from `bv-kwe0` before epic closure.

## Dependencies

- Upstream: none.
- Downstream: `bv-n8vq`, `bv-yfa0`, `bv-s1jn`, `bv-a0uk`, `bv-3eud`, `bv-8a55`, `bv-sds5`, `bv-kwe0`.

## Artifacts

- Updated code, docs, release/automation config, and ticket notes.
- Session notes in `DEVELOPMENT/sessions/YYYY-MM-DD.md`.

## Notes

**2026-02-21T00:56:24Z**

All planned cutover children are complete, including final compatibility sweep and full test stabilization. Project now passes full build/vet/test validation under tk-first cutover branch state.
