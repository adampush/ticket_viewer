---
id: bv-dpkp
status: closed
deps: []
links: []
created: 2026-02-20T16:45:56Z
type: task
priority: 2
assignee: Adam Push
parent: bv-0kqn
---
# Third-pass review tkv cutover plan


## Notes

**2026-02-20T16:46:26Z**

Fresh-eyes review found additional cutover gaps: legacy Beads wording in tutorial/help/agents surfaces and bead-named correlation flags/fields policy not explicitly captured. Updated docs/tkv-fork-rename-cutover-plan.md checklist/risks/validation with targeted text-audit checks and explicit naming-policy decisions.

**2026-02-20T16:48:05Z**

Another fresh-eyes pass found remaining ambiguity in unresolved policy choices. Updated cutover plan with explicit pre-cutover decision gates, added agent blurb marker migration surface, included pkg/agents tests in validation, and expanded text-audit checks for legacy beads_viewer naming.
