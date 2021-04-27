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

	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/products", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}

