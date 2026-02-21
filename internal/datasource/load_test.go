package datasource

import (
	"testing"
)

func TestTestModeLegacyFallbackEnabled(t *testing.T) {
	t.Setenv("BV_TEST_MODE", "")
	if testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback disabled for empty BV_TEST_MODE")
	}

	t.Setenv("BV_TEST_MODE", "1")
	if !testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback enabled for BV_TEST_MODE=1")
	}

	t.Setenv("BV_TEST_MODE", "true")
	if testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback disabled for BV_TEST_MODE=true")
	}

	t.Setenv("BV_TEST_MODE", " 1 ")
	if !testModeLegacyFallbackEnabled() {
		t.Fatalf("expected fallback enabled for spaced BV_TEST_MODE=1")
	}
}
