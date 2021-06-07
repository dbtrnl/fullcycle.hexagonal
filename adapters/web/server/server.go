package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/dnbtr/fullcycle.hexagonal/adapters/web/server/handler"
	"github.com/dnbtr/fullcycle.hexagonal/application"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	// Mux to Route
	router := mux.NewRouter()

	// Negroni to define middlewares
	negroni := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(router, negroni, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil { log.Fatal(err) }
}
