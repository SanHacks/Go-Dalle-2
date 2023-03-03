package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Map the routes to the handlers in the backend api handler
// Handle 404 errors and redirect to /platform
func platformRouter(router *mux.Router) {
	router.NotFoundHandler = routeNotFoundError()
	//Backend API Routes Definition (API) (ie: open route and map to handler)
	router.HandleFunc("/backend/api/products", ProductsHandler).Methods("GET")

	//File Server for Platform CDN and Frontend Engine Resources (CSS, JS, Images, Fonts and stuff)
	router.PathPrefix("/generatedProducts/").Handler(http.StripPrefix("/generatedProducts/", http.FileServer(http.Dir("./generatedProducts/"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets/css/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./frontend/assets/img/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/frontend/assets/js/", http.FileServer(http.Dir("./frontend/assets/js/"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images /", http.FileServer(http.Dir("./frontend/assets/img/"))))
	router.PathPrefix("/frontend/assets/fonts/").Handler(http.StripPrefix("/frontend/assets/fonts/", http.FileServer(http.Dir("./frontend/assets/fonts/"))))
	router.PathPrefix("/frontend/assets/").Handler(http.StripPrefix("/frontend/assets/", http.FileServer(http.Dir("./frontend/assets/"))))
	router.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend/"))))
}
