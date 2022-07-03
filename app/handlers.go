package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/eggysetiawan/banking-go/service"
	"github.com/gorilla/mux"
)

// type Customer struct {
// 	Name       string `json:"name" xml:"name"`
// 	City       string `json:"city" xml:"city"`
// 	PostalCode string `json:"postalCode" xml:"postalcode"`
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) indexCustomer(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandlers) showCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["customer"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}
