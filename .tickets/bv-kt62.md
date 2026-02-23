---
id: bv-kt62
status: closed
deps: []
links: []
created: 2026-02-23T01:32:25Z
type: task
priority: 3
assignee: Adam Push
---
# Define policy for historical docs and test fixture legacy terms


## Notes

**2026-02-23T01:32:54Z**

Scope from audit bv-u39x: decide and document what legacy terms remain intentionally in historical docs and fixtures (e.g., docs/*plan archives, tests/testdata synthetic IDs like bd-* and beads.jsonl fixtures). Produce allowlist policy and boundaries so future cleanups avoid churning archival material while still preventing regressions in active user-facing surfaces.

**2026-02-23T02:27:20Z**

Defined explicit legacy terminology policy in AGENTS.md. Added allowlist/disallowlist boundaries for where beads/bv/br terms are acceptable (historical docs, fixtures, compatibility markers, retained path names) and where they are prohibited (active operational docs, runtime/user-facing strings, new workflow examples). Included requirement to mark intentional legacy usage with historical/compatibility notes when touching those files.
