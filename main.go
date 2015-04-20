package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	b := NewBroker()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), b.router))
}
