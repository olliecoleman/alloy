package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/olliecoleman/alloy/app/models"
	"github.com/olliecoleman/alloy/app/views"
	"github.com/pkg/errors"
)

func ShowPage(w http.ResponseWriter, r *http.Request) error {
	slug := chi.URLParam(r, "slug")
	slug = strings.Split(slug, "#")[0]

	page, err := models.GetPageBySlug(slug)
	if err != nil {
		return StatusError{Code: 404, Err: errors.Wrapf(err, "slug requested: %s", slug)}
	}

	v := views.New(r)
	v.Vars["Page"] = page

	s := page.Layout.String
	if s == "" {
		s = "two-col"
	}

	v.Render(w, fmt.Sprintf("pages/show-%s", s))
	return nil
}
