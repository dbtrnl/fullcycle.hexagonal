package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dnbtr/fullcycle.hexagonal/adapters/web/dto"
	"github.com/dnbtr/fullcycle.hexagonal/application"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
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

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var productDto dto.Product

		// Decoding JSON body and assigning it to productDto
		err := json.NewDecoder(request.Body).Decode(&productDto)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return;
		}

		// Create the Product on the Database
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return;
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return;
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Parsing the URL and extracting the id
		muxVars := mux.Vars(request)
		id := muxVars["id"]

		// Getting the product by ID
		product, err := service.Get(id)
		if err != nil {	writer.WriteHeader(http.StatusNotFound);	return }

		result, err := service.Enable(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return;
		}

		// NewEncoder().Encode() returns the parsed JSON to the Writer
		err = json.NewEncoder(writer).Encode(result)
		if err != nil {	writer.WriteHeader(http.StatusInternalServerError);	return }
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Parsing the URL and extracting the id
		muxVars := mux.Vars(request)
		id := muxVars["id"]

		// Getting the product by ID
		product, err := service.Get(id)
		if err != nil {	writer.WriteHeader(http.StatusNotFound);	return }

		result, err := service.Disable(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return;
		}

		// NewEncoder().Encode() returns the parsed JSON to the Writer
		err = json.NewEncoder(writer).Encode(result)
		if err != nil {	writer.WriteHeader(http.StatusInternalServerError);	return }
	})
}