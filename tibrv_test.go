package tibrv

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Setenv("TEST_SERVICE", "")
	os.Setenv("TEST_NETWORK", "")
	os.Setenv("TEST_DAEMON", "")

	code := m.Run()

	os.Exit(code)
}
