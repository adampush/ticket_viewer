package main_test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEmitScript_BashAndFish(t *testing.T) {
	bv := buildBvBinary(t)
	env := t.TempDir()

	// Ensure at least one actionable recommendation.
	writeTickets(t, env, map[string]string{
		"a.md": "---\nid: A\nstatus: open\npriority: 1\ntype: task\n---\n# Unblocker\n",
		"b.md": "---\nid: B\nstatus: open\npriority: 2\ntype: task\ndeps: [A]\n---\n# Blocked\n",
	})

	tests := []struct {
		name        string
		formatFlag  string
		wantShebang string
		wantExtra   string
	}{
		{name: "bash", formatFlag: "bash", wantShebang: "#!/usr/bin/env bash", wantExtra: "set -euo pipefail"},
		{name: "fish", formatFlag: "fish", wantShebang: "#!/usr/bin/env fish", wantExtra: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(bv, "--emit-script", "--script-limit=1", "--script-format="+tt.formatFlag)
			cmd.Dir = env
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("run failed: %v\n%s", err, out)
			}
			s := string(out)
			if !strings.Contains(s, tt.wantShebang) {
				t.Fatalf("missing shebang %q:\n%s", tt.wantShebang, s)
			}
			if tt.wantExtra != "" && !strings.Contains(s, tt.wantExtra) {
				t.Fatalf("missing %q:\n%s", tt.wantExtra, s)
			}
			if !strings.Contains(s, "tk show A") {
				t.Fatalf("missing tk show command for A:\n%s", s)
			}
			if !strings.Contains(s, "# Data hash:") {
				t.Fatalf("missing data hash header:\n%s", s)
			}
		})
	}
}

func TestEmitScript_TKCommands(t *testing.T) {
	bv := buildBvBinary(t)
	env := t.TempDir()
	writeTickets(t, env, map[string]string{
		"tk-a.md": "---\nid: A\nstatus: open\npriority: 1\ntype: task\n---\n# Unblocker\n",
		"tk-b.md": "---\nid: B\nstatus: open\npriority: 2\ntype: task\ndeps: [A]\n---\n# Blocked\n",
	})

	cmd := exec.Command(bv, "--emit-script", "--script-limit=1", "--script-format=bash")
	cmd.Dir = env
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run failed: %v\n%s", err, out)
	}
	s := string(out)
	if !strings.Contains(s, "tk show A") {
		t.Fatalf("missing tk show command for A:\n%s", s)
	}
	if !strings.Contains(s, "# To claim: tk start A") {
		t.Fatalf("missing tk claim command for A:\n%s", s)
	}
	if strings.Contains(s, "br show A") {
		t.Fatalf("did not expect beads command in tk mode:\n%s", s)
	}
}
