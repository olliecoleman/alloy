package models_test

import (
	"os"
	"testing"

	"github.com/olliecoleman/alloy/testutils"
	"github.com/gobuffalo/envy"
)

func TestMain(m *testing.M) {
	envy.Set("MIGRATIONS_DIR", "../migrations")
	testutils.SetupDB()

	c := m.Run()

	testutils.DropDB()
	os.Exit(c)
}
