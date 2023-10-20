package main

import (
	"encoding/json"
	"fmt"
	"goproject/api/register/client"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	// SERVER

	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	mux.HandleFunc("/register", Handler)

	http.ListenAndServe(":8080", mux)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	fmt.Println("Valor de r", r)

	var request client.Client

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("valor de r.Body decode", request)

	response := client.ClientResponse{
		Account_id: uuid.New().String(),
		CPF:        request.CPF,
		Name:       request.Name,
		Email:      request.Email,
		Address:    client.Address{Street: "Rua Joaquim Jerônimo", Number: 386}}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("Request Processada com sucesso")
		//Imprime no browser
		//w.Write([]byte("Request Processada com sucesso"))
		w.Write([]byte(response))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("Request cancelada pelo client")

	}
}
