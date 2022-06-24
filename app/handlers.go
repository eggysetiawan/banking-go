package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name       string `json:"name" xml:"name"`
	City       string `json:"city" xml:"city"`
	PostalCode string `json:"postalCode" xml:"postalcode"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func indexCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Rahmat", "Jakarta", "13460"},
		{"Dani", "Kediri", "22450"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func storeCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You hit me!")
}

func showCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprint(w, vars["customer"])
}
