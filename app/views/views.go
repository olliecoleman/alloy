package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gorilla/csrf"
	"github.com/olliecoleman/alloy/app/services"
)

type View struct {
	Vars    map[string]interface{}
	Request *http.Request
}

func (v *View) IsCurrentURL(url string) bool {
	return url == v.Request.URL.Path
}

func (v *View) Render(w http.ResponseWriter, name string) {
	t, err := GetTemplate(name)
	if err != nil {
		log.Printf("The template %s does not exist.\n", name)
		v.RenderError(w, 404)
		return
	}

	session := services.Session(v.Request)
	if flashes := session.Flashes(); len(flashes) > 0 {
		v.Vars["flashes"] = make([]Flash, len(flashes))
		for i, f := range flashes {
			switch f.(type) {
			case Flash:
				v.Vars["flashes"].([]Flash)[i] = f.(Flash)
			default:
				v.Vars["flashes"].([]Flash)[i] = Flash{f.(string), "alert-info"}
			}
		}
		session.Save(v.Request, w)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "base", v.Vars)
}

func (v *View) RenderError(w http.ResponseWriter, status int) {
	var name string
	switch status {
	case http.StatusNotFound:
		name = "404"
	case http.StatusUnauthorized:
		name = "401"
	default:
		name = "500"
	}

	t, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The %s does not exist.", name), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	t.ExecuteTemplate(w, "base", v.Vars)
}

func New(req *http.Request) *View {
	return &View{
		Vars: map[string]interface{}{
			"CSRF_TOKEN": csrf.Token(req),
			"CSRF":       csrf.TemplateField(req),
			"Meta": map[string]interface{}{
				"Env":  envy.Get("ENVIRONMENT", "development"),
				"Path": req.URL.Path,
			},
		},
		Request: req,
	}
}

func GetTemplate(templateName string) (*template.Template, error) {
	t, ok := templates[templateName]
	if !ok {
		return nil, fmt.Errorf("the template %s does not exist", templateName)
	}

	return t, nil
}
