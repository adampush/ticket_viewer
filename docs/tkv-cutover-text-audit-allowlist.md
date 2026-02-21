# tkv Cutover Text Audit Allowlist

This document records intentional legacy references discovered during cutover text-audit sweeps.

## Scope

- Audit date: 2026-02-20
- Ticket: `bv-kwe0`
- Primary check: regex scans for legacy CLI command strings (`bv --`) and naming references.

## Allowlisted Matches

1. `README.md`
   - Remaining `bv --...` examples are kept in deep historical/legacy sections.
   - Rationale: preserve historical context while top-level quickstart/install/help sections now use `tkv` and tk-first semantics.
   - Mitigation: explicit historical-note banner near the README top clarifies legacy naming retention.

2. Legacy term references (`bead`, `beads`) in flag names and API surfaces
   - Rationale: correlation and related robot flags remain bead-named by explicit cutover decision lock (`docs/tkv-cutover-decisions.md`).
   - Examples: `--robot-file-beads`, `--bead-history`, `--suggest-bead`.

## Non-Allowlisted Matches

- `cmd/bv/main.go`: no remaining `bv --` command examples in help output strings.
- `pkg/ui/*.go`: no remaining `bv --` tutorial/help snippets.
- `AGENTS.md`: canonical command examples now use `tkv`.
