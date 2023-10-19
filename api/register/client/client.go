package client

import (
	"encoding/json"
	"log"
)

type Address struct {
	Street string `json:"street"`
	Number int    `json:"number"`
}

type Client struct {
	Account_id int     `json:"account_id"`
	CPF        string  `json:"cpf"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Address    Address `json:"addres"`
}

func NewClient(c Client) string {
	res, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(res))
	return string(res)

}