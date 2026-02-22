---
id: bv-hqio
status: closed
deps: [bv-ueib]
links: []
created: 2026-02-22T01:48:45Z
type: task
priority: 1
assignee: Adam Push
parent: bv-jzv8
---
# Migrate Go module path and rewrite internal imports to ticket_viewer

## Context / Why

This is Workstream 5 from `docs/tkv-native-cleanup-plan.md`. The module/import namespace must be fully migrated to `github.com/adampush/ticket_viewer` for a true tkv-native identity.

## Scope

### In Scope

- Update `go.mod` module path to `github.com/adampush/ticket_viewer`.
- Rewrite internal Go imports to new module path.
- Ensure build/test/vet passes after migration.

### Out of Scope

- Packaging/release channel updates (Workstream 6).
- Unrelated behavior refactors.

## Ticket Type & Granularity

- Type: Task
- Granularity target: 1-3 commits focused on module/import namespace migration.

## Prerequisites & Dev State

- Upstream env/config migration `bv-ueib` complete.
- Workstream 5 contract defaults locked in plan.

## Assumptions

- Partial migration is invalid; this ticket must land as a complete cutover in one PR.

## Open Questions

- None currently. Any integration-specific fallout must be documented as follow-up tickets.

## Implementation Guidance

- Update module declaration in `go.mod`.
- Rewrite Go imports across `cmd`, `pkg`, `internal`, `tests`, and helper Go scripts.
- Keep runtime behavior unchanged; this is a namespace cutover.
- Inherit NFR constraints: no regressions in determinism, timeout behavior, or test-mode browser guards.

## Acceptance Criteria

- `go.mod` points to `github.com/adampush/ticket_viewer`.
- No Go imports reference `github.com/Dicklesworthstone/beads_viewer`.
- `go build ./...`, `go vet ./...`, and `go test ./...` pass.

## Deployment & State

- Breaking/Non-breaking: breaking for downstream integrations pinned to old module path.
- Migration requirements: release notes include module path cutover note.
- Safe to deploy independently: yes, after Workstream 4.

## Validation Plan

- grep check for `github.com/Dicklesworthstone/beads_viewer` in Go sources
- `go build ./...`
- `go vet ./...`
- `go test ./...`

Expected evidence:

- grep output proving no legacy import references
- full-suite outputs attached in PR checklist
- `tk add-note bv-hqio "validation: ..."` with results

## Dependencies

- Upstream: `bv-ueib`
- Downstream: `bv-4nc1`

## Artifacts

- Module/import migration evidence
- Release note draft snippet for module path change

## Notes

**2026-02-22T03:15:55Z**

Migrated module path to github.com/adampush/ticket_viewer and rewrote internal imports/references across cmd/pkg/internal/tests/scripts Go sources. Validation: rg -n 'github.com/Dicklesworthstone/beads_viewer' --glob '*.go' (no matches); gofmt -l cmd/bv/burndown_test.go
cmd/bv/main.go
cmd/bv/main_test.go
cmd/bv/profile_test.go
cmd/bv/search_output.go
internal/datasource/diff.go
internal/datasource/load.go
internal/datasource/load_test.go
internal/datasource/sqlite.go
pkg/agents/blurb.go
pkg/agents/tty_guard.go
pkg/analysis/advanced_insights.go
pkg/analysis/advanced_insights_test.go
pkg/analysis/bench_generators_test.go
pkg/analysis/bench_pathological_test.go
pkg/analysis/bench_realdata_test.go
pkg/analysis/bench_test.go
pkg/analysis/benchmark_test.go
pkg/analysis/betweenness_approx_test.go
pkg/analysis/cache.go
pkg/analysis/cache_extra_test.go
pkg/analysis/cache_test.go
pkg/analysis/config.go
pkg/analysis/cycle_warnings.go
pkg/analysis/cycle_warnings_test.go
pkg/analysis/dependency_suggest.go
pkg/analysis/dependency_suggest_test.go
pkg/analysis/diff.go
pkg/analysis/diff_extended_test.go
pkg/analysis/diff_test.go
pkg/analysis/duplicates.go
pkg/analysis/duplicates_test.go
pkg/analysis/e2e_startup_test.go
pkg/analysis/eta.go
pkg/analysis/eta_test.go
pkg/analysis/golden_test.go
pkg/analysis/graph.go
pkg/analysis/graph_cycles_test.go
pkg/analysis/graph_extra_test.go
pkg/analysis/graph_test.go
pkg/analysis/insights_signals_test.go
pkg/analysis/insights_test.go
pkg/analysis/invariance_test.go
pkg/analysis/label_health.go
pkg/analysis/label_health_test.go
pkg/analysis/label_suggest.go
pkg/analysis/label_suggest_test.go
pkg/analysis/perf_invariants_test.go
pkg/analysis/plan.go
pkg/analysis/plan_extended_test.go
pkg/analysis/plan_test.go
pkg/analysis/priority.go
pkg/analysis/priority_test.go
pkg/analysis/real_data_test.go
pkg/analysis/risk.go
pkg/analysis/risk_test.go
pkg/analysis/sample_integration_test.go
pkg/analysis/status_fullstats_test.go
pkg/analysis/suggest_all.go
pkg/analysis/suggest_all_test.go
pkg/analysis/suggestions.go
pkg/analysis/triage.go
pkg/analysis/triage_context.go
pkg/analysis/triage_context_test.go
pkg/analysis/triage_test.go
pkg/analysis/whatif_test.go
pkg/cass/correlation.go
pkg/cass/correlation_test.go
pkg/correlation/causality.go
pkg/correlation/cocommit.go
pkg/correlation/cocommit_test.go
pkg/correlation/correlator.go
pkg/correlation/feedback.go
pkg/correlation/file_index.go
pkg/correlation/network.go
pkg/correlation/network_test.go
pkg/correlation/orphan.go
pkg/correlation/related.go
pkg/correlation/reverse.go
pkg/correlation/scorer.go
pkg/correlation/temporal.go
pkg/correlation/temporal_test.go
pkg/correlation/types.go
pkg/debug/debug.go
pkg/drift/drift.go
pkg/drift/drift_test.go
pkg/export/cloudflare.go
pkg/export/external_tools_test.go
pkg/export/gh_pages_e2e_test.go
pkg/export/github.go
pkg/export/github_test.go
pkg/export/graph_export.go
pkg/export/graph_export_test.go
pkg/export/graph_interactive.go
pkg/export/graph_interactive_test.go
pkg/export/graph_render_beautiful.go
pkg/export/graph_render_golden_test.go
pkg/export/graph_snapshot.go
pkg/export/graph_snapshot_bench_test.go
pkg/export/graph_snapshot_svg_test.go
pkg/export/graph_snapshot_test.go
pkg/export/integration_test.go
pkg/export/main_test.go
pkg/export/markdown.go
pkg/export/markdown_test.go
pkg/export/mermaid_generator.go
pkg/export/sqlite_export.go
pkg/export/sqlite_export_metrics_test.go
pkg/export/sqlite_export_test.go
pkg/export/sqlite_types.go
pkg/export/wizard.go
pkg/hooks/config.go
pkg/hooks/executor_test.go
pkg/loader/benchmark_test.go
pkg/loader/bom_test.go
pkg/loader/fuzz_test.go
pkg/loader/git.go
pkg/loader/git_test.go
pkg/loader/loader.go
pkg/loader/loader_extra_test.go
pkg/loader/loader_test.go
pkg/loader/pool.go
pkg/loader/real_data_test.go
pkg/loader/robustness_test.go
pkg/loader/sprint.go
pkg/loader/sprint_test.go
pkg/loader/synthetic_test.go
pkg/loader/tickets_loader.go
pkg/loader/tickets_loader_test.go
pkg/metrics/timing.go
pkg/recipe/loader_test.go
pkg/recipe/types_test.go
pkg/search/config.go
pkg/search/config_test.go
pkg/search/documents.go
pkg/search/documents_test.go
pkg/search/embedder.go
pkg/search/embedder_test.go
pkg/search/hybrid_scorer_real_test.go
pkg/search/metrics_cache_impl.go
pkg/search/metrics_cache_test.go
pkg/search/search_pipeline_real_test.go
pkg/search/short_query_hybrid_test.go
pkg/search/vector_index.go
pkg/testutil/assertions.go
pkg/testutil/generator.go
pkg/testutil/generator_test.go
pkg/ui/actionable.go
pkg/ui/actionable_test.go
pkg/ui/agent_prompt_modal.go
pkg/ui/attention.go
pkg/ui/attention_test.go
pkg/ui/background_worker.go
pkg/ui/background_worker_test.go
pkg/ui/benchmark_test.go
pkg/ui/board.go
pkg/ui/board_test.go
pkg/ui/cass_session_modal.go
pkg/ui/cass_session_modal_test.go
pkg/ui/coverage_extra_test.go
pkg/ui/delegate.go
pkg/ui/delegate_test.go
pkg/ui/flow_matrix.go
pkg/ui/flow_matrix_test.go
pkg/ui/graph.go
pkg/ui/graph_bench_test.go
pkg/ui/graph_golden_test.go
pkg/ui/graph_test.go
pkg/ui/helpers.go
pkg/ui/helpers_test.go
pkg/ui/history.go
pkg/ui/history_test.go
pkg/ui/insights.go
pkg/ui/insights_test.go
pkg/ui/integration_test.go
pkg/ui/item.go
pkg/ui/item_test.go
pkg/ui/label_dashboard.go
pkg/ui/label_dashboard_test.go
pkg/ui/logic_test.go
pkg/ui/main_test.go
pkg/ui/model.go
pkg/ui/model_test.go
pkg/ui/recipe_picker.go
pkg/ui/recipe_picker_test.go
pkg/ui/semantic_search.go
pkg/ui/semantic_search_test.go
pkg/ui/shortcuts_sidebar.go
pkg/ui/snapshot.go
pkg/ui/snapshot_test.go
pkg/ui/sprint_view.go
pkg/ui/sprint_view_keys_test.go
pkg/ui/tree.go
pkg/ui/tree_bench_test.go
pkg/ui/tree_test.go
pkg/ui/triage_preservation_test.go
pkg/ui/tutorial_content.go
pkg/ui/update_keys_test.go
pkg/ui/update_modal.go
pkg/ui/update_test.go
pkg/ui/velocity_comparison.go
pkg/ui/velocity_comparison_test.go
pkg/ui/visuals_test.go
pkg/ui/workspace_filter_test.go
pkg/updater/updater.go
pkg/updater/updater_test.go
pkg/version/version.go
pkg/watcher/watcher.go
pkg/watcher/watcher_test.go
pkg/workspace/loader.go
pkg/workspace/loader_test.go
pkg/workspace/types_test.go
scripts/generate_testdata.go
tests/e2e/agents_integration_e2e_test.go
tests/e2e/board_e2e_test.go
tests/e2e/board_swimlane_e2e_test.go
tests/e2e/cass_modal_e2e_test.go
tests/e2e/common_test.go
tests/e2e/correlation_e2e_test.go
tests/e2e/export_pages_test.go
tests/e2e/history_timeline_e2e_test.go
tests/e2e/robot_history_test.go
tests/e2e/robot_matrix_test.go
tests/e2e/robot_search_hybrid_test.go
tests/e2e/robot_search_test.go
tests/e2e/robot_suggest_test.go
tests/e2e/search_benchmark_test.go
tests/e2e/tui_hybrid_search_test.go
tests/e2e/tui_snapshot_test.go
tests/e2e/update_flow_test.go
tests/e2e/wizard_flow_e2e_test.go (clean); go build ./...; go vet ./...; go test ./... (pass).
