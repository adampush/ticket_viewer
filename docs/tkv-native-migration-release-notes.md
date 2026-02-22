# tkv-native cleanup migration notes

This note captures the final user-facing migration from legacy `bv`/Beads-era naming to `tkv`/ticket-native naming.

## Breaking changes

- Binary / CLI identity: use `tkv` as the primary command surface.
- Go module path moved to:
  - `github.com/adampush/ticket_viewer`
- Active environment namespace moved to:
  - `TKV_*` (legacy `BV_*` is no longer used by active runtime paths).
- Homebrew install channel is:
  - `brew install adampush/tap/tkv`

## Before / after examples

- Module import path:
  - Before: `github.com/Dicklesworthstone/beads_viewer/pkg/...`
  - After: `github.com/adampush/ticket_viewer/pkg/...`

- Search env defaults:
  - Before: `BV_SEARCH_MODE`, `BV_SEARCH_PRESET`, `BV_SEARCH_WEIGHTS`
  - After: `TKV_SEARCH_MODE`, `TKV_SEARCH_PRESET`, `TKV_SEARCH_WEIGHTS`

- Robot output format env:
  - Before: `BV_OUTPUT_FORMAT=toon`
  - After: `TKV_OUTPUT_FORMAT=toon`

- Browser/test env guards:
  - Before: `BV_NO_BROWSER`, `BV_TEST_MODE`
  - After: `TKV_NO_BROWSER`, `TKV_TEST_MODE`

- Background mode env:
  - Before: `BV_BACKGROUND_MODE`
  - After: `TKV_BACKGROUND_MODE`

- Homebrew install:
  - Before: legacy tap formulas
  - After: `brew tap adampush/tap && brew install adampush/tap/tkv`

## Validation summary

- Engineering gates:
  - `go build ./...`
  - `go vet ./...`
  - `go test ./...`
- Product flows:
  - `go run ./cmd/bv --help`
  - `go run ./cmd/bv --robot-help`
  - `go run ./cmd/bv --robot-schema`
- Contract checks:
  - No legacy `BV_*` in active code/docs (`cmd`, `pkg`, `internal`, `README.md`, `AGENTS.md`; non-test Go paths)
  - No legacy module path `github.com/Dicklesworthstone/beads_viewer` in active Go/packageing sources

## Known constraints

- Historical ticket IDs and legacy directory names (for example `.bv/`) remain where they are intentionally part of persisted project metadata/layout.
- The legacy `bv` executable may still exist in some local environments; `tkv` is the canonical command going forward.
