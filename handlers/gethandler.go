package handlers

import (
	"encoding/json"
	"fmt"
	"goproject/api/register/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

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
	err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
	if err != nil {
		log.Println(err)
	}

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
