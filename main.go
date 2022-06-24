package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name       string `json:"name"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
}

func main() {

	// defines route
	http.HandleFunc("/index", index)
	http.HandleFunc("/customers", customer)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func customer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Rahmat", "Jakarta", "13460"},
		{"Dani", "Kediri", "22450"},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
