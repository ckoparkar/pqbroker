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

	cli := &http.Client{}
	res, err := cli.Do(req)

	// should get an 401
	if res.StatusCode != http.StatusUnauthorized {
		t.Error(res.StatusCode)
	}

	req.SetBasicAuth("admin", "admin")
	res, err = cli.Do(req)
	failIf(t, err)

	// should return 200
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
