---
id: bv-dztk
status: closed
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

**2026-02-23T02:23:51Z**

Completed docs cleanup pass for README.md and AGENTS.md: normalized project identity to ticket_viewer/tkv in active operational sections, updated remaining AGENTS terminology (ticket-to-commit phrasing and thread example), and added explicit legacy-context notices in README deep-dive/appendix sections so inherited upstream bv/beads terminology is clearly marked non-operational. Also modernized nearby FAQ wording to tkv/tk terminology.
