package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld)
	return router
}
