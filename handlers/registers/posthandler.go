package registers

import (
	"encoding/json"
	"fmt"
	"goproject/api/register/models"
	"goproject/db"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	var request models.Client

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	matchCPF, _ := regexp.MatchString(`^([\d]{3})([.]?)([\d]{3})([.]?)([\d]{3})([.|-]?)([\d]{2})$`, request.CPF)
	if !matchCPF {
		fmt.Println("O cpf ÉSTÁ errado", matchCPF)
	}

	matchEmail, _ := regexp.MatchString(`(\w{1,}[.]?)\w{1,}[@]\w{1,}[.]\w{3}([.][a-z]+)?`, request.Email)

	if !matchCPF && request.Name == "" && !matchEmail {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_PARAMETERS",
			ErrorMessage: "Parâmetro parâmetros ausentes",
		}
		err = json.NewEncoder(w).Encode(&response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	} else if !matchCPF {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_CPF",
			ErrorMessage: "Parâmetro CPF ausente ou mal formatado",
		}
		err = json.NewEncoder(w).Encode(&response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			return
		}
		return
	} else if request.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_NAME",
			ErrorMessage: "Parâmetro NAME ausente",
		}
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	} else if !matchEmail {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_EMAIL",
			ErrorMessage: "Parâmetro EMAIL ausente OU mal formatado",
		}
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	}

	response := models.ClientResponse{
		Account_id: uuid.New().String(),
		CPF:        request.CPF,
		Name:       request.Name,
		Email:      request.Email,
		Address:    models.Address{Street: request.Address.Street, Number: request.Address.Number}}

	validate := db.ValidateCPF(response)

	if validate == true {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}

		db.Insert(response)
	} else if validate == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_CPF",
			ErrorMessage: "CPF já cadastrado em nossa base de dados",
		}
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return

	}

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line stdout
		log.Println("Request Processada com sucesso")
		//Imprime no browser
		//w.Write([]byte("Request Processada com sucesso"))
		//w.Write([]byte(response))
	case <-ctx.Done():
		//imprime no comand line stdout
		log.Println("Request cancelada pelo client")

	}

	defer log.Println("Request Finalizada")

}
