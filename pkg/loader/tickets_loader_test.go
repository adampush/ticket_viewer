package loader

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Dicklesworthstone/beads_viewer/pkg/model"
)

func TestLoadIssuesFromTicketsDir_ValidTicket(t *testing.T) {
	dir := t.TempDir()
	writeTicket(t, dir, "tk-1.md", "---\nid: tk-1\nstatus: open\npriority: 3\ntype: task\ntags: [backend]\ndeps: [tk-0]\nlinks: [tk-2]\nparent: epic-1\ncreated: 2026-02-19T12:00:00Z\nupdated: 2026-02-20T12:00:00Z\n---\n# Build parser\nDetails\n")

	issues, err := LoadIssuesFromTicketsDir(dir)
	if err != nil {
		t.Fatalf("LoadIssuesFromTicketsDir failed: %v", err)
	}
	if len(issues) != 1 {
		t.Fatalf("expected 1 issue, got %d", len(issues))
	}
	issue := issues[0]
	if issue.ID != "tk-1" {
		t.Fatalf("unexpected id: %s", issue.ID)
	}
	if issue.Title != "Build parser" {
		t.Fatalf("unexpected title: %s", issue.Title)
	}
	if issue.Status != model.StatusOpen {
		t.Fatalf("unexpected status: %s", issue.Status)
	}
	if issue.Priority != 3 {
		t.Fatalf("unexpected priority: %d", issue.Priority)
	}
	if issue.IssueType != model.TypeTask {
		t.Fatalf("unexpected issue type: %s", issue.IssueType)
	}
	if len(issue.Dependencies) != 3 {
		t.Fatalf("expected 3 dependencies, got %d", len(issue.Dependencies))
	}
}

func TestLoadIssuesFromTicketsDir_SkipsMalformedAndSucceeds(t *testing.T) {
	dir := t.TempDir()
	writeTicket(t, dir, "good.md", "---\nid: tk-1\nstatus: open\ntype: task\n---\n# Good\n")
	writeTicket(t, dir, "bad.md", "# no frontmatter\n")

	issues, err := LoadIssuesFromTicketsDir(dir)
	if err != nil {
		t.Fatalf("expected success with at least one valid ticket: %v", err)
	}
	if len(issues) != 1 {
		t.Fatalf("expected 1 valid issue, got %d", len(issues))
	}
}

func TestLoadIssuesFromTicketsDir_AllMalformedFails(t *testing.T) {
	dir := t.TempDir()
	writeTicket(t, dir, "bad.md", "# no frontmatter\n")

	_, err := LoadIssuesFromTicketsDir(dir)
	if err == nil {
		t.Fatal("expected error when all ticket files are malformed")
	}
}

func TestLoadIssuesFromTicketsDir_DefaultsAndFallbacks(t *testing.T) {
	dir := t.TempDir()
	path := writeTicket(t, dir, "tk-1.md", "---\nid: tk-1\nstatus: \ntype: \npriority: 0\ncreated: invalid\nupdated: invalid\n---\nBody only\n")
	mod := time.Date(2026, 2, 20, 10, 0, 0, 0, time.UTC)
	if err := os.Chtimes(path, mod, mod); err != nil {
		t.Fatalf("failed to set modtime: %v", err)
	}

	issues, err := LoadIssuesFromTicketsDir(dir)
	if err != nil {
		t.Fatalf("LoadIssuesFromTicketsDir failed: %v", err)
	}
	issue := issues[0]
	if issue.Status != model.StatusOpen {
		t.Fatalf("expected default open status, got %s", issue.Status)
	}
	if issue.IssueType != model.TypeTask {
		t.Fatalf("expected default task type, got %s", issue.IssueType)
	}
	if issue.Priority != 2 {
		t.Fatalf("expected default priority 2, got %d", issue.Priority)
	}
	if issue.UpdatedAt.IsZero() || !issue.UpdatedAt.Equal(mod) {
		t.Fatalf("expected updated_at from file modtime, got %v", issue.UpdatedAt)
	}
	if issue.CreatedAt.IsZero() || !issue.CreatedAt.Equal(issue.UpdatedAt) {
		t.Fatalf("expected created_at fallback to updated_at, got %v / %v", issue.CreatedAt, issue.UpdatedAt)
	}
	if issue.Title != "Untitled" {
		t.Fatalf("expected Untitled title fallback, got %s", issue.Title)
	}
}

func TestLoadIssuesFromTicketsDir_UnknownStatusMalformed(t *testing.T) {
	dir := t.TempDir()
	writeTicket(t, dir, "tk-1.md", "---\nid: tk-1\nstatus: unknown_status\ntype: task\n---\n# bad\n")

	_, err := LoadIssuesFromTicketsDir(dir)
	if err == nil {
		t.Fatal("expected error for unknown non-empty status")
	}
}

func TestLoadIssuesFromTicketsDir_PreservesUnknownType(t *testing.T) {
	dir := t.TempDir()
	writeTicket(t, dir, "tk-1.md", "---\nid: tk-1\nstatus: open\ntype: molecule\n---\n# custom\n")

	issues, err := LoadIssuesFromTicketsDir(dir)
	if err != nil {
		t.Fatalf("LoadIssuesFromTicketsDir failed: %v", err)
	}
	if issues[0].IssueType != model.IssueType("molecule") {
		t.Fatalf("expected custom type preserved, got %s", issues[0].IssueType)
	}
}

func writeTicket(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed writing ticket %s: %v", name, err)
	}
	return path
}
