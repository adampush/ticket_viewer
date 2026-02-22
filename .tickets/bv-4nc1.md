---
id: bv-4nc1
status: open
deps: [bv-hqio]
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 2
assignee: Adam Push
parent: bv-jzv8
---
# Finalize tkv packaging and distribution channels

## Context / Why

This is Workstream 6 from `docs/tkv-native-cleanup-plan.md`. Distribution channels and installer/release metadata must consistently represent `tkv` and the migrated namespace.

## Scope

### In Scope

- Align packaging/release metadata with `tkv` naming.
- Update installer/help output and release workflow references.
- Validate Homebrew tap install path (`adampush/tap/tkv`) end-to-end.

### Out of Scope

- Core runtime behavior changes.
- New feature implementation unrelated to naming/distribution consistency.

## Ticket Type & Granularity

- Type: Task
- Granularity target: 1-3 commits focused on release/install surface.

## Prerequisites & Dev State

- Upstream module/import migration `bv-hqio` complete.
- Current release and installer scripts available in-repo.

## Assumptions

- Distribution channels are authoritative and must match docs exactly after this ticket.

## Open Questions

- Is automated tap verification available in current CI, or must this run manually for now?
  - Owner: Adam Push
  - Timing: before close
  - Blocking: no (manual evidence allowed with follow-up automation ticket)

## Implementation Guidance

- Update `.goreleaser.yaml`, `install.sh`, `install.ps1`, and relevant release workflow metadata.
- Ensure help/install output references `tkv` consistently.
- Inherit NFR constraints: no installer security regressions and no accidental channel breakage.

Likely files:

- `.goreleaser.yaml`
- `install.sh`
- `install.ps1`
- `.github/workflows/*.yml` (release/install related)

## Acceptance Criteria

- Installer output and release metadata are consistently `tkv`.
- `brew install adampush/tap/tkv` path is validated and documented.
- Documentation references match actual distribution channels.

## Deployment & State

- Breaking/Non-breaking: potentially breaking if external automation depended on old channel names.
- Migration requirements: explicit release notes with updated install commands.
- Safe to deploy independently: yes, after Workstream 5.

## Validation Plan

- installer dry-runs (`install.sh`, `install.ps1`) where environment allows
- release metadata lint/checks in current tooling
- tap install verification (`brew install adampush/tap/tkv`) in a clean environment
- `go build ./...` sanity check

Expected evidence:

- installer logs + tap verification logs
- PR checklist with before/after install command evidence
- `tk add-note bv-4nc1 "validation: ..."` with results

## Dependencies

- Upstream: `bv-hqio`
- Downstream: `bv-y61k`

## Artifacts

- Updated install docs examples
- Distribution validation logs
- Any manual-only validation follow-up ticket IDs
