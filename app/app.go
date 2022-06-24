package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// router := http.NewServeMux()
	router := mux.NewRouter()

	// defines route
	router.HandleFunc("/", index).Methods(http.MethodGet)
	router.HandleFunc("/customers", indexCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", storeCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer:[0-9]+}", showCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
