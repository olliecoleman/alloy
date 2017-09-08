package admin

import (
	"net/http"

	"github.com/olliecoleman/alloy/app/views"
)

func Dashboard(w http.ResponseWriter, r *http.Request) error {
	v := views.New(r)
	v.Render(w, "admin/dashboard/index")
	return nil
}
