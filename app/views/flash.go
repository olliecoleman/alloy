package views

import (
	"encoding/gob"
	"net/http"

	"github.com/olliecoleman/alloy/app/services"
)

var (
	FlashInfo    = "alert-info"
	FlashWarning = "alert-warning"
	FlashSuccess = "alert-success"
	FlashError   = "alert-danger"
)

type Flash struct {
	Message string
	Class   string
}

func init() {
	gob.Register(Flash{})
}

func ErrorFlash(w http.ResponseWriter, r *http.Request, message string) {
	sess := services.Session(r)
	sess.AddFlash(Flash{message, FlashError})
	sess.Save(r, w)
}

func SuccessFlash(w http.ResponseWriter, r *http.Request, message string) {
	sess := services.Session(r)
	sess.AddFlash(Flash{message, FlashSuccess})
	sess.Save(r, w)
}

func InfoFlash(w http.ResponseWriter, r *http.Request, message string) {
	sess := services.Session(r)
	sess.AddFlash(Flash{message, FlashInfo})
	sess.Save(r, w)
}
