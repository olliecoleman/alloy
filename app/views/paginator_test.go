package views

import (
	"net/url"
	"strings"
	"testing"
)

func Test_NewPaginator(t *testing.T) {
	u, _ := url.Parse("http://localhost:3000/test")
	p, _ := NewPaginator(5, 15, u)

	if p.End != 75 {
		t.Errorf("expected end to be %d, got %d", 75, p.End)
	}

	if p.Offset != 60 {
		t.Errorf("expected offset to be %d, got %d", 60, p.Offset)
	}
}

func Test_Render(t *testing.T) {
	u, _ := url.Parse("http://localhost:3000/test")
	p, _ := NewPaginator(5, 15, u)
	res := p.Render(100)
	if !p.HasNext {
		t.Error("Should have next items")
	}
	if !p.HasPrev {
		t.Error("Should have prev items")
	}

	if !strings.Contains(res, "5 / 7") {
		t.Errorf("%s \n Should have current page and total page count.", res)
	}

	p, _ = NewPaginator(5, 15, u)
	res = p.Render(70)
	if p.HasNext {
		t.Error("Should not have next items")
	}
	if !p.HasPrev {
		t.Error("Should have prev items")
	}

	p, _ = NewPaginator(5, 15, u)
	res = p.Render(70)
	if p.NextPageURL != "http://localhost:3000/test?page=6" {
		t.Errorf("Next page url incorrect: %s", p.NextPageURL)
	}
	if p.PrevPageURL != "http://localhost:3000/test?page=4" {
		t.Errorf("Prev page url incorrect: %s", p.PrevPageURL)
	}

}
