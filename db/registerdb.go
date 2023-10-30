package db

import (
	"database/sql"
	"encoding/json"
	"goproject/api/register/models"
	"log"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func ValidateCPF(c models.ClientResponse) bool {
	//Database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("select cpf from bank where cpf = ?")

	var v models.ClientResponse
	err = stmt.QueryRow(c.CPF).Scan(&v.CPF)
	if err != nil {
		log.Println(err)
	}
	if v.CPF == c.CPF {
		return false
	}

	defer stmt.Close()
	return true
}

func Insert(c models.ClientResponse) (*models.ClientResponse, error) {
	//Database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into bank(account_id, cpf, name, email,transactiontype, value, currency, transactiondate, description, balance, status, message) values (?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(c.Account_id, c.CPF, c.Name, c.Email, "Register", 0, "BRL", "-", "register", 0, "Success", "The transaction was successful.")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return &c, nil
}

func GetInsert() (*[]models.ClientResponse, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select account_id, name, cpf, email from bank")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sc []models.ClientResponse
	for rows.Next() {
		var c models.ClientResponse
		err = rows.Scan(&c.Account_id, &c.Name, &c.CPF, &c.Email)
		if err != nil {
			return nil, err
		}
		sc = append(sc, c)
	}
	return &sc, nil
}

func GetInsertID(params map[string]string) (*models.ClientResponse, error) {
	id := params["id"]
	//Database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select account_id, name, cpf, email from bank where account_id = ?")
	if err != nil {
		log.Println(err)
		return nil, err

	}
	var v models.ClientResponse
	err = stmt.QueryRow(id).Scan(&v.Account_id, &v.Name, &v.CPF, &v.Email)
	if err != nil {
		log.Println(err)
		return nil, err

	}
	defer stmt.Close()
	return &v, nil
}

func Update(params map[string]string, w http.ResponseWriter, r *http.Request) (any, error) {
	id := params["id"]
	var request models.ClientResponse

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	matchCPF, _ := regexp.MatchString(`^([\d]{3})([.]?)([\d]{3})([.]?)([\d]{3})([.|-]?)([\d]{2})$`, request.CPF)

	matchEmail, _ := regexp.MatchString(`(\w{1,}[.]?)\w{1,}[@]\w{1,}[.]\w{3}([.][a-z]+)?`, request.Email)

	if !matchCPF && request.Name == "" && !matchEmail {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_PARAMETERS",
			ErrorMessage: "Parâmetro parâmetros ausentes",
		}

		return &response, nil
	} else if !matchCPF {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_CPF",
			ErrorMessage: "Parâmetro CPF ausente ou mal formatado",
		}
		if err != nil {
			log.Println(err)
		}
		return &response, nil
	} else if request.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_NAME",
			ErrorMessage: "Parâmetro NAME ausente",
		}
		if err != nil {
			log.Println(err)
		}
		return &response, nil
	} else if !matchEmail {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.Erros{
			ErrorCode:    "INVALID_EMAIL",
			ErrorMessage: "Parâmetro EMAIL ausente OU mal formatado",
		}
		if err != nil {
			log.Println(err)
		}
		return &response, nil
	}

	//Database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update bank set cpf = ?, name = ?, email = ? where account_id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(request.CPF, request.Name, request.Email, id)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return &request, nil
}

func Delete(params map[string]string) error {
	id := params["id"]

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from bank where account_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil

}
