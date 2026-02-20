---
id: bv-pgvp
status: closed
deps: []
links: []
created: 2026-02-20T17:43:05Z
type: task
priority: 2
assignee: Adam Push
parent: bv-0kqn
---
# Seventeenth-pass review tkv cutover plan


## Notes

**2026-02-20T17:43:31Z**

Fresh-eyes pass found remaining config-namespace gap: plan covered .bv/ and BEADS_FILE but did not explicitly govern ~/.config/bv migration policy. Updated docs/tkv-fork-rename-cutover-plan.md with user config directory decision gate, checklist/risk additions, and targeted audit command for .config/bv references.
