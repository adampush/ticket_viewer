---
id: bv-c56w
status: closed
deps: [bv-v3f4, bv-92o5]
links: []
created: 2026-02-20T13:03:45Z
type: task
priority: 1
assignee: Adam Push
parent: bv-0kqn
---
# Make robot command hints tracker-aware

## Context / Why

Robot outputs currently suggest Beads commands even when analyzing `tk` data. This creates incorrect action guidance. Command hints must align to active tracker mode.

## Scope

### In Scope

- Update robot command hint generation to switch between `br` and `tk` forms.
- Update `--emit-script` output to tracker-aware claim/show commands.
- Keep JSON schema stable; only command values differ by mode.
- Cover command hints emitted from both CLI assembly and triage payload structures.
- Update exported machine-action guidance that is generated from triage data (for example markdown exports that currently emit `br` commands).

### Out of Scope

- Core ranking/recommendation logic changes.
- Datasource detection (`bv-92o5`) and loader parsing (`bv-v3f4`).
- Full docs updates (handled in `bv-7k1w`).
- Breaking API changes to core analysis entrypoints.

## Assumptions

- Active tracker mode is available where robot payloads are assembled.
- `tk start <id>` is canonical claim action for `tk`.
- `tk` does not support Beads-style `--json` flags on `show/ready/blocked`; generated `tk` commands must use valid `tk` syntax.

## Open Questions

- Should we expose selected tracker mode as an explicit top-level JSON field in robot responses?
  - Owner: maintainer
  - Timing: before release cut
  - Blocking: non-blocking (default: no schema expansion in MVP)

- Where should mixed-source diagnostic note live without schema changes?
  - Owner: maintainer
  - Timing: before merge of this ticket
  - Blocking: non-blocking (default: append concise text item to existing `usage_hints` in robot-triage paths)

## Implementation Spec

- Update command generation in `cmd/bv/main.go` for:
  - `--robot-next` `claim_command` / `show_command`
  - triage grouped recommendations command hints
  - `--emit-script` generated command comments/lines
- Keep current Beads outputs unchanged in Beads mode.
- Ensure triage command helper strings are tracker-aware (either generated with mode or rewritten at response assembly).
- Ensure triage helper commands (`claim_top`, `show_top`, `list_ready`, `list_blocked`) are tracker-valid in both modes.
- Use explicit mapping for helper commands:
  - Beads: `CI=1 br update <id> --status in_progress --json`, `CI=1 br show <id> --json`, `CI=1 br ready --json`, `CI=1 br blocked --json`
  - tk: `tk start <id>`, `tk show <id>`, `tk ready`, `tk blocked`
- For export snippets, avoid emitting unsupported `tk` commands (for example map note/comment actions to `tk add-note` patterns).
- Add mode-aware helper to avoid duplicated conditionals.
- Consume tracker/source metadata from datasource layer rather than re-detecting ad hoc in multiple output paths.
- For export package integration, prefer a backward-compatible options path (for example `GenerateMarkdownWithOptions`) while preserving existing `GenerateMarkdown(...)` callers.
- Update export callsites that should emit tracker-aware snippets to pass tracker options explicitly; leave untouched callsites behavior-stable by default.

Likely files:

- `cmd/bv/main.go`
- `pkg/analysis/triage.go` (if mode-aware command helper generation is done at source)
- `pkg/export/markdown.go` (tracker-aware actionable command snippets)
- `cmd/bv/main.go` (export callsite wiring, if needed)
- `cmd/bv/*_test.go`
- `pkg/export/markdown_test.go`
- `tests/e2e/robot_*` tests where command strings are asserted

## Acceptance Criteria

- In `tk` mode, robot command hints use `tk start` and `tk show`.
- In Beads mode, command hints remain `br update ... --status=in_progress` and `br show`.
- In `tk` mode, export-generated actionable command snippets use `tk` equivalents.
- In `tk` mode, mixed-source diagnostic note is emitted in robot payloads that already include usage hints/diagnostics (no schema expansion).
- In `tk` mode, helper command strings do not include unsupported Beads-only flags (for example `--json`).
- Output JSON structure remains compatible with existing consumers.
- Existing export callsites that do not pass tracker options continue to compile and preserve current behavior.

## Validation Plan

- `go test ./cmd/bv/...`
- `go test ./pkg/export/...`
- `go test ./tests/e2e/... -run Robot`
- `go build ./...`
- `go vet ./...`
- `gofmt -l .`

Evidence:

- Contract test output showing mode-specific command strings.

## Dependencies

- Upstream:
  - `bv-v3f4`
- Downstream:
  - `bv-g3nq`
  - `bv-7k1w`

## Artifacts

- Updated command-hint tests and expected strings.
- Session notes entry documenting schema-compatibility decision.

## Notes

**2026-02-20T19:54:34Z**

Implemented tracker-aware command generation for robot outputs and script/export surfaces. Changes: (1) cmd/bv robot-next and --emit-script now emit tk commands in tk mode via helper functions; (2) robot-triage usage_hints now append mixed-source diagnostic note from datasource metadata when both sources are present; (3) analysis triage command helpers (claim_top/show_top/list_ready/list_blocked) and grouped claim_command fields now honor TriageOptions.TrackerMode; (4) export markdown command snippets are tracker-aware via new MarkdownOptions and SaveMarkdownToFileWithOptions, with tk-safe commands (tk start/show/add-note) and no unsupported br flags in tk mode; (5) added e2e tests for tk robot-next and tk emit-script commands plus export tests for tk markdown options.

**2026-02-20T19:54:42Z**

Validation: go test ./pkg/analysis/... passed; go test ./pkg/export/... passed; go test ./cmd/bv/... passed; go test ./tests/e2e/... -run 'RobotNextContractActionable|EmitScript' passed; go build ./... passed; go vet ./... passed. gofmt -l . reports pre-existing unrelated/vendor files; all touched files were formatted with gofmt -w.
