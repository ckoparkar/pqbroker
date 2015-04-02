package main

import (
	"log"
	"net/http"

	"github.com/cskksc/postgresqlbroker/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", api.Handlers()))
}
