---
id: bv-0qst
status: closed
deps: []
links: []
created: 2026-02-23T01:32:21Z
type: task
priority: 2
assignee: Adam Push
---
# Update scripts and CI references from beads_viewer/bv to tkv


## Notes

**2026-02-23T01:32:49Z**

Scope from audit bv-u39x: scripts/workflow/tooling references still using legacy module or bv branding. Hotspots include scripts/coverage.sh (old github.com/Dicklesworthstone/beads_viewer package prefixes), scripts/test_toon_e2e.sh branding, scripts/capture_baseline.sh BV labels and cmd examples, and any CI helper references. Acceptance: scripts and CI helper logic align with github.com/adampush/ticket_viewer + tkv naming and execute without regressions.

**2026-02-23T02:26:43Z**

Updated script/CI-adjacent tooling references to tkv naming in scripts/coverage.sh, scripts/capture_baseline.sh, scripts/test_toon_e2e.sh, scripts/test_datasource_e2e.sh, scripts/e2e_hybrid_search.sh, scripts/e2e_web_hybrid_scoring.js, and scripts/benchmark.sh. Replaced legacy module path thresholds with github.com/adampush/ticket_viewer and switched script command invocations from bv to tkv where operationally appropriate. Validation: grep for Dicklesworthstone/beads_viewer in scripts/ and .github/ returns no matches.
