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
	r.HandleFunc("/register/{id}", handlers.GetHandlerID).Methods("GET")
	r.HandleFunc("/register/{id}", handlers.UpdateHandler).Methods("PUT")
	r.HandleFunc("/register/{id}", handlers.DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
