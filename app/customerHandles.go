package app

import (
	"encoding/json"
	"encoding/xml"
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

func (ch *CustomerHandlers) IndexCustomerActive(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomerActive()
	if err != nil {
		writeResponse(w, r.Header.Get("Content-Type"), err.Code, err.AsMessage())
	} else {
		writeResponse(w, r.Header.Get("Content-Type"), http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) indexCustomer(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, r.Header.Get("Content-Type"), err.Code, err.AsMessage())
	} else {
		writeResponse(w, r.Header.Get("Content-Type"), http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) indexCustomerInactive(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomerInactive()

	if err != nil {
		writeResponse(w, r.Header.Get("Content-Type"), err.Code, err.AsMessage())
	} else {
		writeResponse(w, r.Header.Get("Content-Type"), http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) showCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["customer"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, r.Header.Get("Content-Type"), err.Code, err.AsMessage())
	} else {
		writeResponse(w, r.Header.Get("Content-Type"), http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, header string, code int, data interface{}) {
	w.Header().Add("Content-Type", header)

	w.WriteHeader(code)

	if header == "application/xml" {
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}

	}
}
