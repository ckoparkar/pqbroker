package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func panicIf(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestHelloWorld(t *testing.T) {
	server := httptest.NewServer(Routes())
	defer server.Close()

	res, err := http.Get(server.URL)
	panicIf(t, err)

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}
}

func TestCatalog(t *testing.T) {
	server := httptest.NewServer(Routes())
	defer server.Close()

	url := server.URL + "/v2/catalog"
	req, err := http.NewRequest("GET", url, nil)
	panicIf(t, err)

	req.SetBasicAuth("admin", "admin")
	cli := &http.Client{}
	res, err := cli.Do(req)
	panicIf(t, err)

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	panicIf(t, err)

	content := string(body)
	if !strings.Contains(content, "postgresql-db") {
		t.Error(err)
	}
}
