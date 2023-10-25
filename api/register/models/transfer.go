package models

type Transfer struct {
	Account_id_from string `json:"account_id_from"`
	Account_id_to   string `json:"account_id_to"`
	TransactionType string `json:"transactiontype"`
	Value           int    `json:"value"`
	Currency        string `json:"currency"`
	TransactionDate string `json:"transactiondate"`
	Description     string `json:"description"`
}

type TransferResponse struct {
	Account_id_from string `json:"account_id_from"`
	Account_id_to   string `json:"account_id_to"`
	Name_to         string `json:"name_to"`
	Email_to        string `json:"email_to"`
	TransactionType string `json:"transactiontype"`
	Value           int    `json:"value"`
	Currency        string `json:"currency"`
	TransactionDate string `json:"transactiondate"`
	Description     string `json:"description"`
	Balance         int    `json:"balance"`
	Status          string `json:"status"`
	Message         string `json:"message"`
}
