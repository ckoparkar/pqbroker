package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	server := httptest.NewServer(Routes())
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}
}

func TestCatalog(t *testing.T) {
	server := httptest.NewServer(Routes())
	defer server.Close()

	res, err := http.Get(server.URL + "/v2/catalog")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	content := string(body)
	if !strings.Contains(content, "postgresql-db") {
		t.Error(err)
	}
}
