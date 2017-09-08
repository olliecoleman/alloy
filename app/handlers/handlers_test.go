package handlers_test

import (
	"os"
	"testing"

	"github.com/olliecoleman/alloy/app/views"
	"github.com/olliecoleman/alloy/testutils"
	"github.com/gobuffalo/envy"
)

func TestMain(m *testing.M) {
	envy.Set("MIGRATIONS_DIR", "../migrations")
	testutils.SetupServer()
	testutils.SetupDB()

	envy.Set("TEMPLATES_DIR", "../templates")
	views.LoadTemplates()

	c := m.Run()

	testutils.DropDB()
	testutils.CloseServer()
	os.Exit(c)
}
