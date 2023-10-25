package db

import (
	"encoding/json"
	"goproject/api/register/models"
	"net/http"
)

var userDeposits []models.DepositResponse

func SumBalance(d *models.DepositResponse) models.DepositResponse {

	if userDeposits != nil {
		lastUserDeposits := userDeposits[len(userDeposits)-1]
		if lastUserDeposits.Account_id == d.Account_id {
			d.Balance = lastUserDeposits.Balance + d.Value
			return *d
		}
	}

	return *d

}

func InsertDeposit(d models.DepositResponse) *[]models.DepositResponse {
	userDeposits = append(userDeposits, d)
	return &userDeposits

}

func GetInsertDeposit() *[]models.DepositResponse {
	return &userDeposits
}

func GetInsertIDDeposit(params map[string]string) *[]models.DepositResponse {
	var userDepositsID []models.DepositResponse
	for _, deposit := range userDeposits {
		if deposit.Account_id == params["id"] {
			userDepositsID = append(userDepositsID, deposit)
			continue
		}

		return &userDepositsID
	}
	return &userDepositsID
}

func UpdateDeposit(params map[string]string, r *http.Request) *models.DepositResponse {
	for index, deposit := range userDeposits {
		if deposit.Account_id == params["id"] {
			userDeposits = append(userDeposits[:index], userDeposits[index+1:]...)

			var userDeposit models.DepositResponse
			_ = json.NewDecoder(r.Body).Decode(&userDeposit)
			userDeposit.Account_id = params["id"]
			userDeposits = append(userDeposits, userDeposit)
			return &userDeposit
		}
	}
	return nil
}

func DeleteDeposit(params map[string]string, r *http.Request) *[]models.DepositResponse {
	for _, deposit := range userDeposits {
		if deposit.Account_id == params["id"] {
			userDeposits = append(userDeposits[:], userDeposits[:]...)
			return &userDeposits
		}
	}
	return nil
}
