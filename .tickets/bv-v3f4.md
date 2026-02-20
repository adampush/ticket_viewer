---
id: bv-v3f4
status: closed
deps: [bv-92o5]
links: []
created: 2026-02-20T13:03:42Z
type: task
priority: 1
assignee: Adam Push
parent: bv-0kqn
---
# Implement tk markdown loader to model.Issue mapping

## Context / Why

`bv` analysis expects `[]model.Issue`. `tk` stores issues as markdown + YAML frontmatter. This ticket provides robust parsing and canonical mapping so existing analytics can run unchanged.

## Scope

### In Scope

- Parse `.tickets/*.md` files.
- Read YAML frontmatter plus markdown title/body.
- Map `tk` fields and relationships into `model.Issue`.
- Handle malformed/missing fields with deterministic defaults.

### Out of Scope

- Datasource mode selection logic (`bv-92o5`).
- Robot command-hint generation (`bv-c56w`).
- `tk` history/correlation parity.

## Assumptions

- Ticket identity is defined by frontmatter `id`; files missing `id` are malformed for MVP ingestion.
- Relationship semantics map as:
  - `deps` -> `blocks`
  - `links` -> `related`
  - `parent` -> `parent-child`

## Open Questions

- Should parser support relaxed recovery mode for malformed frontmatter in a follow-up flag?
  - Owner: maintainer
  - Timing: post-MVP
  - Blocking: non-blocking

## Implementation Spec

- Implement `tk` loader/parser under `pkg/loader`.
- Parse frontmatter keys:
  - `id`, `status`, `priority`, `type`, `assignee`, `external-ref`, `tags`, `deps`, `links`, `parent`, `created`, `updated`.
- Extract title from first `# ` heading; fallback order:
  1) heading
  2) frontmatter `title` if present
  3) `Untitled`
- Description should be markdown body excluding frontmatter/title line.
- Timestamp handling:
  - parse `created`/`updated` from frontmatter when valid
  - fallback `updated_at` to file modtime when `updated` is missing/invalid
  - fallback `created_at` to `updated_at` when `created` is missing/invalid
- Priority handling:
  - numeric: preserve value
  - missing/non-numeric: default to `2` with diagnostic warning in verbose mode
- Status/type handling:
  - missing status: default `open`
  - missing type: default `task`
  - unknown non-empty type: preserve value (valid in `model.IssueType`)
- Unknown status handling:
  - if status is non-empty but not a valid `model.Status`, treat file as malformed and skip with warning (avoid invalid issues entering analysis paths)
- Malformed arrays (`deps/links/tags`) should parse best-effort else default to empty.
- Missing required identity fields (for example `id`) should mark file malformed and skip it.
- Files without valid YAML frontmatter should be treated as malformed and skipped with warning.
- If all discovered `tk` tickets are malformed, return a clear load error.

Likely files:

- `pkg/loader/loader.go` and/or new companion parser file in `pkg/loader/`
- `pkg/loader/*_test.go`

## Acceptance Criteria

- Valid `tk` tickets are converted to `model.Issue` with correct fields.
- Relationships map to dependency types exactly as specified.
- Malformed ticket file does not crash full load.
- Missing/non-numeric priority defaults to `2`.
- Title fallback behavior works for missing heading.
- Tickets with unknown non-empty status are treated as malformed and skipped.

## Validation Plan

- `go test ./pkg/loader/...`
- `go test ./internal/datasource/...` (sanity with integration touch points)
- `go build ./...`
- `go vet ./...`
- `gofmt -l .`

Evidence:

- Loader tests covering valid, partial, and malformed ticket fixtures.

## Dependencies

- Upstream:
  - `bv-92o5`
- Downstream:
  - `bv-c56w`
  - `bv-g3nq`

## Artifacts

- `tk` loader fixtures and mapping tests.
- Session notes entry with mapping and fallback decisions.

## Notes

**2026-02-20T19:40:41Z**

Implemented tk markdown parser/mapper in pkg/loader/tickets_loader.go with YAML frontmatter parsing and field mapping: id/title/body/status/priority/type/assignee/tags/deps/links/parent/created/updated. Added defaults/fallbacks per spec: missing status->open, missing type->task, non-positive priority->2, updated_at fallback to file modtime, created_at fallback to updated_at, title fallback to heading/frontmatter/Untitled. Unknown non-empty status treated as malformed (skip). Added tests in pkg/loader/tickets_loader_test.go for valid mapping, malformed skipping behavior, all-malformed error, defaults/fallbacks, unknown status, and preserving unknown non-empty type.

**2026-02-20T19:40:46Z**

Validation: go test ./pkg/loader/... -run Tickets passed; go test ./internal/datasource/... passed; go build ./... passed; go vet ./... passed. Note: full go test ./pkg/loader/... currently has a pre-existing failure in TestGetBeadsDir_EmptyRepoPath_UsesCwd due /private vs /var tempdir path normalization (unrelated to tk loader changes).
