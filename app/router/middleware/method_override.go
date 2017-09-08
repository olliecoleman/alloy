package middleware

import (
	"net/http"
	"strings"
)

func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			override := r.FormValue("_method")
			if override != "" {
				r.Method = strings.ToUpper(override)
				r.Form.Del("_method")
				r.PostForm.Del("_method")
			}
		}
		next.ServeHTTP(w, r)
	})
}
