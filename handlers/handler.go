package handlers

import (
	"encoding/json"
	"goproject/api/register/models"
	"goproject/db"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	var request models.Client

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	if request.CPF == "" && request.Name == "" && request.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_PARAMETERS",
			ErrorMessage: "Parâmetro parâmetros ausentes",
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	} else if request.CPF == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_CPF",
			ErrorMessage: "Parâmetro CPF ausente",
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	} else if request.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_NAME",
			ErrorMessage: "Parâmetro NAME ausente",
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	} else if request.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_EMAIL",
			ErrorMessage: "Parâmetro EMAIL ausente",
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	}

	response := models.ClientResponse{
		Account_id: uuid.New().String(),
		CPF:        request.CPF,
		Name:       request.Name,
		Email:      request.Email,
		Address:    models.Address{Street: request.Address.Street, Number: request.Address.Number}}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response) // Defina o código de status antes de escrever o corpo
	if err != nil {
		log.Println(err)
	}

	db.Insert(response)

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

	defer log.Println("Request Finalizada")

}
