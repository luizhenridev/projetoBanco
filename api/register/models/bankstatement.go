package models

type BankStatementResponse struct {
	Account_id string           `json:"account_id"`
	Name       string           `json:"name"`
	Email      string           `json:"email"`
	Transfers  TransferResponse `json:"transfers"`
	Deposits   DepositResponse  `json:"deposits"`
	Withdraw   WithdrawResponse `json:"withdraws"`
}
