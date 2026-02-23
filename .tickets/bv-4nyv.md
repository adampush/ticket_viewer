---
id: bv-4nyv
status: closed
deps: []
links: []
created: 2026-02-23T01:32:08Z
type: task
priority: 1
assignee: Adam Push
---
# Replace runtime/user-facing legacy beads/bv strings in Go paths


## Notes

**2026-02-23T01:32:32Z**

Scope from audit bv-u39x: runtime/user-visible strings only (non-test Go + embedded viewer assets). Primary hotspots: cmd/bv/main.go (error/help/schema strings), internal/datasource/load.go and source.go (Beads source wording), pkg/updater/updater.go (beads-viewer user-agent and bv binary wording), pkg/ui/tutorial*.go and pkg/ui/model.go package/title/help text, pkg/agents/blurb.go legacy blurb text, pkg/export/viewer_embed.go + pkg/export/viewer_assets/index.html/viewer.js/coi-serviceworker.js title/branding strings. Acceptance: user-facing CLI/help/TUI/export strings prefer tkv/ticket language except explicitly grandfathered compatibility text documented in ticket notes.

**2026-02-23T01:51:45Z**

Implemented additional runtime/user-facing naming cleanup: updated tutorial copy and command examples from bv/br to tkv/tk, refreshed update modal and quit prompt text, updated generated markdown footer branding, modernized AGENTS blurb template to tk workflow, and retitled embedded export viewer UI copy from Beads Viewer to Ticket Viewer. Validation: go build ./... passed; go vet ./... passed; go test ./pkg/ui/... ./pkg/export/... ./pkg/agents/... passed. Note: gofmt -l . still reports pre-existing formatting entries in vendor/* and internal/datasource/watch.go not touched in this ticket.

**2026-02-23T01:53:04Z**

Follow-up cleanup pass completed: updated remaining UI legacy strings in tutorial markdown sections, tree empty-state dependency command (tk dep), and agent prompt modal wording. Updated affected tests in pkg/ui, pkg/export, and pkg/agents to match new runtime copy and tk workflow examples. Re-validated: go test ./pkg/ui/... ./pkg/export/... ./pkg/agents/... passed; go build ./... passed; go vet ./... passed.

**2026-02-23T01:54:35Z**

Final pass: cleaned remaining ui package header comments referencing beads_viewer and re-validated UI suite (go test ./pkg/ui/... passed).

**2026-02-23T02:16:57Z**

Updated embedded COI service worker cache namespace from legacy beads-viewer-coi-v2 to tkv-viewer-coi-v2 in exported viewer assets; export test suite remains green (go test ./pkg/export/...).
