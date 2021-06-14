package transactionmodel

type Transaction struct {
	ID          int    `json:"id"`
	Customer_id int    `json:"customer_id"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
}
