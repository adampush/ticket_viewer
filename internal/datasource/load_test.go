package datasource

import (
	"testing"
)

func TestTestModeLegacyFallbackEnabled(t *testing.T) {
	t.Setenv("TKV_TEST_MODE", "")
	if testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback disabled for empty TKV_TEST_MODE")
	}

	t.Setenv("TKV_TEST_MODE", "1")
	if !testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback enabled for TKV_TEST_MODE=1")
	}

	t.Setenv("TKV_TEST_MODE", "true")
	if testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback disabled for TKV_TEST_MODE=true")
	}

	t.Setenv("TKV_TEST_MODE", " 1 ")
	if !testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback enabled for spaced TKV_TEST_MODE=1")
	}
}
