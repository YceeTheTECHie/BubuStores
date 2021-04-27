package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
 
func main() {

	r := mux.NewRouter()

	// defining routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/products", getProducts).Methods("GET")
	api.HandleFunc("/products/create", createProduct).Methods("POST")
	api.HandleFunc("/products/{id}", getProduct).Methods("GET")
	api.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	api.HandleFunc("/products", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}

