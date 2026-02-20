---
id: bv-vuhp
status: closed
deps: []
links: []
created: 2026-02-20T17:40:23Z
type: task
priority: 2
assignee: Adam Push
parent: bv-0kqn
---
# Sixteenth-pass review tkv cutover plan


## Notes

**2026-02-20T17:40:45Z**

Fresh-eyes pass found a remaining installer-path coupling gap: install.ps1 uses a command template tied to cmd path + binary name (/cmd/@latest). Updated docs/tkv-fork-rename-cutover-plan.md checklist and risks to explicitly cover installer command-template/path coupling under cmd/bv vs cmd/tkv policy.

**2026-02-20T17:44:51Z**

Follow-up review found config-namespace validation gap: workspace/hooks/drift packages were not explicitly listed in targeted checks despite .bv and user-config migration scope. Updated cutover plan to add go test ./pkg/workspace/... ./pkg/hooks/... ./pkg/drift/... checks.
