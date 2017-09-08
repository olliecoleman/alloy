package views

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"math"
	"net/url"
	"strconv"
)

type Paginator struct {
	Page        int
	PerPage     int
	Offset      int
	Limit       int
	NumPages    int
	Start       int
	End         int
	HasPrev     bool
	HasNext     bool
	BaseURL     *url.URL
	PrevPageURL string
	NextPageURL string
}

var (
	tmpl = `
		{{ if .NumPages }}
		<nav class="app-pagination">
			<ul class="pagination justify-content-center">
				{{ if .HasPrev }}
					<li class="page-item"><a class="page-link" href="{{.PrevPageURL}}"><span aria-hidden="true">&larr;</span> Previous</a></li>
				{{ else }}
					<li class="page-item disabled"><a class="page-link" href="#"><span aria-hidden="true">&larr;</span> Previous</a></li>
				{{ end }}

				<li class="page-item disabled">
      				<a href="#" class="page-link">{{ .Page }} / {{ .NumPages }}</a>
    			</li>

				{{ if .HasNext }}
					<li class="page-item"><a class="page-link" href="{{.NextPageURL}}">Next <span aria-hidden="true">&rarr;</span></a></li>
				{{ else }}
					<li class="page-item disabled"><a class="page-link" href="#">Next <span aria-hidden="true">&rarr;</span></a></li>
				{{ end }}
			</ul>
		</nav>
		{{ end }}
	`
)

func NewPaginator(page, perPage int, url *url.URL) (*Paginator, error) {
	if page < 1 {
		return nil, errors.New("Invalid page number")
	}

	p := &Paginator{
		Page:        page,
		PerPage:     perPage,
		Offset:      (page - 1) * perPage,
		Limit:       perPage,
		HasPrev:     false,
		HasNext:     false,
		BaseURL:     url,
		PrevPageURL: "",
		NextPageURL: "",
	}

	p.Start = p.Offset
	p.End = p.Start + perPage

	return p, nil
}

func (p *Paginator) Render(totalItems int) string {
	if p.Page > 1 {
		p.HasPrev = true
	}

	if p.End < totalItems {
		p.HasNext = true
	}

	p.NumPages = int(math.Ceil(float64(totalItems) / float64(p.PerPage)))

	if p.Page > p.NumPages {
		return ""
	}

	p.PrevPageURL = p.PageLink(p.Page - 1)
	p.NextPageURL = p.PageLink(p.Page + 1)

	t, err := template.New("*").Parse(tmpl)
	if err != nil {
		log.Println(err)
		return ""
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, p); err != nil {
		log.Println(err)
		return ""
	}
	return buf.String()
}

func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.BaseURL.String())
	values := link.Query()
	if page == 1 {
		values.Del("page")
	} else {
		values.Set("page", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return link.String()
}
