package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Broker struct {
	router *httprouter.Router
}

func NewBroker() *Broker {
	b := new(Broker)
	b.router = httprouter.New()
	b.registerRoutes()
	return b
}

func (b *Broker) registerRoutes() {
	b.router.GET("/", b.serveHelloWorld)
}

func (b *Broker) serveHelloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, World")
}
