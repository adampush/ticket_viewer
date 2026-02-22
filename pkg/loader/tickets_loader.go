package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/adampush/ticket_viewer/pkg/model"
	"gopkg.in/yaml.v3"
)

type ticketFrontmatter struct {
	ID       string   `yaml:"id"`
	Title    string   `yaml:"title"`
	Status   string   `yaml:"status"`
	Priority int      `yaml:"priority"`
	Type     string   `yaml:"type"`
	Assignee string   `yaml:"assignee"`
	Tags     []string `yaml:"tags"`
	Deps     []string `yaml:"deps"`
	Links    []string `yaml:"links"`
	Parent   string   `yaml:"parent"`
	Created  string   `yaml:"created"`
	Updated  string   `yaml:"updated"`
}

// LoadIssuesFromTicketsDir loads tk ticket markdown files from a .tickets directory.
func LoadIssuesFromTicketsDir(ticketsDir string) ([]model.Issue, error) {
	entries, err := os.ReadDir(ticketsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read tickets directory: %w", err)
	}

	paths := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		paths = append(paths, filepath.Join(ticketsDir, e.Name()))
	}
	sort.Strings(paths)

	if len(paths) == 0 {
		return nil, fmt.Errorf("no ticket markdown files found in %s", ticketsDir)
	}

	issues := make([]model.Issue, 0, len(paths))
	invalidCount := 0
	for _, p := range paths {
		issue, parseErr := parseTicketFile(p)
		if parseErr != nil {
			invalidCount++
			continue
		}
		issues = append(issues, issue)
	}

	if len(issues) == 0 {
		return nil, fmt.Errorf("all ticket files were malformed in %s", ticketsDir)
	}

	_ = invalidCount
	return issues, nil
}

func parseTicketFile(path string) (model.Issue, error) {
	var zero model.Issue

	b, err := os.ReadFile(path)
	if err != nil {
		return zero, fmt.Errorf("read %s: %w", path, err)
	}

	fmText, body, ok := splitFrontmatter(string(b))
	if !ok {
		return zero, fmt.Errorf("missing YAML frontmatter")
	}

	var fm ticketFrontmatter
	if err := yaml.Unmarshal([]byte(fmText), &fm); err != nil {
		return zero, fmt.Errorf("invalid frontmatter: %w", err)
	}
	if strings.TrimSpace(fm.ID) == "" {
		return zero, fmt.Errorf("missing id")
	}

	status, err := parseTicketStatus(fm.Status)
	if err != nil {
		return zero, err
	}

	issueType := strings.TrimSpace(fm.Type)
	if issueType == "" {
		issueType = string(model.TypeTask)
	}

	title := strings.TrimSpace(firstHeading(body))
	if title == "" {
		title = strings.TrimSpace(fm.Title)
	}
	if title == "" {
		title = "Untitled"
	}

	createdAt := parseTimeOrZero(fm.Created)
	updatedAt := parseTimeOrZero(fm.Updated)
	if updatedAt.IsZero() {
		if info, statErr := os.Stat(path); statErr == nil {
			updatedAt = info.ModTime()
		}
	}
	if createdAt.IsZero() {
		createdAt = updatedAt
	}

	issue := model.Issue{
		ID:          fm.ID,
		Title:       title,
		Description: strings.TrimSpace(body),
		Status:      status,
		Priority:    normalizePriority(fm.Priority),
		IssueType:   model.IssueType(issueType),
		Assignee:    fm.Assignee,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Labels:      append([]string(nil), fm.Tags...),
	}
	if err := issue.Validate(); err != nil {
		return zero, err
	}

	deps := make([]*model.Dependency, 0, len(fm.Deps)+len(fm.Links)+1)
	for _, dep := range fm.Deps {
		dep = strings.TrimSpace(dep)
		if dep == "" {
			continue
		}
		deps = append(deps, &model.Dependency{IssueID: fm.ID, DependsOnID: dep, Type: model.DepBlocks, CreatedAt: time.Now()})
	}
	for _, dep := range fm.Links {
		dep = strings.TrimSpace(dep)
		if dep == "" {
			continue
		}
		deps = append(deps, &model.Dependency{IssueID: fm.ID, DependsOnID: dep, Type: model.DepRelated, CreatedAt: time.Now()})
	}
	if strings.TrimSpace(fm.Parent) != "" {
		deps = append(deps, &model.Dependency{IssueID: fm.ID, DependsOnID: strings.TrimSpace(fm.Parent), Type: model.DepParentChild, CreatedAt: time.Now()})
	}
	issue.Dependencies = deps

	return issue, nil
}

func splitFrontmatter(content string) (frontmatter string, body string, ok bool) {
	lines := strings.Split(content, "\n")
	if len(lines) < 3 || strings.TrimSpace(lines[0]) != "---" {
		return "", "", false
	}
	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end == -1 {
		return "", "", false
	}
	frontmatter = strings.Join(lines[1:end], "\n")
	body = strings.Join(lines[end+1:], "\n")
	return frontmatter, body, true
}

func parseTicketStatus(v string) (model.Status, error) {
	s := strings.TrimSpace(v)
	if s == "" {
		return model.StatusOpen, nil
	}
	status := model.Status(s)
	if !status.IsValid() {
		return "", fmt.Errorf("invalid status: %s", s)
	}
	return status, nil
}

func firstHeading(body string) string {
	for _, line := range strings.Split(body, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
		}
	}
	return ""
}

func parseTimeOrZero(v string) time.Time {
	v = strings.TrimSpace(v)
	if v == "" {
		return time.Time{}
	}
	if parsed, err := time.Parse(time.RFC3339, v); err == nil {
		return parsed
	}
	if parsed, err := time.Parse("2006-01-02", v); err == nil {
		return parsed
	}
	if unixValue, err := strconv.ParseInt(v, 10, 64); err == nil {
		return time.Unix(unixValue, 0)
	}
	return time.Time{}
}

func normalizePriority(v int) int {
	if v <= 0 {
		return 2
	}
	return v
}
