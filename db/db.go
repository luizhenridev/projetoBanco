package db

import (
	"encoding/json"
	"goproject/api/register/models"
	"net/http"
)

var userPosts []models.ClientResponse

func Insert(c models.ClientResponse) *[]models.ClientResponse {
	userPosts = append(userPosts, c)
	return &userPosts
}

func GetInsert() *[]models.ClientResponse {
	return &userPosts
}

func GetInsertID(params map[string]string) *models.ClientResponse {
	for _, item := range userPosts {
		if item.Account_id == params["id"] {
			return &item
		}
	}
	return nil
}

func Update(params map[string]string, r *http.Request) *models.ClientResponse {
	for index, item := range userPosts {
		if item.Account_id == params["id"] {
			userPosts = append(userPosts[:index], userPosts[index+1:]...)

			var userPost models.ClientResponse
			_ = json.NewDecoder(r.Body).Decode(&userPost)
			userPost.Account_id = params["id"]
			userPosts = append(userPosts, userPost)
			return &userPost
		}
	}
	return nil
}

func Delete(params map[string]string, r *http.Request) *[]models.ClientResponse {
	for _, item := range userPosts {
		if item.Account_id == params["id"] {
			userPosts = append(userPosts[:], userPosts[:]...)
			return &userPosts
		}
	}
	return nil
}
