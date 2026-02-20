---
id: bv-0kqn
status: closed
deps: []
links: []
created: 2026-02-20T13:03:37Z
type: epic
priority: 1
assignee: Adam Push
---
# Dual-mode tracker support (beads + tk)

## Context / Why

`bv` analytics are strong but currently tied to Beads-centric discovery and command hints. Supporting `tk` directly allows teams using `.tickets/` to use the same triage engine without running migration flows.

This epic delivers dual-mode tracker support with deterministic auto-detect and minimal blast radius.

## Scope

### In Scope

- Add `tk` tracker ingestion path (markdown + frontmatter).
- Auto-detect tracker mode (`tk` preferred when both trackers exist for MVP).
- Keep core analysis logic unchanged by adapting at loader/datasource boundaries.
- Make robot command hints tracker-aware (`br` vs `tk`).
- Add validation coverage for detection, parsing, and robot contracts.
- Update user/developer docs to reflect dual-mode behavior.

### Out of Scope

- Re-architecting analysis/scoring packages.
- Full history/correlation parity for `tk` in MVP.
- Major TUI redesign.
- New persistent config system for tracker selection in MVP.

## Assumptions

- `tk` ticket markdown has parseable YAML frontmatter in most files.
- Existing `model.Issue` is sufficient for `tk` field mapping.
- In mixed repos, `.tickets` precedence is acceptable for MVP.

## Open Questions

- Should explicit `--tracker` override land in MVP or immediate follow-up?
  - Owner: product/maintainer
  - Timing: before or immediately after MVP merge
  - Blocking: non-blocking for MVP (default is follow-up)

## Implementation Spec

- Execute child tickets in dependency order.
- Keep tracker-specific logic isolated to datasource/loader and command hint generation.
- Enforce deterministic behavior and edge-case handling from extension plan.

## Acceptance Criteria

- All child tickets complete and closed with validation evidence.
- `bv --robot-next`, `--robot-triage`, `--robot-plan`, `--robot-insights` operate in `tk` repos.
- Beads behavior remains compatible.
- Documentation reflects shipped behavior.

## Validation Plan

- Aggregate child-ticket validation outputs.
- Run final integration sweep:
  - `go build ./...`
  - `go vet ./...`
  - `gofmt -l .`
  - `go test ./internal/datasource/... ./pkg/loader/... ./pkg/export/... ./cmd/bv/...`

## Dependencies

- Upstream: none (epic root)
- Downstream:
  - `bv-92o5`
  - `bv-v3f4`
  - `bv-c56w`
  - `bv-g3nq`
  - `bv-7k1w`

## Artifacts

- Implementation changes in datasource/loader/CLI output files.
- Updated docs including README and robot docs/help text.
- Session notes in `DEVELOPMENT/sessions/YYYY-MM-DD.md`.

## Notes

**2026-02-20T20:00:04Z**

Completed child tickets bv-92o5, bv-v3f4, bv-c56w, bv-g3nq, and bv-7k1w. Delivered dual-mode source detection with tk precedence, tk markdown loader mapping to model.Issue, tracker-aware robot/emit-script/export command hints, expanded dual-mode/e2e test coverage, and docs/help/robot-docs updates with verification evidence.
