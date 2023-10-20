package main

import (
	"goproject/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// SERVER
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Handler).Methods("POST")
	r.HandleFunc("/register", handlers.GetHandler).Methods("GET")

	http.ListenAndServe(":8080", r)

}
