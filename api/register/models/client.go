package models

import (
	"encoding/json"
	"log"
)

type Address struct {
	Street string `json:"street"`
	Number int    `json:"number"`
}

type Client struct {
	CPF     string  `json:"cpf"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"addres"`
}

type ClientResponse struct {
	Account_id string  `json:"account_id"`
	CPF        string  `json:"cpf"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Address    Address `json:"addres"`
}

func NewClientJson(c Client) []byte {
	res, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(res))
	return res

}
