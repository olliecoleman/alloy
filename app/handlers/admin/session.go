package admin

import (
	"net/http"

	"github.com/olliecoleman/alloy/app/models"
	"github.com/olliecoleman/alloy/app/services"
	"github.com/olliecoleman/alloy/app/views"
)

func GetLogin(w http.ResponseWriter, r *http.Request) error {
	v := views.New(r)
	v.Vars["Email"] = ""
	v.Render(w, "admin/sessions/new")
	return nil
}

func PostLogin(w http.ResponseWriter, r *http.Request) error {
	session := services.Session(r)

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		loginErr(w, r, email)
		return nil
	}

	adminUser, err := models.GetAdminByEmail(email)
	if err != nil {
		loginErr(w, r, email)
		return nil
	}

	valid := adminUser.CheckAuth(password)

	if valid == true {
		session.Values["admin_user_id"] = adminUser.ID
		session.Save(r, w)

		if err != nil {
			loginErr(w, r, email)
		}

		views.SuccessFlash(w, r, "Logged in successfully")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return nil
	}

	loginErr(w, r, email)
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	session := services.Session(r)
	services.EmptySession(session)

	views.SuccessFlash(w, r, "Logged out successfully.")
	http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
	return nil
}

func loginErr(w http.ResponseWriter, r *http.Request, email string) {
	v := views.New(r)
	views.ErrorFlash(w, r, "Invalid credentials.")
	v.Vars["Email"] = email
	v.Render(w, "admin/sessions/new")
}
