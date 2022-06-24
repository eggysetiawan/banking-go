package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name       string `json:"name" xml:"name"`
	City       string `json:"city" xml:"city"`
	PostalCode string `json:"postalCode" xml:"postalcode"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func customer(w http.ResponseWriter, r *http.Request) {
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
