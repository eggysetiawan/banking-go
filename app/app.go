package app

import (
	"log"
	"net/http"

	"github.com/eggysetiawan/banking-go/domain"
	"github.com/eggysetiawan/banking-go/service"
	"github.com/gorilla/mux"
)

func Start() {
	// router := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// defines route
	router.HandleFunc("/customers", ch.indexCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer:[0-9]+}", ch.showCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
