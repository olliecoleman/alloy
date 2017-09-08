package views

import (
	"net/http"
	"testing"
)

func Test_NewView(t *testing.T) {
	req, err := http.NewRequest("get", "http://localhost:3000/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	v := New(req)
	if !v.IsCurrentURL("/home") {
		t.Fail()
	}

	if v.IsCurrentURL("/") {
		t.Fail()
	}
}
