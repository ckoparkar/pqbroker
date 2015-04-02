package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cskksc/postgresqlbroker/api"
)

func TestHelloWorld(t *testing.T) {
	server := httptest.NewServer(api.Handlers())
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.StatusCode)
	}
}
