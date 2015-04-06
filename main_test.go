package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setCorrectBasicAuth(r *http.Request) {
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("admin:admin")))
}

func TestHelloWorld(t *testing.T) {
	server := httptest.NewServer(Router())
	defer server.Close()

	res, err := http.Get(server.URL)
	failIf(t, err)

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}
}

func TestCatalog(t *testing.T) {
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
