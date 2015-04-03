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

func Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", helloWorld)
	return router
}

func main() {
	router := Routes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
