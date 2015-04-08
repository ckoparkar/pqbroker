package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func createInstance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dbname := fmt.Sprintf("d%s", ps.ByName("instance_id"))
	dashboard_url, err := createDatabase(strings.Replace(dbname, "-", "_", -1))

	response := make(map[string]string)

	if err == nil {
		response["dashboard_url"] = dashboard_url
	} else {
		w.WriteHeader(err.Code)
		response["description"] = err.Err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(response)
	w.Write(j)
}

func deleteInstance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dbname := fmt.Sprintf("d%s", ps.ByName("instance_id"))

	response := make(map[string]string)
	err := deleteDatabase(strings.Replace(dbname, "-", "_", -1))
	if err != nil {
		w.WriteHeader(err.Code)
		response["description"] = err.Err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(response)
	w.Write(j)
}

func Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", helloWorld)
	router.GET("/v2/catalog", basicAuth(catalog))
	router.PUT("/v2/service_instances/:instance_id", basicAuth(createInstance))
	router.DELETE("/v2/service_instances/:instance_id", basicAuth(deleteInstance))
	router.PUT("/v2/service_instances/:instance_id/service_bindings/:binding_id",
		basicAuth(createBinding))
	return router
}

func main() {
	router := Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
