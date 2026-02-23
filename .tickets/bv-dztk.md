---
id: bv-dztk
status: in_progress
deps: []
links: []
created: 2026-02-23T01:32:14Z
type: task
priority: 2
assignee: Adam Push
---
# Clean README and AGENTS legacy bv/beads references


## Notes

**2026-02-23T01:32:41Z**

Scope from audit bv-u39x: top-level docs cleanup in README.md and AGENTS.md. Replace remaining operational 'bv --' command guidance with tkv equivalents where current behavior supports it; update beads/beads_viewer wording where not intentionally historical. Keep explicit fork attribution note and preserve clearly marked historical sections if policy requires. Acceptance: README/AGENTS operational guidance is tkv-first, and any intentional legacy mentions are called out as historical context.

**2026-02-23T02:18:48Z**

Started docs cleanup for tkv-first language. Updated AGENTS.md project identity sections (beads_viewer/bv -> ticket_viewer/tkv), adjusted architecture/examples (including /dp/ticket_viewer and tkv TUI command), and refreshed early README operational wording to ticket-centric language in Core Experience + Architecture snippets. Kept legacy-compatible config path note (~/.config/bv/agent-prompts/) where behavior is still implemented that way.
