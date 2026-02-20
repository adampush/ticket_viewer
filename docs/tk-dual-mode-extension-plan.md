# Dual-Mode Tracker Extension Plan (`beads` + `tk`)

## Planning Quality Bar (used for review before implementation)

Any implementation plan in this project should include all items below so we avoid hidden assumptions and vague execution.

1. **Explicit assumptions register**: every assumption written down, with confidence and validation approach.
2. **Clear user outcomes**: who uses it, what they do, and what success looks like in real workflows.
3. **Architecture and file-level design**: exact packages/files to touch, plus why.
4. **Blast-radius analysis**: what existing behavior may change and how we constrain impact.
5. **Edge-case matrix**: malformed input, mixed modes, empty data, and conflict cases.
6. **Test strategy with commands**: unit/integration/e2e scope, fixtures, and exact validation commands.
7. **Definition of done**: measurable pass/fail criteria.
8. **Risk + rollback plan**: what can fail in production and how to recover safely.
9. **Non-goals and deferred work**: what we are intentionally not solving in this phase.

## Assumptions Register

1. **Assumption:** `tk` ticket files use YAML frontmatter consistently enough for tolerant parsing.
   - **Confidence:** medium
   - **Validation:** parser tests with malformed and partial frontmatter fixtures.

2. **Assumption:** `model.Issue` can represent `tk` relationships without schema changes.
   - **Confidence:** high
   - **Validation:** map `deps/links/parent` into existing dependency types and run analysis tests.

3. **Assumption:** command hints are the only user-visible places tightly coupled to `br` in robot outputs.
   - **Confidence:** medium
   - **Validation:** grep scan + robot output contract tests for all `--robot-*` commands.

4. **Assumption:** preferring `.tickets/` when both trackers exist is acceptable for MVP.
   - **Confidence:** medium
   - **Validation:** document behavior + add deterministic selection tests.

5. **Assumption:** full history/correlation parity is not required for MVP.
   - **Confidence:** high
   - **Validation:** acceptance criteria scoped to triage/plan/insights/next outputs.

## Objective

Extend `bv` so it can analyze projects using either:

- Beads data (`.beads/` JSONL/SQLite), or
- `tk` data (`.tickets/*.md` with YAML frontmatter)

without forking the product and without regressing existing Beads behavior.

## Product Decision

- **Mode:** dual-mode auto-detect
- **Detection rule (initial):**
  1. If `.tickets/` exists and has ticket files, run in `tk` mode
  2. Else use existing Beads source detection
  3. If both exist, prefer `tk` (can later add override flag/env)

## User Flows and Desired Outcomes

1. **`tk`-only repo user**
   - Runs `bv --robot-triage` from project root.
   - Receives valid recommendations with `tk` claim/show commands.

2. **Beads-only repo user**
   - Existing behavior remains unchanged.
   - Existing automation/scripts continue to work.

3. **Repo containing both `.tickets` and `.beads`**
   - Deterministic mode selection (`tk` first for MVP).
   - Output makes selected source/mode explicit to avoid confusion.

4. **Empty or partially initialized repo**
   - Clear, actionable error message explaining missing ticket sources.

## Scope

### In Scope (MVP)

1. Add `tk` datasource support with markdown parsing.
2. Map `tk` ticket fields into existing `model.Issue`.
3. Keep all current analysis/triage/robot logic unchanged.
4. Make robot command hints tracker-aware (`br` vs `tk`).
5. Add tests for auto-detect + parser + robot command output.
6. Update user and operator documentation for dual-mode behavior.

### Out of Scope (MVP)

1. Rewriting analytics/scoring models.
2. New UI paradigm for `tk`.
3. Full `tk`-native history/correlation parity (can be follow-up).

## Documentation Deliverables

Documentation updates are part of done criteria, not optional polish.

1. **Primary user docs**
   - Update root `README.md` to explain dual-mode tracker support.
   - Add `tk` examples for key robot flows (`--robot-triage`, `--robot-next`, `--robot-plan`).
   - Document mixed-repo precedence behavior (`.tickets` preferred over `.beads` in MVP).

2. **CLI help / robot docs**
   - Update help text and machine-readable docs sections that currently imply Beads-only command hints.
   - Ensure command examples are tracker-aware or clearly mode-annotated.

3. **Developer docs**
   - Add a short architecture note describing where tracker-specific logic lives (datasource/loader boundaries).
   - Document parser assumptions and failure behavior for malformed `tk` tickets.

4. **Migration/compatibility note**
   - Include guidance for teams with both `.tickets` and `.beads` present.
   - State current limits (for example history/correlation parity if still deferred).

## `tk` Data Model Mapping

Source format: markdown files in `.tickets/` with YAML frontmatter.

### Field mapping

- `id` -> `Issue.ID`
- first markdown `# Heading` -> `Issue.Title`
- body markdown (minus frontmatter/title) -> `Issue.Description`
- `status` -> `Issue.Status` (`open`, `in_progress`, `closed`)
- `priority` -> `Issue.Priority`
- `type` -> `Issue.IssueType`
- `assignee` -> `Issue.Assignee`
- `external-ref` -> `Issue.ExternalRef`
- `tags` -> `Issue.Labels`
- `created` -> `Issue.CreatedAt` (best-effort parse)
- `updated` (if present) -> `Issue.UpdatedAt`; otherwise fallback to file modtime or `CreatedAt`

### Relationship mapping

- `deps: [A,B]` -> `Dependency{Type: "blocks", DependsOnID: A|B}`
- `links: [A,B]` -> `Dependency{Type: "related", DependsOnID: A|B}`
- `parent: P` -> `Dependency{Type: "parent-child", DependsOnID: P}`

## Architecture Changes

### 1) Datasource layer

Update `internal/datasource` to include `tk` source type and loader path.

Planned changes:

- Add source type constant for `tk` markdown directory/files.
- Add discovery branch for `.tickets/*.md`.
- Add validation branch for `tk` sources (readability + minimal required fields).
- Wire `LoadFromSource` to dispatch to `tk` loader.

Files expected to change (MVP):

- `internal/datasource/source.go`
- `internal/datasource/load.go`
- `internal/datasource/validate.go`
- `internal/datasource/select.go` (source-type reason strings)

### 2) Loader layer

Add `tk` parser in `pkg/loader`.

Planned functions:

- Discover `.tickets` directory from repo root.
- Enumerate `*.md` ticket files.
- Parse YAML frontmatter and markdown sections.
- Build `[]model.Issue` with robust defaults.

Files expected to change (MVP):

- `pkg/loader/loader.go` (or companion file under `pkg/loader/`)
- new parser tests under `pkg/loader/*_test.go`

### 3) Robot command emission

Update command hints currently hardcoded to Beads CLI.

Examples:

- Beads mode: `br update <id> --status=in_progress`, `br show <id>`
- `tk` mode: `tk start <id>`, `tk show <id>`

Affected paths include `--robot-next`, `--robot-triage`, and `--emit-script` outputs.

Files expected to change (MVP):

- `cmd/bv/main.go`
- robot contract tests in `tests/e2e/` and/or `cmd/bv/*_test.go`

## Blast Radius Control

1. Keep all analysis packages (`pkg/analysis`, scoring logic) unchanged.
2. Contain tracker-specific logic at datasource/loader boundaries.
3. Keep output schema stable; only command-hint values vary by tracker mode.
4. Avoid broad refactors during MVP; add targeted abstractions only where coupling exists.

## Edge Cases and Handling

1. `.tickets/` exists but has zero `.md` files -> fall back to Beads detection.
2. `.tickets/` file missing `id` -> skip with warning in verbose mode; do not crash full load.
3. Missing markdown `# title` -> use frontmatter `title` if available, else `Untitled`.
4. Malformed `deps/links/tags` arrays -> parse best effort; default to empty array.
5. Unknown `status/type/priority` values -> preserve raw value where possible, apply safe defaults for scoring.
6. Cyclic deps in `tk` data -> allow; existing graph analysis handles cycles.
7. Both trackers present -> deterministic `tk` precedence for MVP.

## Implementation Steps

1. Add tracker-mode detection utility (`beads` vs `tk`).
2. Implement `tk` ticket file parser and `Issue` mapper.
3. Integrate `tk` source into datasource discovery/selection/loading.
4. Make robot command generation depend on detected tracker mode.
5. Add/update tests.
6. Run validation gates (`go build`, `go vet`, `gofmt -l`, targeted tests).

## Test Plan

### Unit tests

1. `tk` frontmatter parsing (valid/missing fields/malformed arrays).
2. Relationship mapping (`deps`, `links`, `parent`).
3. Timestamp fallback behavior.
4. Source detection precedence (`.tickets` vs `.beads`).
5. Fallback title/body extraction behavior.
6. Invalid ticket file handling (single bad file does not fail entire load unless strict mode is introduced).

### Integration tests

1. `--robot-next` on `tk` fixture emits `tk start` + `tk show` commands.
2. `--robot-triage` works on `tk` fixture and returns recommendations.
3. Existing Beads robot tests continue to pass with unchanged outputs.
4. Mixed-repo fixture (`.tickets` + `.beads`) confirms deterministic mode selection.

### Non-regression checks

- Build: `go build ./...`
- Vet: `go vet ./...`
- Format check: `gofmt -l .`
- Optional static check: `ubs <changed-go-files>`
- Targeted tests:
  - `go test ./internal/datasource/...`
  - `go test ./pkg/loader/...`
  - `go test ./cmd/bv/...`

## Validation Matrix

1. **Functional correctness**: robot outputs valid JSON and meaningful recommendations in both modes.
2. **Compatibility**: existing Beads snapshots/contracts unchanged (except where explicitly mode-conditional).
3. **Usability**: command hints are executable and tracker-correct.
4. **Stability**: no panics on malformed or partial `.tickets` data.
5. **Documentation parity**: README/help/robot docs accurately match shipped behavior.

## Documentation Verification Checklist

Use this checklist before calling the feature complete.

1. **README accuracy**
   - Confirm dual-mode behavior is documented.
   - Confirm mixed-repo precedence is documented.
   - Confirm `tk` examples run as written.

2. **CLI help accuracy**
   - Run `bv --help` and confirm tracker wording is not Beads-only where dual-mode now applies.
   - Confirm command examples do not imply wrong claim/show commands in `tk` mode.

3. **Robot docs accuracy**
   - Run `bv --robot-docs guide` and `bv --robot-docs commands`.
   - Confirm examples and command-hint descriptions match dual-mode behavior.

4. **Runtime spot checks (both modes)**
   - In a `tk` fixture repo, run `bv --robot-next` and verify `claim_command` / `show_command` use `tk`.
   - In a Beads fixture repo, run `bv --robot-next` and verify `claim_command` / `show_command` still use `br`.

5. **Suggested verification commands**
   - `bv --help`
   - `bv --robot-docs guide`
   - `bv --robot-docs commands`
   - `bv --robot-next`
   - `bv --robot-triage`

## Risks and Mitigations

1. **Risk:** Ambiguous mode when both `.tickets` and `.beads` exist.
   - **Mitigation:** deterministic precedence + document behavior; add override in follow-up.

2. **Risk:** `tk` markdown variability breaks parser.
   - **Mitigation:** tolerant parser + fallback defaults + focused tests.

3. **Risk:** Hidden Beads assumptions in downstream features (history/correlation).
   - **Mitigation:** keep MVP surface focused; explicitly gate unsupported behavior where needed.

4. **Risk:** Excessive invasive edits increase regression probability.
   - **Mitigation:** isolate changes to loader/datasource/command-hints and avoid broad package-level refactors.

## Rollout and Recovery

1. Land behind deterministic auto-detect behavior with no flags required.
2. If critical regression is found, temporarily disable `tk` discovery path and preserve Beads-only behavior.
3. Follow up with explicit `--tracker` override for operational control.

## Acceptance Criteria (MVP)

1. In a repo with only `.tickets/`, robot commands (`--robot-next`, `--robot-triage`, `--robot-plan`, `--robot-insights`) return valid output.
2. Command hints in robot output use `tk` commands in `tk` mode.
3. In a Beads repo, behavior remains unchanged.
4. New tests pass and no regressions in touched packages.
5. Edge-case fixtures (missing fields, malformed arrays, mixed trackers) produce deterministic and non-crashing behavior.
6. Documentation updates are merged and validated against behavior (README + CLI help + robot docs references).

## Resolved Defaults (spec lock for implementation)

1. **Mixed tracker warning behavior**
   - **Default:** Yes, emit a concise warning/diagnostic note when both `.tickets` and `.beads` are present and `tk` is selected.
   - **Rationale:** prevents silent mode surprises and makes auto-detect transparent.

2. **Malformed `tk` ticket handling**
   - **Default:** skip malformed ticket files with warning in verbose/diagnostic paths; do not hard-fail entire load unless all tickets fail.
   - **Rationale:** resilient operation in real repos with partial corruption.

3. **Out-of-range `tk` priority handling**
   - **Default:** preserve parsed numeric value if possible; if non-numeric/missing, use safe default `2` and warn in verbose mode.
   - **Rationale:** avoid destructive normalization while keeping scoring stable.

4. **Tracker override control**
   - **Default:** keep auto-detect only in MVP; add explicit `--tracker` override as immediate follow-up item.
   - **Rationale:** keeps MVP blast radius smaller and delivery faster.

5. **Mode precedence lock**
   - **Default:** `.tickets` with at least one ticket file wins over Beads sources.
   - **Rationale:** aligns with the selected dual-mode strategy and prevents ambiguous selection.

## Follow-Up Backlog (Post-MVP)

1. Add explicit tracker override flag/env (e.g., `--tracker=tk|beads`).
2. Extend watch/reload UX to deeply optimize `.tickets` updates.
3. Add `tk`-aware history/correlation parity.
4. Update README and robot docs with first-class `tk` examples.
