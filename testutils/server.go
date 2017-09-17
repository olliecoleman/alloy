package testutils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/olliecoleman/alloy/app/router"
)

var (
	server *httptest.Server
)

func SetupServer() {
	envy.Set("ENVIRONMENT", "test")
	server = httptest.NewServer(router.HandleRoutes())
}

func CloseServer() {
	server.Close()
}

func Get(t *testing.T, url string) (int, string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", server.URL, url), nil)

	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, string(b)
}

func PostForm(t *testing.T, url string, data url.Values) (int, string) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", server.URL, url), strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, string(b)
}

func PostJSON(t *testing.T, url string, jsonStr string) (int, string) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", server.URL, url), bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonStr)))

	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, string(b)
}
