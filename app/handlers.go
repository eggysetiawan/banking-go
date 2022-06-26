package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/eggysetiawan/banking-go/service"
)

type Customer struct {
	Name       string `json:"name" xml:"name"`
	City       string `json:"city" xml:"city"`
	PostalCode string `json:"postalCode" xml:"postalcode"`
}

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
