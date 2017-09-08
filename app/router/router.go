package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/envy"
	"github.com/gorilla/csrf"
	"github.com/olliecoleman/alloy/app/handlers"
	admin "github.com/olliecoleman/alloy/app/handlers/admin"
	mw "github.com/olliecoleman/alloy/app/router/middleware"
)

// New creates a new router instance and server
func New() error {
	mux := HandleRoutes()

	addr := ":" + envy.Get("PORT", "1212")
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// production server will use nginx + letsencrypt
	return server.ListenAndServe()
}

func HandleRoutes() *chi.Mux {
	r := chi.NewRouter()
	env := envy.Get("ENVIRONMENT", "development")

	// Middlewares
	r.Use(mw.MethodOverride)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RedirectSlashes)
	r.Use(mw.Recoverer)
	r.Use(middleware.DefaultCompress)

	// Base
	r.NotFound(handlers.NotFoundHandler)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		if env == "production" {
			w.Header().Set("Vary", "Accept-Encoding")
			w.Header().Set("Cache-Control", "public, max-age=604800")
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		fs := http.StripPrefix("/public", http.FileServer(http.Dir("public")))

		fs.ServeHTTP(w, r)
	})
	serveSingle(r, "/robots.txt", "public/robots.txt")
	serveSingle(r, "/favicon.ico", "public/favicon.ico")

	r.Route("/", func(r chi.Router) {
		r.NotFound(handlers.NotFoundHandler)

		csrfKey := []byte(envy.Get("CSRF_KEY", "g4k827b582367a77cb27d1e5dc268912"))

		if env != "test" {
			r.Use(csrf.Protect(
				csrfKey,
				csrf.Secure(false), // change to true after switching to https
			))
		}

		// home signup
		r.Get("/", h(handlers.Home))

		// contact
		r.Get("/contact/new", h(handlers.NewSupportMessage))
		r.Post("/contact", h(handlers.CreateSupportMessage))
		r.Get("/pages/{slug}", h(handlers.ShowPage))

		// Admin auth
		r.Get("/admin/sessions/new", h(admin.GetLogin))
		r.Post("/admin/sessions", h(admin.PostLogin))
		r.Delete("/admin/sessions", h(admin.Logout))

		r.Route("/admin", func(r chi.Router) {
			r.Use(mw.RequireAdminAuth)

			// Dashboard
			r.Get("/", h(admin.Dashboard))

			// Support Messages
			r.Route("/support-messages", func(r chi.Router) {
				r.Get("/", h(admin.ListMessages))
				r.Delete("/{ID}", h(admin.DeleteMessage))
			})

			// Pages
			r.Route("/pages", func(r chi.Router) {
				r.Get("/", h(admin.ListPages))
				r.Get("/new", h(admin.NewPage))
				r.Post("/", h(admin.CreatePage))

				r.Route("/{ID}", func(r chi.Router) {
					r.Use(admin.PageContext)
					r.Get("/", h(admin.GetPage))
					r.Get("/edit", h(admin.EditPage))
					r.Put("/", h(admin.UpdatePage))
					r.Delete("/", h(admin.DeletePage))
				})
			})

		})

	})

	return r
}

func h(fn handlers.Handler) http.HandlerFunc {
	return handlers.Handler(fn).ServeHTTP
}

func serveSingle(mux *chi.Mux, pattern, filename string) {
	mux.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
