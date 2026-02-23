---
id: bv-u39x
status: open
deps: [bv-4nyv, bv-dztk, bv-0qst, bv-kt62]
links: []
created: 2026-02-23T01:27:54Z
type: task
priority: 2
assignee: Adam Push
---
# Audit lingering legacy references (beads/bv/beads_viewer/bd)


## Notes

**2026-02-23T01:28:02Z**

Audit executed without code changes. Searched for lingering tokens: beads_viewer, beads, bv, bd. Primary commands: rg -n --hidden -S 'beads_viewer|beads-viewer|beads viewer' ... ; rg -n --hidden -S '\bbv --' ... ; rg -n --hidden -S '\bbd\b' ... ; rg -n --hidden -S 'beads_viewer|\bbeads\b|\bbv\b|\bbd\b' cmd pkg internal --glob '!**/*_test.go'.

**2026-02-23T01:28:09Z**

Runtime/user-facing hotspots identified: cmd/bv/main.go (many 'beads'/'beads viewer' strings in error/help/robot schema text; examples around lines 4832, 4944, 5942, 6310, 6534, 7851), internal/datasource/source.go (header comment still says beads_viewer + beads source terminology), internal/datasource/load.go (messages mention Beads sources), pkg/updater/updater.go (User-Agent 'beads-viewer-updater', binary extraction comments still 'bv'), pkg/ui/tutorial_content.go + pkg/ui/tutorial.go + pkg/ui/snapshot.go + pkg/ui/background_worker.go + pkg/ui/model.go (tutorial/package comments still beads_viewer wording), pkg/agents/blurb.go (agent blurb text still says beads_viewer and includes br/bv workflow snippets), pkg/export/viewer_embed.go and pkg/export/viewer_assets/* (embedded UI title still 'Beads Viewer').

**2026-02-23T01:28:15Z**

Docs/scripts hotspots identified: README.md has extensive legacy  command examples and Beads terminology throughout (including legacy install URLs to Dicklesworthstone/beads_viewer around lines ~3541+); AGENTS.md still labels repo as beads_viewer and uses bv framing in many sections; scripts/coverage.sh still references old module path github.com/Dicklesworthstone/beads_viewer in threshold mapping; scripts/test_toon_e2e.sh still branded as 'BV (BEADS VIEWER)'; scripts/capture_baseline.sh emits 'BV Performance Baseline' and uses go run ./cmd/bv examples; docs/*plan and migration artifacts intentionally preserve historical beads/bv references (tkv-native-cleanup-plan, tkv-fork-rename-cutover-plan, tk-dual-mode-extension-plan, bead-history-feature-plan, complementary-features-analysis, performance).

**2026-02-23T01:28:31Z**

Correction/clarification: README legacy command examples are primarily bv --... forms (robot/history/graph/export/search/alerts/workspace/help examples), plus mixed .bv/.beads path references.

**2026-02-23T01:28:38Z**

bd token findings: (a) mostly ticket IDs in comments/tests (e.g., cmd/bv/main.go comments bd-2kxo/bd-2v50, cmd/bv/main_robot_test.go bd-2lmf), (b) docs/performance.md contains user-facing example 'Use bd to break cycles: bd unblock A --from B' (likely stale), (c) tests fixtures include bd-* IDs (tests/testdata/synthetic_complex.jsonl, tests/testdata/real/cass.jsonl) that appear intentional synthetic/historical data.

**2026-02-23T01:28:45Z**

Additional non-doc string hotspots where legacy bv command text is emitted/generated: pkg/analysis/triage.go (RefreshTriage: 'bv --robot-triage'), pkg/export/markdown.go (tables include 'Run bv --robot-triage'), pkg/export/github.go commit message 'Deploy static site via bv --pages', pkg/export/wizard.go preview hint uses 'bv --preview-pages', pkg/drift/config.go help reference 'bv --help'.
