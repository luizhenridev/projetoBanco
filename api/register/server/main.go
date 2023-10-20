package main

import (
	"encoding/json"
	"fmt"
	"goproject/api/register/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	// SERVER
	r := mux.NewRouter()

	r.HandleFunc("/register", Handler).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	var request models.Client

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	if request.CPF == "" && request.Name == "" && request.Email == "" {
		http.Error(w, "Parâmetros obrigatórios ausentes", http.StatusBadRequest)
		return
	} else if request.CPF == "" {
		http.Error(w, "Parâmetro CPF ausente", http.StatusBadRequest)
		return
	} else if request.Name == "" {
		http.Error(w, "Parâmetro Name ausente", http.StatusBadRequest)
		return
	} else if request.Email == "" {
		http.Error(w, "Parâmetro Email ausente", http.StatusBadRequest)
		return
	}

	response := models.ClientResponse{
		Account_id: uuid.New().String(),
		CPF:        request.CPF,
		Name:       request.Name,
		Email:      request.Email,
		Address:    models.Address{Street: request.Address.Street, Number: request.Address.Number}}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("Request Processada com sucesso")
		//Imprime no browser
		w.Write([]byte("Request Processada com sucesso"))
		//w.Write([]byte(response))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("Request cancelada pelo client")

	}
}
func GetHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	fmt.Println("Valor de r", r)

	var request models.Client

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("valor de r.Body decode", request)

	response := models.ClientResponse{
		Account_id: uuid.New().String(),
		CPF:        request.CPF,
		Name:       request.Name,
		Email:      request.Email,
		Address:    models.Address{Street: request.Address.Street, Number: request.Address.Number}}

	//jsonResponse, err := json.Marshal(response)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("Request Processada com sucesso")
		//Imprime no browser
		w.Write([]byte("Request Processada com sucesso"))
		//w.Write([]byte(response))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("Request cancelada pelo client")

	}
}
