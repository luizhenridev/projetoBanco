package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"goproject/api/register/client"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	// SERVER

	mux := http.NewServeMux()
	mux.HandleFunc("/register", Handler)
	http.ListenAndServe(":8080", mux)

	//CLIENT
	ac := client.Client{
		CPF:     "00000000",
		Name:    "Luiz",
		Email:   "luiz.henri882@gmail.com",
		Address: client.Address{Street: "Rua Joaquim Jerônimo", Number: 386}}
	reader := client.NewClientJson(ac)
	ctxClient, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("Reader", reader)

	req, err := http.NewRequestWithContext(ctxClient, "POST", "http://localhost:8080/register", bytes.NewBuffer(reader))
	if err != nil {
		log.Println(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	println("Resposta", string(body))

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

	bodyRes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("valor de r.Body", bodyRes)
	decoder, err := json.Marshal(bodyRes)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("vaalor de decoder", decoder)

	response := client.ClientResponse{
		Account_id: uuid.New().String(),
		CPF:        request.CPF,
		Name:       request.Name,
		Email:      request.Email,
		Address:    client.Address{Street: "Rua Joaquim Jerônimo", Number: 386}}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Defina o código de status antes de escrever o corpo

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("Request Processada com sucesso")
		//Imprime no browser
		//w.Write([]byte("Request Processada com sucesso"))
		w.Write([]byte(jsonResponse))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("Request cancelada pelo client")

	}
}
