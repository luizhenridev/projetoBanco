package main

import (
	"goproject/handlers/registers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// SERVER
	r := mux.NewRouter()

	r.HandleFunc("/register", registers.Handler).Methods("POST")
	r.HandleFunc("/register", registers.GetHandler).Methods("GET")
	r.HandleFunc("/register/{id}", registers.GetHandlerID).Methods("GET")
	r.HandleFunc("/register/{id}", registers.UpdateHandler).Methods("PUT")
	r.HandleFunc("/register/{id}", registers.DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
