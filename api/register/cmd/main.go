package main

import (
	"bytes"
	"context"
	"fmt"
	"goproject/api/register/client"

	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	//CLIENT
	ac := client.Client{
		CPF:     "00000000",
		Name:    "Luiz",
		Email:   "luiz.henri882@gmail.com",
		Address: client.Address{Street: "Rua Joaquim Jerônimo", Number: 386}}
	reader := client.NewClientJson(ac)
	ctxClient, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
