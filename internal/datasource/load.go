package datasource

import (
	"fmt"
	"os"
	"strings"

	"github.com/adampush/ticket_viewer/pkg/loader"
	"github.com/adampush/ticket_viewer/pkg/model"
)

func testModeLegacyFallbackEnabled() bool {
	return strings.TrimSpace(os.Getenv("TKV_TEST_MODE")) == "1"
}

// LoadMetadata captures source-selection details for callers that need diagnostics.
type LoadMetadata struct {
	SelectedSource DataSource
	Sources        []DataSource
	Reason         string
	MixedSources   bool
}

// LoadIssues performs smart multi-source detection and loading.
// Falls back to legacy JSONL-only loading when smart detection fails, except
// when tk ticket sources were selected but could not be parsed.
func LoadIssues(repoPath string) ([]model.Issue, error) {
	issues, meta, smartErr := LoadIssuesDetailed(repoPath)
	if smartErr == nil {
		return issues, nil
	}
	_ = meta
	return nil, smartErr
}

// LoadIssuesDetailed performs smart loading and returns source metadata.
func LoadIssuesDetailed(repoPath string) ([]model.Issue, *LoadMetadata, error) {
	if strings.TrimSpace(repoPath) == "" {
		if cwd, cwdErr := os.Getwd(); cwdErr == nil {
			repoPath = cwd
		}
	}

	beadsDir, err := loader.GetBeadsDir(repoPath)
	if err != nil {
		beadsDir = ""
	}

	issues, meta, smartErr := loadSmart(beadsDir, repoPath)
	if smartErr == nil {
		return issues, meta, nil
	}

	return nil, meta, smartErr
}

// LoadIssuesFromDir performs smart source detection within a known beads directory.
// This is useful when the caller already knows the .beads path.
func LoadIssuesFromDir(beadsDir string) ([]model.Issue, error) {
	issues, _, smartErr := loadSmart(beadsDir, "")
	if smartErr == nil {
		return issues, nil
	}

	// Fall back to JSONL
	jsonlPath, err := loader.FindJSONLPath(beadsDir)
	if err != nil {
		return nil, err
	}
	return loader.LoadIssuesFromFile(jsonlPath)
}

// loadSmart discovers sources, validates, selects the best, and loads from it.
func loadSmart(beadsDir, repoPath string) ([]model.Issue, *LoadMetadata, error) {
	sources, err := DiscoverSources(DiscoveryOptions{
		BeadsDir:               beadsDir,
		RepoPath:               repoPath,
		ValidateAfterDiscovery: true,
		IncludeInvalid:         false,
	})
	if err != nil {
		return nil, nil, err
	}
	if len(sources) == 0 {
		return nil, nil, fmt.Errorf("no valid sources discovered")
	}

	ticketSources := make([]DataSource, 0, len(sources))
	legacySources := make([]DataSource, 0, len(sources))
	for _, source := range sources {
		if source.Type == SourceTypeTicketsMarkdown {
			ticketSources = append(ticketSources, source)
			continue
		}
		switch source.Type {
		case SourceTypeSQLite, SourceTypeJSONLLocal, SourceTypeJSONLWorktree:
			legacySources = append(legacySources, source)
		}
	}
	if len(ticketSources) == 0 {
		if testModeLegacyFallbackEnabled() && strings.TrimSpace(beadsDir) != "" {
			if jsonlPath, findErr := loader.FindJSONLPath(beadsDir); findErr == nil {
				issues, loadErr := loader.LoadIssuesFromFile(jsonlPath)
				if loadErr == nil {
					selected := DataSource{Type: SourceTypeJSONLLocal, Path: jsonlPath}
					return issues, &LoadMetadata{
						SelectedSource: selected,
						Sources:        append([]DataSource{selected}, sources...),
						Reason:         "test-mode legacy JSONL fallback",
						MixedSources:   hasMixedSources(append([]DataSource{selected}, sources...)),
					}, nil
				}
			}
		}
		if testModeLegacyFallbackEnabled() && len(legacySources) > 0 {
			result, err := SelectBestSourceDetailed(legacySources, DefaultSelectionOptions())
			if err != nil {
				return nil, nil, err
			}

			issues, err := LoadFromSource(result.Selected)
			if err != nil {
				return nil, &LoadMetadata{
					SelectedSource: result.Selected,
					Sources:        result.Candidates,
					Reason:         result.Reason,
					MixedSources:   hasMixedSources(result.Candidates),
				}, err
			}

			return issues, &LoadMetadata{
				SelectedSource: result.Selected,
				Sources:        result.Candidates,
				Reason:         result.Reason,
				MixedSources:   hasMixedSources(result.Candidates),
			}, nil
		}
		return nil, nil, fmt.Errorf("no tk ticket sources discovered (.tickets/*.md)")
	}

	result, err := SelectBestSourceDetailed(ticketSources, DefaultSelectionOptions())
	if err != nil {
		return nil, nil, err
	}

	issues, err := LoadFromSource(result.Selected)
	if err != nil {
		return nil, &LoadMetadata{
			SelectedSource: result.Selected,
			Sources:        result.Candidates,
			Reason:         result.Reason,
			MixedSources:   hasMixedSources(result.Candidates),
		}, err
	}

	meta := &LoadMetadata{
		SelectedSource: result.Selected,
		Sources:        result.Candidates,
		Reason:         result.Reason,
		MixedSources:   hasMixedSources(result.Candidates),
	}

	return issues, meta, nil
}

// LoadFromSource loads issues from a specific DataSource, dispatching to the
// appropriate reader based on source type.
func LoadFromSource(source DataSource) ([]model.Issue, error) {
	switch source.Type {
	case SourceTypeSQLite:
		reader, err := NewSQLiteReader(source)
		if err != nil {
			return nil, fmt.Errorf("failed to open SQLite source %s: %w", source.Path, err)
		}
		defer reader.Close()
		return reader.LoadIssues()

	case SourceTypeJSONLLocal, SourceTypeJSONLWorktree:
		return loader.LoadIssuesFromFile(source.Path)

	case SourceTypeTicketsMarkdown:
		return loader.LoadIssuesFromTicketsDir(source.Path)

	default:
		return nil, fmt.Errorf("unknown source type: %s", source.Type)
	}
}

func hasMixedSources(sources []DataSource) bool {
	hasTickets := false
	hasBeads := false
	for _, s := range sources {
		switch s.Type {
		case SourceTypeTicketsMarkdown:
			hasTickets = true
		case SourceTypeSQLite, SourceTypeJSONLLocal, SourceTypeJSONLWorktree:
			hasBeads = true
		}
	}
	return hasTickets && hasBeads
}

// BuildMixedSourceUsageHint returns a concise note for robot payload usage_hints.
func BuildMixedSourceUsageHint(meta *LoadMetadata) string {
	if meta == nil || !meta.MixedSources {
		return ""
	}
	if meta.SelectedSource.Type == SourceTypeTicketsMarkdown {
		return "Both .tickets and Beads sources were detected; tk tickets were selected by precedence."
	}
	if strings.TrimSpace(meta.SelectedSource.Path) == "" {
		return "Both .tickets and Beads sources were detected; Beads source was selected."
	}
	return fmt.Sprintf("Both .tickets and Beads sources were detected; selected source: %s.", meta.SelectedSource.Path)
}
