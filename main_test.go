package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setCorrectBasicAuth(r *http.Request) {
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("admin:admin")))
}

func TestHelloWorld(t *testing.T) {
	server := httptest.NewServer(Router())
	defer server.Close()

	res, _ := http.Get(server.URL)

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}
}

func TestAuth(t *testing.T) {
	ts := Router()

	r, _ := http.NewRequest("GET", "/v2/catalog", nil)
	w := httptest.NewRecorder()

	ts.ServeHTTP(w, r)
	// should return 401
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Got %d, wanted 401.", w.Code)
	}

	// should return 200
	w = httptest.NewRecorder()
	setCorrectBasicAuth(r)

	ts.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Got %d, wanted 200.", w.Code)
	}
}

func TestCatalog(t *testing.T) {
	ts := Router()
	r, _ := http.NewRequest("GET", "/v2/catalog", nil)
	w := httptest.NewRecorder()

	setCorrectBasicAuth(r)
	ts.ServeHTTP(w, r)

	b, _ := ioutil.ReadAll(w.Body)
	body := string(b)

	if !strings.Contains(body, "postgresql-db") {
		t.Errorf("Expected body to contain postgresql-db. Got %s", body)
	}
}
