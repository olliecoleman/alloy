package views

import (
	"fmt"
	"html/template"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/microcosm-cc/bluemonday"
)

var templates map[string]*template.Template
var funcs = template.FuncMap{
	"assetURL":     assetURLFn,
	"empty":        emptyFn,
	"date":         dateFn,
	"readableDate": readableDateFn,
	"raw":          raw,
	"lower":        lower,
	"upper":        upper,
}

func LoadTemplates() {
	defaults := []string{"layouts/application.html", "shared/header.html", "shared/footer.html"}
	admin := []string{"layouts/admin.html", "admin/shared/header.html", "admin/shared/footer.html"}
	mailer := []string{"mailer/layouts/transactional.html", "mailer/shared/button.html", "mailer/shared/unsubscribe.html", "mailer/shared/footer.html"}

	templates = map[string]*template.Template{
		"404":                                  newTemplate("layouts/application.html", "errors/404.html"),
		"500":                                  newTemplate("layouts/application.html", "errors/500.html"),
		"401":                                  newTemplate("layouts/application.html", "errors/401.html"),
		"home":                                 newTemplate(append(defaults, "static-pages/home.html")...),
		"pages/show-full":                      newTemplate(append(defaults, "pages/show-full.html")...),
		"pages/show-two-col":                   newTemplate(append(defaults, "pages/show-two-col.html")...),
		"support-messages/new":                 newTemplate(append(defaults, "support-messages/new.html")...),
		"admin/sessions/new":                   newTemplate("layouts/admin.html", "admin/sessions/new.html"),
		"admin/dashboard/index":                newTemplate(append(admin, "admin/dashboard/index.html")...),
		"admin/support-messages/index":         newTemplate(append(admin, "admin/support-messages/index.html")...),
		"admin/pages/index":                    newTemplate(append(admin, "admin/pages/index.html")...),
		"admin/pages/show":                     newTemplate(append(admin, "admin/pages/show.html")...),
		"admin/pages/edit":                     newTemplate(append(admin, "admin/pages/_form.html", "admin/pages/edit.html")...),
		"admin/pages/new":                      newTemplate(append(admin, "admin/pages/_form.html", "admin/pages/new.html")...),
		"mailer/support-messages/new":          newTemplate(append(mailer, "mailer/support-messages/new.html")...),
		"mailer/support-messages/notification": newTemplate(append(mailer, "mailer/support-messages/notification.html")...),
	}
}

func newTemplate(files ...string) *template.Template {
	f := []string{}
	cwd := envy.Get("TEMPLATES_DIR", "app/templates")

	for _, file := range files {
		f = append(f, filepath.Join(cwd, file))
	}

	return template.Must(template.New("*").Funcs(funcs).ParseFiles(f...))
}

// github.com/Masterminds/sprig/blob/master/defaults.go
func emptyFn(input interface{}) bool {
	g := reflect.ValueOf(input)
	if !g.IsValid() {
		return true
	}

	switch g.Kind() {
	default:
		return g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		return g.Bool() == false
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Struct:
		return false
	}
}

func dateFn(format string, input interface{}) string {
	var t time.Time
	switch date := input.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	}

	return t.Format(format)
}

func readableDateFn(t time.Time) string {
	if time.Now().Before(t) {
		return ""
	}
	diff := time.Now().Sub(t)
	day := 24 * time.Hour
	month := 30 * day
	year := 12 * month

	switch {
	case diff < time.Second:
		return "just now"
	case diff < 5*time.Minute:
		return "a few minutes ago"
	case diff < time.Hour:
		return fmt.Sprintf("%d minutes ago", diff/time.Minute)
	case diff < day:
		return fmt.Sprintf("%d hours ago", diff/time.Hour)
	case diff < month:
		return fmt.Sprintf("%d days ago", diff/day)
	case diff < year:
		return fmt.Sprintf("%d months ago", diff/month)
	default:
		return fmt.Sprintf("%d years ago", diff/year)
	}
}

func raw(input string) template.HTML {
	p := bluemonday.UGCPolicy()
	return template.HTML(p.Sanitize(input))
}

func lower(input string) string {
	return strings.ToLower(input)
}

func upper(input string) string {
	return strings.ToUpper(input)
}

func assetURLFn(input string) string {
	url := envy.Get("ASSET_URL", "")
	return fmt.Sprintf("%s%s", url, input)
}
