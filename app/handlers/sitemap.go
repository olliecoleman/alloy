package handlers

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/envy"
	"github.com/olliecoleman/alloy/app/models"
	"github.com/pkg/errors"
)

func RenderSitemap(w http.ResponseWriter, r *http.Request) error {
	host := envy.Get("FULL_URL", "http://localhost:1212")
	cwd := envy.Get("TEMPLATES_DIR", "app/templates")
	t := template.Must(template.New("*").ParseFiles(filepath.Join(cwd, "sitemap/index.xml")))

	v := map[string]interface{}{}
	v["Host"] = host
	pages, err := models.ListAllPages()
	if err != nil {
		return errors.WithStack(err)
	}
	v["Pages"] = pages

	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "base", v)
	return nil
}
