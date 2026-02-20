# tkv Cutover Decision Lock

Date: 2026-02-20
Ticket: `bv-n8vq`

## 1) Command package path policy

- Decision: keep command package path as `cmd/bv` for now.
- Rationale: minimizes churn in CI/scripts/import paths while still allowing binary/UX rename to `tkv`.

## 2) Module path policy

- Decision: defer `go.mod` module path rename to follow-up ticket after functional cutover is stable.
- Rationale: module rename is high blast-radius and not required for runtime/CLI identity cutover.

## 3) Flag/env policy

- Decision: keep existing `BV_*` and related flags during this cutover; no strict env rename now.
- Rationale: avoid breaking user automation during core rename.

## 4) Correlation/history naming policy (`bead-*`)

- Decision: keep existing flag names for this cutover and document as legacy naming.
- Rationale: behavior stability first; naming cleanup can be done in follow-up once tkv baseline is stable.

## 5) Updater/release policy

- Decision: point release/updater metadata to `adampush/ticket_viewer`.
- Rationale: fork is canonical delivery target now.

## 6) Nix packaging policy

- Decision: add/maintain `.#tkv` package output; keep temporary `.#bv` compatibility alias during transition.
- Rationale: smooth migration for existing Nix workflows.

## 7) Tooling policy (`goreleaser`, `nix`)

- Decision: run local checks when installed; otherwise use CI-equivalent checks and record evidence in ticket notes.
- Rationale: ensures consistent verification despite local tool availability differences.

## 8) Workspace/config namespace policy (`.bv/*`)

- Decision: keep `.bv/*` namespace in this cutover.
- Rationale: prevents config/workspace breakage while runtime/CLI identity changes land.

## 9) User config directory policy (`~/.config/bv`)

- Decision: keep `~/.config/bv` in this cutover.
- Rationale: avoid silent user config migration risk; follow-up ticket can introduce explicit migration plan.
