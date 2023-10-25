package models

type BalanceResponse struct {
	Account_id_from string `json:"account_id_from"`
	Name_to         string `json:"name_to"`
	Email_to        string `json:"email_to"`
	TransactionType string `json:"transactiontype"`
	Currency        string `json:"currency"`
	TransactionDate string `json:"transactiondate"`
	Description     string `json:"description"`
	Balance         int    `json:"balance"`
	Status          string `json:"status"`
	Message         string `json:"message"`
}
