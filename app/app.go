package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	// defines route
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/customers", customer)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
