package balance

type Balance struct {
	ID             int    `json:"id"`
	Transaction_id int    `json:"transaction_id"`
	Balance        int    `json:"balance"`
	Date           string `json:"date"`
}
