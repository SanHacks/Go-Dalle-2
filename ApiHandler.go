package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func platformRouter(router *mux.Router) {
	//Map the routes to the handlers in the backend api handler
	router.HandleFunc("/backend/api/products", ProductsHandler).Methods("GET")
	router.PathPrefix("/generatedProducts/").Handler(http.StripPrefix("/generatedProducts/", http.FileServer(http.Dir("./generatedProducts/"))))

}
