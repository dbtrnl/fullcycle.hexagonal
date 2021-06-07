package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dnbtr/fullcycle.hexagonal/application"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Parsing the URL and extracting the id
		muxVars := mux.Vars(request)
		id := muxVars["id"]

		// Getting the product by ID
		product, err := service.Get(id)
		if err != nil {	writer.WriteHeader(http.StatusNotFound);	return }

		// NewEncoder().Encode() returns the parsed JSON to the Writer
		err = json.NewEncoder(writer).Encode(product)
		if err != nil {	writer.WriteHeader(http.StatusInternalServerError);	return }
	})
}