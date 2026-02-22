# tkv-Native Cleanup Plan

Ticket: `bv-jzv8`

## Goal

Complete the transition from legacy `beads_viewer` naming to a fully `tkv`-native project identity across CLI, robot APIs, docs, code terminology, config/env namespaces, module path, and release tooling.

## User Outcomes and Success Criteria

### Primary user flows

1. **CLI users** run `tkv --help` and all examples without seeing stale bead-era naming.
2. **Automation users** run `--robot-*` commands and receive ticket-native keys that match published schema/docs.
3. **CI/integration users** consume env/config defaults using only `TKV_*` variables.
4. **Installer users** install/upgrade via documented channels and receive consistent `tkv` naming.

### Observable success criteria (binary)

- `tkv --help`, `tkv --robot-help`, and `tkv --robot-docs all` contain no legacy bead terms except explicitly marked historical notes.
- `tkv --robot-schema` matches emitted payload keys for all robot commands under tests.
- No active code path reads `BV_*`; all active env/config reads use `TKV_*`.
- No `.go` import references `github.com/Dicklesworthstone/beads_viewer`.
- `go build ./...`, `go vet ./...`, and `go test ./...` pass on `main` after merge.

## Scope

In scope:

- CLI/robot command and flag naming
- Robot JSON schema and payload field naming
- User-facing docs/help/tutorial content
- Internal naming in core packages where legacy terms leak into behavior or public output
- Environment variable and config key namespace migration (`BV_*` -> `TKV_*`)
- Go module path migration to fork namespace
- Release/install surface consistency (including Homebrew tap readiness)

Out of scope:

- Backward compatibility shims for renamed flags/env vars (explicitly avoided)
- Non-essential refactors unrelated to naming consistency

## Assumptions Register

| ID | Assumption | Confidence | Validation |
|---|---|---|---|
| A1 | Breaking changes are acceptable for CLI/env/schema names in this phase. | High | Release note includes explicit breaking-change section; no compatibility code added. |
| A2 | `ticket` is canonical terminology for all public interfaces. | High | Grep-based checks for bead-era naming in help/docs/schemas. |
| A3 | Existing tests are sufficient to catch most behavior regressions from naming migration. | Medium | Full build/vet/test + targeted e2e robot contract checks. |
| A4 | Module path migration to `github.com/adampush/ticket_viewer` is authoritative and final. | High | `go.mod` + import grep checks + clean `go test ./...`. |
| A5 | Homebrew distribution target remains `adampush/tap/tkv`. | Medium | End-to-end install test from clean environment. |

## Locked Defaults (Resolved Decisions)

1. **No compatibility wrappers** for old flags/env vars/schema fields.
2. **Failure semantics for legacy env vars:** old `BV_*` values are ignored; behavior follows defaults unless corresponding `TKV_*` is set.
3. **Precedence order:** CLI flags > `TKV_*` env vars > config file defaults.
4. **Historical naming allowance:** only in explicitly marked migration/historical docs sections.
5. **Module cutover policy:** single-phase hard cut; all internal imports updated in same PR as `go.mod` change.

## Interface Contracts Matrix (Old -> New)

This plan requires explicit interface-contract tracking for all externally visible renames. The table below defines contract classes and required behavior. Each workstream PR must include its exact old->new key/flag mapping as validation evidence.

| Workstream | Surface | Old | New | Legacy Behavior | Validation |
|---|---|---|---|---|---|
| 1 | CLI flags/help text | bead-era flag/term | ticket-native equivalent | old names hard-fail | `tkv --help`, `tkv --robot-help`, docs parity checklist |
| 2 | Robot payload + schema keys | bead-named payload keys | ticket-native payload keys | old keys removed from active payload | `tkv --robot-schema` + contract tests |
| 3 | User-visible internal terminology | mixed bead/ticket terms in primary paths | ticket-native terminology | no compatibility aliasing | targeted package tests + output snapshot checks |
| 4 | Environment/config namespace | `BV_*` vars | `TKV_*` vars | legacy `BV_*` ignored, defaults apply | grep checks + env-driven tests |
| 5 | Module/import namespace | `github.com/Dicklesworthstone/beads_viewer` | `github.com/adampush/ticket_viewer` | partial migration invalid | full import grep + build/test |
| 6 | Packaging/install metadata | legacy project naming in install/release text | `tkv` naming | old naming removed from active channels | installer dry-runs + tap install verification |

## Design Principles

1. No compatibility wrappers: do direct cutover.
2. One semantic model: `ticket` is canonical domain term.
3. Public interface first: robot/API/CLI/docs must match before internals are considered complete.
4. Deterministic validation at each phase with full `go build`, `go vet`, and `go test ./...`.
5. Minimize blast radius: isolate each workstream into narrow PRs with explicit file boundaries.

## Architecture and File-Level Plan

### Workstream 1: Public CLI and Robot Surface Rename

**Boundary:** public command/flag/help text only; no deep algorithm changes.

Primary touchpoints:

- `cmd/bv/main.go` (flags, usage text, examples)
- robot docs/help emitters (in `cmd/bv`)
- tests validating CLI/help content

Acceptance criteria:

- No public help text references bead terms unless explicitly historical.
- No bead-named flags remain in active CLI surface.

### Workstream 2: Robot Payload and Schema Normalization

**Boundary:** payload key names + schema contract; preserve semantic meaning of values.

Primary touchpoints:

- robot output builders in `cmd/bv` and `pkg/ui`
- schema emitters + contract tests in `tests/e2e`

Acceptance criteria:

- Robot payload keys are ticket-native and documented.
- `--robot-schema` aligns with emitted JSON in all commands.

### Workstream 3: Internal Terminology Cleanup

**Boundary:** identifiers/types in core code paths that affect readability or user-visible output.

Primary touchpoints:

- `pkg/analysis`, `pkg/correlation`, `pkg/export`, `pkg/ui`, `pkg/loader`

Acceptance criteria:

- Core code paths use ticket-native naming conventions.
- No confusing mixed terminology in primary execution paths.

### Workstream 4: Env/Config Namespace Migration

**Boundary:** env/config surface only; no behavior changes beyond name migration.

Primary touchpoints:

- env lookups across `cmd/bv`, `pkg/ui`, `pkg/export`, `pkg/agents`, `internal/datasource`
- docs: `README.md`, `AGENTS.md`
- test env setup in `tests/e2e` and package tests

Acceptance criteria:

- All documented env vars are `TKV_*`.
- Legacy `BV_*` names removed from active codepaths.

### Workstream 5: Module Path and Import Migration

**Boundary:** module declaration + import rewrites only.

Primary touchpoints:

- `go.mod`
- all `.go` files with module imports
- generated/version references that embed module path

Acceptance criteria:

- `go build ./...` and `go test ./...` pass with new module path.
- No imports reference `github.com/Dicklesworthstone/beads_viewer`.

### Workstream 6: Packaging and Distribution Finalization

**Boundary:** release/install metadata and scripts; no unrelated product changes.

Primary touchpoints:

- `.goreleaser.yaml`, `install.sh`, `install.ps1`, release workflows
- tap repo and formula pipeline

Acceptance criteria:

- `brew install adampush/tap/tkv` works end-to-end.
- Installer output and docs match real distribution channels.

## Edge Cases and Failure Semantics

1. **Mixed env modes (`BV_*` + `TKV_*`)**
   - Behavior: only `TKV_*` is honored; `BV_*` ignored.
   - Handling: deterministic defaulting; no warning spam unless explicitly added in a later task.

2. **Malformed/unknown robot consumer expectations**
   - Behavior: emit only normalized schema-compliant keys.
   - Handling: contract tests enforce emitted payload/schema parity.

3. **Empty/partial ticket datasets**
   - Behavior: existing deterministic empty-state outputs preserved.
   - Handling: non-regression tests verify unchanged semantics.

4. **Old automation scripts using legacy names**
   - Behavior: legacy CLI flags/schema keys hard-fail; legacy `BV_*` env vars are ignored and defaults apply.
   - Handling: release notes + migration examples are mandatory to prevent silent misconfiguration.

5. **Module path partial migration**
   - Behavior: treated as invalid intermediate state.
   - Handling: single PR with full import rewrite and full test gates required.

## Blast Radius Controls

- One active high-risk stream per PR (module migration and packaging always isolated).
- Avoid refactors in algorithmic/analysis logic during naming-only streams.
- Preserve stable subsystems (`pkg/analysis` algorithms, scheduler/math behavior) unless naming leaks into output.
- Require green gates before moving to next stream.

## Testing Matrix (Required Minimum)

| Change Type | Required Coverage | Additional Requirement | Evidence |
|---|---|---|---|
| Implementation behavior change | unit + integration | happy-path, edge-case, failure-path | test output + `tk` note/PR checklist |
| Public interface change (CLI/API/schema/env/config) | unit + integration + e2e/non-regression | explicit contract old->new verification | command output + schema/help snapshots |
| Migration/rename/refactor | non-regression + targeted changed-behavior tests | unchanged-behavior assertions must be listed | regression test evidence + grep checks |

Manual-only validation is allowed only when automation is infeasible; rationale and follow-up automation ticket are required.

## Execution Order

1. Public CLI and robot surface rename
2. Robot payload/schema normalization
3. Internal terminology cleanup
4. Env/config namespace migration
5. Module path migration
6. Packaging/distribution finalization

## Ownership, Sequencing, and Checkpoints

| Workstream | Owner | Entry Gate | Exit Gate | Checkpoint Artifact |
|---|---|---|---|---|
| 1) CLI/robot surface | Adam Push | Plan approved + scope frozen | Help/docs free of legacy terms in active surface | PR checklist + `tk add-note` with command outputs |
| 2) Robot payload/schema | Adam Push | Workstream 1 complete | schema and emitted payload parity proven | contract test results + schema snapshot |
| 3) Internal terminology | Adam Push | Workstream 2 complete | no mixed terminology in primary execution paths | package test evidence + reviewer checklist |
| 4) Env/config migration | Adam Push | Workstream 3 complete | active codepaths use `TKV_*`; legacy behavior matches defaults policy | grep evidence + env tests |
| 5) Module/import migration | Adam Push | Workstream 4 complete | no legacy module imports; full suite green | `go build`/`go vet`/`go test` evidence |
| 6) Packaging/distribution | Adam Push | Workstream 5 complete | install/release channels consistent and verified | installer/tap verification evidence |

## Validation Plan

### Global gates (required for every stream)

- `go build ./...`
- `go vet ./...`
- `go test ./...`
- expected pass signal: exit code `0` for all commands
- evidence location: `tk add-note <id> "validation: ..."` and PR validation checklist

### Stream-specific validation

1. **CLI/robot surface**
   - `tkv --help`
   - `tkv --robot-help`
   - `tkv --robot-docs all`
   - expected pass signal: output contains only ticket-native active terms
   - evidence location: help/docs snapshots in PR notes

2. **Robot payload/schema**
   - `tkv --robot-schema`
   - e2e contract tests for robot payload key names
   - expected pass signal: emitted payload keys match schema contract
   - evidence location: contract test output + schema diff/check report

3. **Env/config migration**
   - grep checks for `\bBV_[A-Z0-9_]+\b` in active code/docs
   - test updates proving `TKV_*` variables drive behavior
   - note: for migration gating, "active codepaths" excludes test-only fixtures and historical/session notes
   - expected pass signal: active code/docs checks pass with no legacy env matches; env tests pass
   - evidence location: grep result log + targeted test output

4. **Module path migration**
   - grep checks for `github.com/Dicklesworthstone/beads_viewer`
   - full rebuild/test after import rewrite
   - expected pass signal: no legacy import matches in active codepaths + full suite passes
   - evidence location: grep result log + build/vet/test output

5. **Packaging/distribution**
   - installer dry-runs where available
   - Homebrew tap install verification
   - expected pass signal: install commands complete successfully and output matches docs
   - evidence location: installer/tap logs attached to PR checklist

## Unchanged Behavior Verification (Non-Regression)

The following behaviors must remain stable unless explicitly changed by a workstream contract:

1. Graph-analysis semantics and recommendation ranking behavior (outside intentional key/label renames).
2. Deterministic robot command structure for unchanged fields.
3. Existing empty/partial dataset handling behavior.
4. Browser safety gates for test mode (`no browser` behavior in tests).

Verification requirements:

- run `go test ./...` and targeted e2e contract suites
- compare representative `--robot-*` outputs for unchanged sections before/after each stream
- document unchanged-behavior checks in `tk` notes and PR checklist

## Documentation Deliverables and Verification Checklist

Required updates per impacted stream:

- `README.md` examples and env variable tables
- `AGENTS.md` command/env references
- CLI help text (`--help`, `--robot-help`)
- robot docs/schema (`--robot-docs`, `--robot-schema`)
- release notes migration section (breaking changes + before/after examples)

Docs parity checklist (must all pass):

- Help text names match implemented flags/commands.
- Robot docs/schema match emitted payload keys.
- README examples execute as written.
- Env variable names in docs match active code lookups.

## Risks and Mitigations

1. Broad rename churn introduces regressions.
   - Mitigation: phase-by-phase PRs with narrow scope and full test gates.

2. External automation/scripts depend on old names.
   - Mitigation: explicit breaking-change release notes + migration examples.

3. Module path migration breaks tooling/integrations.
   - Mitigation: isolate migration in dedicated PR and revalidate all workflows.

4. Docs drift from behavior during fast iteration.
   - Mitigation: docs parity checklist is required acceptance criteria.

## Rollout and Recovery Strategy

### Rollout

1. Land workstreams in order with focused PRs.
2. Publish breaking-change notes at each externally visible interface change.
3. Merge module-path and packaging changes only after prior streams are stable.

### Recovery

1. If regression appears, revert the specific stream PR (not unrelated streams).
2. Restore last known-good release tag while fix is prepared.
3. Publish hotfix notes with exact mitigation and revised migration guidance.

Temporary constraints during rollout:

- No concurrent large refactors in touched files.
- No unrelated behavior changes bundled with naming PRs.

## PR Strategy

- Use a sequence of focused PRs aligned to workstreams (one or two workstreams per PR max).
- Each PR includes:
  - explicit breaking-change note (if applicable)
  - updated docs for changed interface
  - validation evidence in PR description
  - checklist outcomes for docs/help/schema parity

## Ticketization Rules (`bv-jzv8` Umbrella)

`bv-jzv8` is the umbrella planning ticket for the full cleanup effort. Implementation executes through child tickets derived from this plan.

Ticket graph requirements:

1. Create child tickets aligned to workstreams (or safe sub-splits for high-risk streams).
2. Model dependencies to match execution order using `tk dep`.
3. Keep one active implementation ticket in progress at a time unless explicit parallelization is planned.

Each child ticket must include:

- Context/Why
- Scope (in/out)
- Assumptions
- Open questions (owner + timing + blocking status)
- Implementation spec (likely files/modules/contracts)
- Acceptance criteria (pass/fail + edge/failure behavior)
- Validation plan (commands + expected evidence)
- Dependencies (upstream/downstream)
- Artifacts (docs/config/session notes/migration/versioning)

Child ticket readiness rules:

- must satisfy Definition of Ready before starting
- must satisfy Definition of Done with validation evidence before closure

## Done Definition

The project is considered fully `tkv`-native when all conditions hold:

- Public CLI/robot names are ticket-native and consistent.
- Docs/help/examples are `tkv` only (except clearly marked historical notes).
- Env/config namespace is `TKV_*`.
- Go module/import namespace points to `adampush/ticket_viewer`.
- Release/install channels are working and consistent with docs.
- Full CI test/build/vet suite passes on `main`.
