package main

import (
"fmt"
"net/http"
)

// ProductsHandler Push products as requested from endpoint as JSON
func ProductsHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, ProductsGenerator(w))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}