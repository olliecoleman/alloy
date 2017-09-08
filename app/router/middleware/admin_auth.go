package middleware

import (
	"context"
	"net/http"

	"github.com/olliecoleman/alloy/app/models"
	"github.com/olliecoleman/alloy/app/services"
	"github.com/olliecoleman/alloy/app/views"
)

func RequireAdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := services.Session(r)
		userID, ok := session.Values["admin_user_id"]

		if !ok {
			views.ErrorFlash(w, r, "Please login to view that resource.")
			http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
			return
		}

		// find user
		u, err := models.GetAdminUser(userID.(int64))
		if err != nil {
			views.ErrorFlash(w, r, "Please login to view that resource.")
			http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), services.SessKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
