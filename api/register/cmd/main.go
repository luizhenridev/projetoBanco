package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	ac := Client{Account_id: 1, CPF: "00000000", Name: "Luiz", Email: "luiz.henri882@gmail.com", Address: Address{Street: "Rua Joaquim Jerônimo", Number: 386}}
	readerString := NewClient(ac)

	reader := strings.NewReader(readerString)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8080/register", reader)

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
	println(string(body))

}



func NewClient(c Client) string {
	res, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(res))
	return string(res)

}
