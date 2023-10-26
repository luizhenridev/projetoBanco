package registers

import (
	"encoding/json"
	"goproject/db"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	//v := reflect.ValueOf(&userPosts)
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")

	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	params := mux.Vars(r)

	userPost, err := db.Update(params, r)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&userPost) // Defina o código de status antes de escrever o corpo
	if err != nil {
		log.Println(err)
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
}
