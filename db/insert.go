package db

import (
	"encoding/json"
	"fmt"
	"goproject/api/register/models"
	"log"
)

var userList []string

func Insert(c models.ClientResponse) []string {

	jsonBytes, err := json.Marshal(&c)
	if err != nil {
		log.Fatalf("Error marshaling struct %v", err)
	}

	jsonString := string(jsonBytes)

	userList = append(userList, jsonString)
	fmt.Println("lista de users", userList)

	return userList
}
