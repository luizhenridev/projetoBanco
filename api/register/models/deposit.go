package models

type Deposit struct {
	Account_id      string `json:"account_id"`
	TransactionType string `json:"transactiontype"`
	Value           int    `json:"value"`
	Currency        string `json:"currency"`
	TransactionDate string `json:"transactiondate"`
	Description     string `json:"description"`
}

type DepositResponse struct {
	Account_id      string `json:"account_id"`
	TransactionType string `json:"transactiontype"`
	Value           int    `json:"value"`
	Currency        string `json:"currency"`
	TransactionDate string `json:"transactiondate"`
	Description     string `json:"description"`
	Balance         int    `json:"balance"`
	Status          string `json:"status"`
	Message         string `json:"message"`
}
