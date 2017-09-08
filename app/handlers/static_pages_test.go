package handlers_test

import (
	"strings"
	"testing"

	"net/http"

	"github.com/olliecoleman/alloy/testutils"
)

func Test_Home(t *testing.T) {
	code, resp := testutils.Get(t, "/")

	if code != http.StatusOK {
		t.Errorf("Error code while fetching home: %d", code)
	}

	if !strings.Contains(resp, "Alloy") {
		t.Error("Home does not contain Alloy.")
	}
}
