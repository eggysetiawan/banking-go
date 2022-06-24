package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// defines route
	http.HandleFunc("/index", index)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
