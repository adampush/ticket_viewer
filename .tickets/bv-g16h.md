---
id: bv-g16h
status: closed
deps: []
links: []
created: 2026-02-21T02:11:35Z
type: task
priority: 1
assignee: Adam Push
---
# Harden test-mode fallback guard and fixture YAML quoting


## Notes

**2026-02-21T02:15:48Z**

Implemented hardening follow-up: tightened test-mode legacy fallback to require explicit BV_TEST_MODE=1 in internal/datasource/load.go; added unit test internal/datasource/load_test.go; hardened tests/e2e/common_test.go synthetic ticket writer by quoting YAML scalar frontmatter fields and sanitizing markdown headings. Validation: go build ./..., go vet ./..., go test ./..., plus targeted datasource/e2e tests.
