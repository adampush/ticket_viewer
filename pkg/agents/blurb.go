// Package agents provides AGENTS.md integration for AI coding agents.
// It handles detection, content injection, and preference storage for
// automatically adding ticket workflow instructions to agent configuration files.
package agents

import (
	"regexp"
	"strconv"
	"strings"
)

// BlurbVersion is the current version of the agent instructions blurb.
// Increment this when making breaking changes to the blurb format.
const BlurbVersion = 1

// BlurbStartMarker marks the beginning of injected agent instructions.
const BlurbStartMarker = "<!-- bv-agent-instructions-v1 -->"

// BlurbEndMarker marks the end of injected agent instructions.
const BlurbEndMarker = "<!-- end-bv-agent-instructions -->"

// AgentBlurb contains the instructions to be appended to AGENTS.md files.
const AgentBlurb = `<!-- bv-agent-instructions-v1 -->

---

## Ticket Workflow Integration

This project uses [ticket_viewer](https://github.com/adampush/ticket_viewer) for issue tracking. Ticket files in ` + "`" + `.tickets/` + "`" + ` are tracked in git.

### Essential Commands

` + "```" + `bash
# View issues (launches TUI - avoid in automated sessions)
tkv

# CLI commands for agents (use these instead)
tk ready                               # Show tickets ready to work (no blockers)
tk list --status open                  # All open tickets
tk show <id>                           # Full ticket details
tk create "<title>" -t task -p 2      # Create a new task
tk start <id>                          # Move ticket to in_progress
tk close <id> -r "Completed"          # Close ticket with reason
tk dep <ticket-id> <depends-on-id>     # Add dependency
tk add-note <id> "Implementation ..." # Add implementation notes
` + "```" + `

### Workflow Pattern

1. **Start**: Run ` + "`" + `tk ready` + "`" + ` to find actionable work
2. **Claim**: Use ` + "`" + `tk start <id>` + "`" + `
3. **Work**: Implement the task
4. **Document**: Add notes with ` + "`" + `tk add-note <id> "..."` + "`" + `
5. **Complete**: Use ` + "`" + `tk close <id> -r "Completed"` + "`" + `

### Key Concepts

- **Dependencies**: Tickets can block other tickets. ` + "`" + `tk ready` + "`" + ` shows only unblocked work.
- **Priority**: P0=critical, P1=high, P2=medium, P3=low, P4=backlog (use numbers, not words)
- **Types**: task, bug, feature, epic, question, docs
- **Blocking**: ` + "`" + `tk dep <ticket> <depends-on>` + "`" + ` to add dependencies

### Session Protocol

**Before ending any session, run this checklist:**

` + "```" + `bash
git status              # Check what changed
git add <files>         # Stage code changes
tk add-note <id> "..."  # Record implementation decisions + validation
git commit -m "(tk: <id>) ..."  # Atomic ticket-aligned commit
git push                # Push to remote
` + "```" + `

### Best Practices

- Check ` + "`" + `tk ready` + "`" + ` at session start to find available work
- Update status as you work (` + "`" + `start` + "`" + ` -> ` + "`" + `close` + "`" + `)
- Create new tickets with ` + "`" + `tk create` + "`" + ` when you discover tasks
- Use descriptive titles and set appropriate priority/type
- Keep commits atomic and ticket-scoped

<!-- end-bv-agent-instructions -->`

// SupportedAgentFiles lists the filenames that can contain agent instructions.
var SupportedAgentFiles = []string{
	"AGENTS.md",
	"CLAUDE.md",
	"agents.md",
	"claude.md",
}

// blurbVersionRegex extracts the version number from a blurb marker.
var blurbVersionRegex = regexp.MustCompile(`<!-- bv-agent-instructions-v(\d+) -->`)

// LegacyBlurbPatterns are markers that identify the old blurb format (pre-v1, no HTML markers).
var LegacyBlurbPatterns = []string{
	"### Using bv as an AI sidecar",
	"--robot-insights",
	"--robot-plan",
	"bv already computes the hard parts",
}

// legacyBlurbStartPattern matches the beginning of the legacy blurb.
var legacyBlurbStartPattern = regexp.MustCompile(`(?m)^#{2,3}\s*Using bv as an AI sidecar`)

// legacyBlurbEndPattern matches content near the end of the legacy blurb.
// Uses non-capturing group to make the entire triple-backtick sequence optional.
var legacyBlurbEndPattern = regexp.MustCompile(`(?m)bv already computes the hard parts[^\n]*(?:\n*` + "```" + `)?\n*`)

// legacyBlurbNextSectionPattern matches the start of a new section after the legacy blurb.
// Used as fallback when the end pattern isn't found.
var legacyBlurbNextSectionPattern = regexp.MustCompile(`(?m)^#{1,2}\s+[^#]`)

// ContainsBlurb checks if the content already contains an agent workflow blurb.
func ContainsBlurb(content string) bool {
	return strings.Contains(content, "<!-- bv-agent-instructions-v")
}

// ContainsLegacyBlurb checks if the content contains the old-format blurb (pre-v1, no HTML markers).
// Requires all 4 legacy patterns to match to avoid false positives on content that
// merely references robot flags (like the current AGENTS.md documentation).
func ContainsLegacyBlurb(content string) bool {
	if !legacyBlurbStartPattern.MatchString(content) {
		return false
	}
	matchCount := 0
	for _, pattern := range LegacyBlurbPatterns {
		if strings.Contains(content, pattern) {
			matchCount++
		}
	}
	// Require all patterns - the key differentiator is "bv already computes the hard parts"
	// which only appears in the legacy blurb, not in current documentation
	return matchCount == len(LegacyBlurbPatterns)
}

// ContainsAnyBlurb checks if the content contains either the current or legacy blurb format.
func ContainsAnyBlurb(content string) bool {
	return ContainsBlurb(content) || ContainsLegacyBlurb(content)
}

// GetBlurbVersion extracts the version number from existing blurb content.
func GetBlurbVersion(content string) int {
	matches := blurbVersionRegex.FindStringSubmatch(content)
	if len(matches) < 2 {
		return 0
	}
	version, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}
	return version
}

// NeedsUpdate checks if the content has an older version of the blurb that should be updated.
func NeedsUpdate(content string) bool {
	if ContainsLegacyBlurb(content) {
		return true
	}
	if !ContainsBlurb(content) {
		return false
	}
	return GetBlurbVersion(content) < BlurbVersion
}

// AppendBlurb appends the agent blurb to the given content.
func AppendBlurb(content string) string {
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	content += "\n"
	content += AgentBlurb
	content += "\n"
	return content
}

// RemoveBlurb removes an existing blurb from the content.
func RemoveBlurb(content string) string {
	startIdx := strings.Index(content, "<!-- bv-agent-instructions-v")
	if startIdx == -1 {
		return content
	}
	endIdx := strings.Index(content, BlurbEndMarker)
	if endIdx == -1 {
		return content
	}
	endIdx += len(BlurbEndMarker)
	for endIdx < len(content) && (content[endIdx] == '\n' || content[endIdx] == '\r') {
		endIdx++
	}
	for startIdx > 0 && (content[startIdx-1] == '\n' || content[startIdx-1] == '\r') {
		startIdx--
	}
	return content[:startIdx] + content[endIdx:]
}

// RemoveLegacyBlurb removes the old-format blurb (pre-v1, no HTML markers) from content.
func RemoveLegacyBlurb(content string) string {
	if !ContainsLegacyBlurb(content) {
		return content
	}
	startLoc := legacyBlurbStartPattern.FindStringIndex(content)
	if startLoc == nil {
		return content
	}
	startIdx := startLoc[0]
	endLoc := legacyBlurbEndPattern.FindStringIndex(content[startIdx:])
	var endIdx int
	if endLoc != nil {
		endIdx = startIdx + endLoc[1]
	} else {
		// Fallback: find the next major section heading
		nextLoc := legacyBlurbNextSectionPattern.FindStringIndex(content[startIdx+10:])
		if nextLoc != nil {
			endIdx = startIdx + 10 + nextLoc[0]
		} else {
			endIdx = len(content)
		}
	}
	for endIdx < len(content) && (content[endIdx] == '\n' || content[endIdx] == '\r') {
		endIdx++
	}
	for startIdx > 0 && (content[startIdx-1] == '\n' || content[startIdx-1] == '\r') {
		startIdx--
	}
	if startIdx > 0 {
		startIdx++
	}
	return content[:startIdx] + content[endIdx:]
}

// UpdateBlurb replaces an existing blurb with the current version.
func UpdateBlurb(content string) string {
	content = RemoveLegacyBlurb(content)
	content = RemoveBlurb(content)
	return AppendBlurb(content)
}
