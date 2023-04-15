package app_test

import (
	"os"
	"testing"
)

const defaultPhone = "+79000000000"

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
