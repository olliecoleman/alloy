package admin

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/olliecoleman/alloy/app/handlers"
	"github.com/olliecoleman/alloy/app/models"
	"github.com/olliecoleman/alloy/app/views"
	"github.com/pkg/errors"
)

func ListMessages(w http.ResponseWriter, r *http.Request) error {
	v := views.New(r)
	pageNum := handlers.GetPageNum(r)

	p, err := views.NewPaginator(pageNum, handlers.PerPage, r.URL)
	if err != nil {
		return handlers.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	messages, numMessages, err := models.ListMessages(p.Start, p.Limit)

	if err != nil {
		return handlers.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	pagination := p.Render(numMessages)
	if err != nil {
		return handlers.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	v.Vars["Messages"] = messages
	v.Vars["Pagination"] = template.HTML(pagination)
	v.Render(w, "admin/support-messages/index")
	return nil
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		return errors.WithStack(err)
	}

	message := models.SupportMessage{ID: int64(id)}
	err = message.Delete()
	if err != nil {
		return handlers.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	views.SuccessFlash(w, r, "Support message was deleted successfully.")
	http.Redirect(w, r, "/admin/support-messages", http.StatusSeeOther)
	return nil
}
