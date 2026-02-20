---
id: bv-ka29
status: closed
deps: []
links: []
created: 2026-02-20T16:49:10Z
type: task
priority: 2
assignee: Adam Push
parent: bv-0kqn
---
# Fifth-pass review tkv cutover plan


## Notes

**2026-02-20T16:49:51Z**

Fresh-eyes pass found additional release-config gaps: .goreleaser.yaml identity fields (project/binary/build path/ldflags/Homebrew/Scoop metadata), version fallback/injection surfaces, and need for release-config validation gate. Updated docs/tkv-fork-rename-cutover-plan.md accordingly.

**2026-02-20T16:54:52Z**

Another fresh-eyes pass found plan inconsistencies: risk numbering order and missing explicit Nix flake decision/risk/check coverage. Updated docs/tkv-fork-rename-cutover-plan.md with decision gate for flake attr policy, checklist/risk additions, nix validation checks, and ordered risk numbering.
