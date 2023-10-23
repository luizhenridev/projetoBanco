package deposits

import (
	"encoding/json"
	"goproject/api/register/models"
	"goproject/db"
	"log"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	var request models.Deposit

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	if request.Value == 0 || request.Value < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_AMOUNT",
			ErrorMessage: "O valor do depósito não deve ser menor ou igual a zero",
		}
		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	} else if request.Currency != "BRL" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_CURRENCY",
			ErrorMessage: "A moeda selecionada não é suportada.",
		}

		err = json.NewEncoder(w).Encode(response) // Defina o código de status antes de escrever o corpo
		if err != nil {
			log.Println(err)
		}
		return
	}

	response := models.DepositResponse{
		Account_id:      request.Account_id,
		TransactionType: request.TransactionType,
		Value:           request.Value,
		Currency:        request.Currency,
		TransactionDate: request.TransactionDate,
		Description:     request.Description,
		Balance:         request.Value,
		Status:          "Success",
		Message:         "The transaction was successful.",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response) // Defina o código de status antes de escrever o corpo
	if err != nil {
		log.Println(err)
	}

	db.InsertDeposit(response)

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
