package models_test

import (
	"testing"

	"github.com/olliecoleman/alloy/app/models"
	"github.com/olliecoleman/alloy/testutils"
	"github.com/markbates/pop/nulls"
)

func Test_CheckAdminCredentials(t *testing.T) {
	admin := models.AdminUser{
		Email:    nulls.NewString("admin@example.com"),
		Password: "testing123",
	}
	admin.Create()
	defer testutils.ResetTable("admin_users")

	res := admin.CheckAuth("testing123")
	if res == false {
		t.Error("Incorrect matching.")
	}

	res = admin.CheckAuth("foobar")
	if res != false {
		t.Error("Incorrect matching.")
	}
}
