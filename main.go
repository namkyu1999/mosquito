package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	err := initializeTransactionLog()
	if err != nil {
		panic(err)
	}

	// Create a new mux router
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/v1/{key}", keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")

	r.HandleFunc("/v1", notAllowedHandler)
	r.HandleFunc("/v1/{key}", notAllowedHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
