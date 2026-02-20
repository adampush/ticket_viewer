# `ticket_viewer` / `tkv` Fork & Rename Cutover Plan

## Goal

Fork the current project into a `tk`-only product, rename it to `ticket_viewer`, and expose CLI as `tkv` while preserving core graph-analysis behavior and robot output quality.

## Decision Summary

- Tracker strategy: `tk` only (drop Beads runtime support)
- Product identity: `ticket_viewer`
- CLI binary: `tkv`
- Migration style: staged cutover (compatibility aliases only where strictly needed)

## Pre-Cutover Decision Gates (must be explicit before implementation)

1. Command package path: keep `cmd/bv` or rename to `cmd/tkv`.
2. Module path policy: rename `go.mod` module in same cutover or follow-up phase.
3. Flag/env policy: strict rename vs temporary aliases (and removal timeline).
4. Correlation/history naming policy for `bead-*` surfaces.
5. Updater/release policy for forked repository owner/name and channels.
6. Nix packaging policy: flake package attr rename (`.#bv` -> `.#tkv`) and any transitional alias window.
7. Tooling policy: required external tools in validation (`goreleaser`, `nix`) and fallback checks when unavailable in developer environments.
8. Workspace/config namespace policy: keep `.bv/*` paths and related flags/envs, or rename to new namespace with migration rules.
9. User config directory policy: keep `~/.config/bv` or migrate to new namespace (and define migration/back-compat behavior).

## Scope

### In Scope

- Rename product-facing identity (`beads_viewer`/`bv` -> `ticket_viewer`/`tkv`)
- Remove Beads-specific runtime paths, docs, and command hints
- Standardize loader/discovery around `.tickets/*.md`
- Preserve existing analysis engines and robot output schemas unless deliberate changes are approved

### Out of Scope

- Rewriting analytics algorithms
- New feature development unrelated to tracker cutover
- Historical Beads data migration tooling (unless separately ticketed)

## Architecture Cutover

1. **CLI and branding layer**
   - Rename command entrypoint behavior from `bv` to `tkv`
   - Update help text, usage examples, docs references

2. **Datasource/loader layer**
   - Make `.tickets/*.md` the only supported source path
   - Remove Beads source discovery/selection branches

3. **Action command generation**
   - Ensure all claim/show/list helper commands emit valid `tk` syntax
   - Remove `br`/Beads command generation paths

4. **Export and robot docs**
   - Ensure exported command snippets and robot docs are `tk`-native

## Rename Surface Checklist

- Module/repo naming references in docs and visible strings
- CLI usage examples (`bv` -> `tkv`)
- Command helper text in robot payloads and exports
- Test fixtures and assertions with hardcoded `br`/`bv` strings
- Developer docs (`AGENTS.md`, README, robot docs)
- Go module path and import graph (`go.mod` module name + internal imports)
- Toolchain floor/version declarations across `go.mod`, CI, and installer scripts (avoid stale minimum-version claims)
- Build/test harness paths that currently compile `../../cmd/bv`
- Updater/release metadata that references current repository/binary naming
- Release manifests/config (`.goreleaser.yaml`) including project name, binary name, build path, ldflags module path, Homebrew/Scoop metadata
- Environment variable policy (`BV_*` names keep vs rename; decide once and document)
- Non-`BV_*` legacy env/config surfaces (for example `BEADS_FILE`, `.bv/workspace.yaml`) and migration/compat policy
- User config directory references (for example `~/.config/bv/...`) across runtime, docs, and recipe/config loaders
- Legacy flag naming policy (for example bead-specific flag names): keep, alias, or rename
- CI/CD workflow references (`.github/workflows/*`) that build/test/package `cmd/bv` or use module-path filters
- Local automation surfaces (`Makefile`, `install.sh`, `install.ps1`, `scripts/*.sh`, test harness scripts) with hardcoded `bv`/`cmd/bv`/repo names
- Installer command-template coupling (`install.ps1` currently uses `$REPO/cmd/$BIN_NAME@latest`) and equivalent path assumptions in shell installer logic
- Nix/Homebrew/release packaging references to current binary/tap names
- Nix flake package/output names and docs/workflows that invoke `nix build .#bv`
- Version injection and fallback surfaces (`pkg/version/version.go` ldflags path comments, fallback strategy)
- User-facing terminology in flags/help text (for example `bead-*` flags) with explicit rename/deprecation policy
- Workspace/config path references in flags/help/docs (`.bv/...`) and any required migration guidance
- Agent instruction generators and templates (`--agents-*`, `pkg/agents/*`, AGENTS docs sync helpers)
- Agent blurb version markers/templates (`bv-agent-instructions-v*`) and their tests/migration behavior
- Tutorial/help content surfaces (`pkg/ui/tutorial*`, long-form help text) with legacy Beads wording
- Correlation/history interface naming (`bead-*` flags/fields) policy: rename now, alias temporarily, or defer with explicit rationale

## Compatibility Strategy

- Prefer a clean rename with no long-term alias debt.
- If operationally necessary, allow short-lived alias period (`bv` shim) behind explicit deprecation note and removal ticket.
- For CLI flags/env vars, avoid half-renames: choose one policy up front (strict rename or temporary aliases) and enforce consistently.
- Decide command package path policy up front: keep `cmd/bv` with renamed binary, or rename to `cmd/tkv`; update all automation/tests consistently based on that decision.

## Risk Checks (Must Pass)

1. **Command validity risk**
   - Verify all emitted commands are valid `tk` commands (no `br` leftovers).

2. **Behavioral regression risk**
   - Ensure triage/plan/insights output remains stable aside from intentional naming/command updates.

3. **Documentation drift risk**
   - Ensure README/help/robot docs match runtime outputs exactly.

4. **Test coverage gap risk**
   - Add/adjust tests for renamed CLI and tk-only command generation paths.

5. **Packaging/invocation risk**
   - Validate local build/install and invocation as `tkv`.

6. **Module/import churn risk**
   - `go.mod` module rename can touch most files; stage it deliberately and run full compile/test gates after each large rename step.

7. **Release/update plumbing risk**
   - Ensure updater/version metadata and release URLs point to the fork identity, not the legacy repo.

8. **Release config mismatch risk**
   - Ensure `.goreleaser.yaml`, install script defaults, and updater endpoints all agree on binary/repo/module identity.

9. **Automation drift risk**
   - Ensure CI workflows, coverage scripts, and helper automation are updated together; avoid green local tests with broken CI paths.

10. **Flag/terminology inconsistency risk**
   - Ensure user-facing flag names/help text are internally consistent for a `tk`-only product (no mixed bead/ticket language unless intentionally aliased).

11. **Hidden legacy wording risk**
   - Ensure generated help/tutorial/agent text has no stale Beads operational guidance after cutover.

12. **Nix packaging drift risk**
   - Ensure `flake.nix` package names, `mainProgram`, subpackage path, and workflow invocations are updated consistently.

13. **Config namespace drift risk**
   - Ensure workspace/config/env naming is coherent (no accidental mix of new product identity with legacy `.bv`/`BEADS_*` semantics unless intentionally aliased).

14. **User config migration risk**
   - Ensure `~/.config/bv` migration (or intentional retention) is explicit, documented, and validated so users do not silently lose settings.

15. **Installer path-coupling risk**
   - Ensure installer build/install commands remain valid for chosen command package path policy (`cmd/bv` vs `cmd/tkv`) and binary rename.

16. **Toolchain version drift risk**
   - Ensure minimum Go version/toolchain requirements are consistent across `go.mod`, CI, `install.sh`, and `install.ps1`.

## Validation Gates

```bash
go build ./...
go vet ./...
gofmt -l .
go test ./...
```

Spot checks:

```bash
go run ./cmd/bv --robot-triage
go run ./cmd/bv --robot-next
go run ./cmd/bv --robot-docs commands
```

Note: after CLI rename is complete, replace these with `tkv` invocation checks. If command package path is renamed, also update `go run ./cmd/...` paths accordingly.

Additional required checks during/after rename:

```bash
go test ./tests/e2e/...
go test ./pkg/ui/...
go test ./pkg/export/...
go test ./pkg/agents/...
go test ./pkg/updater/...
go test ./pkg/workspace/...
go test ./pkg/hooks/...
go test ./pkg/drift/...
```

Verify there are no stale command-path assumptions (for example `cmd/bv`) in test helpers/build scripts.

Automation sanity checks:

```bash
go test ./cmd/...
make -n build
nix flake check
```

And verify workflow/script references are updated for renamed paths/binary names before merge.

If `nix` is unavailable locally, run equivalent CI verification for flake/workflow paths and capture evidence in ticket notes.

Release config sanity checks:

```bash
goreleaser check
```

If `goreleaser` is unavailable locally, run equivalent CI verification and capture evidence in ticket notes.

Validate `.goreleaser.yaml` fields after rename: `project_name`, `builds.main`, `builds.binary`, ldflags module path, Homebrew/Scoop names, homepage/description, and sample commands.

Validate installer scripts after rename (`install.sh`, `install.ps1`): repo URL, binary name, command examples, and post-install usage text.

Validate toolchain messaging/version gates after rename: installer minimum Go version checks/messages must match `go.mod`/CI toolchain policy.

Validate `flake.nix` fields after rename: package attr name, `pname`, `subPackages`, ldflags module path, `meta.mainProgram`, shell hints, and workflow invocations (`nix build .#...`).

Targeted text-audit checks before release:

```bash
rg -n "\bbr\b|bead-history|robot-file-beads|bead" cmd pkg tests docs .github
rg -n "\bbv\b|cmd/bv" cmd pkg
rg -n "cmd/bv|\bbv\b" tests scripts .github Makefile install.sh install.ps1
rg -n "beads_viewer|beads-viewer|beads_rust" cmd pkg tests docs scripts .github
rg -n "\.bv/|BEADS_FILE|\.beads/" cmd pkg tests docs scripts .github install.ps1
rg -n "~/.config/bv|\.config/bv" cmd pkg tests docs scripts .github README.md AGENTS.md
rg -n "beads_viewer|beads-viewer|\bbv\b|cmd/bv|\.bv/|BEADS_FILE" README.md AGENTS.md go.mod flake.nix .goreleaser.yaml install.ps1
```

After policy lock, run focused namespace audits matching the chosen decision:

- If renaming env prefix: audit `BV_` occurrences across `cmd`, `pkg`, `tests`, scripts, and docs.
- If renaming workspace path namespace: audit `.bv/` references in runtime/help/config surfaces and validate migration behavior.

Review and resolve remaining matches based on the chosen alias/deprecation policy.

Text-audit hygiene:

- Keep an explicit allowlist for intentional historical references (for example migration/cutover planning docs and ticket/session artifacts).
- Do not treat planned historical-note matches as release blockers unless they appear in runtime/help/user-facing surfaces.
- Record the allowlist in a tracked artifact (ticket note or dedicated markdown file) with per-match rationale so final review is auditable.

## Rollout Sequence

1. Lock and document pre-cutover decisions (path/module/flags/nix/updater/tooling/config namespace).
2. Land `tk`-only runtime cutover first (no Beads paths).
3. Land CLI/product rename (`bv` -> `tkv`, `beads_viewer` -> `ticket_viewer`).
4. Land module/import/release plumbing updates (`go.mod`, updater metadata, build/test harness paths).
5. Land automation/packaging cutover (`.github/workflows`, Makefile/install/scripts, release taps/channels).
6. Run full validation + docs parity pass.
7. Announce cutover and deprecation/removal timeline for any temporary aliases.

## Suggested Ticket Breakdown

1. pre-cutover decision lock ticket (record policy choices + rationale)
2. `tk`-only datasource/loader cutover
3. command/helper/export `tk` normalization
4. CLI/binary rename to `tkv`
5. automation/packaging path updates (CI/workflows/scripts/install)
6. docs + robot docs rename/parity
7. test updates + compatibility cleanup
