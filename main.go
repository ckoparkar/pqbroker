package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func helloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, World")
}

func catalog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	catalog, err := Asset("config/settings.json")
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(catalog))
}

func Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", helloWorld)
	router.GET("/v2/catalog", basicAuth(catalog))
	return router
}

func main() {
	router := Routes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
