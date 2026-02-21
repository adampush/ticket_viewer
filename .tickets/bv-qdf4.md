---
id: bv-qdf4
status: closed
deps: []
links: []
created: 2026-02-20T22:33:50Z
type: bug
priority: 2
assignee: Adam Push
---
# Fix loader cwd path assertion for /private symlink on macOS


## Notes

**2026-02-20T23:19:26Z**

Fixed macOS /private symlink path assertion in pkg/loader TestGetBeadsDir_EmptyRepoPath_UsesCwd by normalizing evaluated paths with /private trimming on darwin before comparison. Validation: go test ./pkg/loader -run TestGetBeadsDir_EmptyRepoPath_UsesCwd passed.
