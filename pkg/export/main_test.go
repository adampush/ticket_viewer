package export

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Prevent any test from accidentally opening a browser
	os.Setenv("TKV_NO_BROWSER", "1")
	os.Setenv("TKV_TEST_MODE", "1")

	os.Exit(m.Run())
}
