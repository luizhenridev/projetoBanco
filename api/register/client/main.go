package main

import (
	"bytes"
	"context"
	"fmt"
	"goproject/api/register/models"

	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	//CLIENT
	ac := models.Client{
		CPF:     "00000000",
		Name:    "Luiz",
		Email:   "luiz.henri882@gmail.com",
		Address: models.Address{Street: "Rua Joaquim Jerônimo", Number: 386}}
	reader := models.NewClientJson(ac)
	ctxClient, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Reader", reader)

	req, err := http.NewRequestWithContext(ctxClient, "POST", "http://localhost:8080", bytes.NewBuffer(reader))
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
