package customermodel

import (
	"transaction/balance"
	"transaction/transactionmodel"
)

type Customer struct {
	ID             int    `json:"id"`
	Account_number int    `json:"account_number"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Phone_number   string `json:"phone_number"`
	Email          string `json:"email"`
	Transaction    transactionmodel.Transaction
	Balance        balance.Balance
}
