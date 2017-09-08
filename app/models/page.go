package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/olliecoleman/alloy/app/services"
	"github.com/lib/pq"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
)

type Page struct {
	ID              int64
	Title           nulls.String      `db:"title"`
	PageTitle       nulls.String      `db:"page_title"`
	MetaDescription nulls.String      `db:"meta_description"`
	Content         nulls.String      `db:"content"`
	Slug            nulls.String      `db:"slug"`
	Layout          nulls.String      `db:"layout"`
	InsertedAt      time.Time         `db:"inserted_at"`
	UpdatedAt       time.Time         `db:"updated_at"`
	Errors          map[string]string `db:"-"`
}

func NewPage() *Page {
	return &Page{
		Layout: nulls.NewString("two-col"),
	}
}

func ListPages(offset, limit int) ([]*Page, int, error) {
	pages := []*Page{}
	err := services.DB.Select(&pages, `
		SELECT 
			id, title, slug, inserted_at
		FROM pages 
		ORDER BY inserted_at DESC
		OFFSET $1 LIMIT $2
	`, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	var count int
	err = services.DB.Get(&count, `SELECT count(*) FROM pages`)
	if err != nil {
		return nil, 0, err
	}

	return pages, count, nil
}

func ListAllPages() ([]*Page, error) {
	pages := []*Page{}
	err := services.DB.Select(&pages, `
		SELECT 
			id, title, slug, inserted_at
		FROM pages 
		ORDER BY inserted_at DESC
	`)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func GetPage(ID int) (*Page, error) {
	page := Page{}
	err := services.DB.Get(&page, `
		SELECT 
			id, title, page_title, meta_description, content, slug, layout, inserted_at
		FROM pages 
		WHERE id = $1
	`, ID)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func GetPageBySlug(slug string) (*Page, error) {
	page := Page{}
	err := services.DB.Get(&page, `
		SELECT 
			id, title, page_title, meta_description, content, slug, layout, inserted_at
		FROM pages 
		WHERE slug = $1
	`, slug)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func (p *Page) Create() error {
	p.InsertedAt = time.Now()
	p.UpdatedAt = time.Now()
	stmt, err := services.DB.PrepareNamed(`
		INSERT INTO pages (slug, title, page_title, meta_description, layout, content, inserted_at, updated_at)
		VALUES 			  (:slug, :title, :page_title, :meta_description, :layout, :content, :inserted_at, :updated_at)
		RETURNING id
	`)

	if err != nil {
		return errors.WithStack(err)
	}

	err = stmt.Get(&p.ID, p)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23505" && pgerr.Constraint == "pages_slug_index" {
				return ErrAlreadyTaken
			}
		}

		return errors.WithStack(err)
	}

	return nil
}

func (p *Page) Delete() error {
	_, err := services.DB.Exec("DELETE from pages WHERE pages.id = $1", p.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *Page) Update() error {
	p.UpdatedAt = time.Now()

	stmt, err := services.DB.PrepareNamed(`
		UPDATE pages 
		SET
			title = :title,
			page_title = :page_title,
			meta_description = :meta_description,
			layout = :layout,
			content = :content,
			slug = :slug,
			updated_at = :updated_at
		WHERE id = :id
		RETURNING id
	`)

	if err != nil {
		return errors.WithStack(err)
	}

	err = stmt.Get(p, p)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23505" && pgerr.Constraint == "pages_slug_index" {
				return ErrAlreadyTaken
			}
		}

		return errors.WithStack(err)
	}

	return nil
}

func (p *Page) Validate() bool {
	p.Errors = make(map[string]string)

	title := strings.TrimSpace(p.Title.String)
	slug := strings.TrimSpace(p.Slug.String)
	content := strings.TrimSpace(p.Content.String)

	if title == "" {
		p.Errors["Title"] = "can't be blank"
	}

	if slug == "" {
		p.Errors["Slug"] = "can't be blank"
	}

	if content == "" {
		p.Errors["Content"] = "can't be blank"
	}

	re := regexp.MustCompile("^[\\w\\-]+$")
	matched := re.Match([]byte(slug))

	if matched == false {
		p.Errors["Slug"] = "URL slug is invalid (only a-z, 0-9 and - allowed)"
	}

	if p.Layout.String == "" {
		p.Errors["Layout"] = "must be valid"
	}

	return len(p.Errors) == 0
}
