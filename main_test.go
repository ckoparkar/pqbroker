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
	failIf(t, err)

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}
}

func TestCatalog(t *testing.T) {
	server := httptest.NewServer(Routes())
	defer server.Close()

	url := server.URL + "/v2/catalog"
	req, err := http.NewRequest("GET", url, nil)
	failIf(t, err)

	req.SetBasicAuth("admin", "admin")
	cli := &http.Client{}
	res, err := cli.Do(req)
	failIf(t, err)

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	failIf(t, err)

	content := string(body)
	if !strings.Contains(content, "postgresql-db") {
		t.Error(err)
	}
}
