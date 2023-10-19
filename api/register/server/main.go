package main

import (
	"log"
	"net/http"
	"time"
	"goproject/api/register/client"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", Handler)
	http.ListenAndServe(":8080", mux)
	ac := client.Client{Account_id: 1, CPF: "00000000", Name: "Luiz", Email: "luiz.henri882@gmail.com", Address: client.Address{Street: "Rua Joaquim Jerônimo", Number: 386}}
	client.NewClient(ac)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")
	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("Request Processada com sucesso")
		//Imprime no browser
		w.Write([]byte("Request Processada com sucesso"))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("Request cancelada pelo client")

	}
}
