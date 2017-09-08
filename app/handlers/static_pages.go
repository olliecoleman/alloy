package handlers

import (
	"net/http"

	"github.com/olliecoleman/alloy/app/views"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	v := views.New(r)
	v.Render(w, "home")
	return nil
}
