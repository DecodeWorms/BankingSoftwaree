package transactioncontroller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transaction/balance"
	"transaction/customermodel"
	"transaction/transactionmodel"

	"github.com/gorilla/mux"
)

var db *sql.DB

type Controller struct{}

func (control Controller) CustomersTransactions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theTransaction transactionmodel.Transaction

		//var transactionId int

		var params = mux.Vars(r)

		err := db.QueryRow("select * from transsactions where customer_id = $1", params["id"], &theTransaction.ID, &theTransaction.Customer_id, &theTransaction.Amount, &theTransaction.Date)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(theTransaction)
	}
}

func (control Controller) Transfer(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var theTransaction transactionmodel.Transaction
		json.NewDecoder(r.Body).Decode(&theTransaction)

		var theBalance balance.Balance

		var tId int
		var bId int

		err := db.QueryRow("insert into transactions(customer_id,amount,date) values($1,$2,$3) RETURNING customer_id",
			theTransaction.Customer_id, theTransaction.Amount, theTransaction.Date).Scan(&tId)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(tId)

		//cheking if this is first transaction from the user

		row := db.QueryRow("select transaction_id ,balance from balance where transaction_id = $1 ", theTransaction.Customer_id)

		row.Scan(&theBalance.Transaction_id, &theBalance.Balance)

		if theBalance.Transaction_id == 0 {
			fmt.Println("am empty")
			err := db.QueryRow("insert into balance(transaction_id,balance,date) values($1,$2,$3) RETURNING transaction_id", theTransaction.Customer_id, theTransaction.Amount, theTransaction.Date).Scan(&bId)

			if err != nil {
				log.Fatal(err)
			}

		} else {
			fmt.Println(theBalance.Transaction_id, theTransaction.Amount, theBalance.Balance)

			number, err := db.Exec("update balance set balance = $1, date = $2 where transaction_id = $3", theBalance.Balance+theTransaction.Amount, theTransaction.Date, theBalance.Transaction_id)

			if err != nil {
				log.Fatal(err)
			}
			num, err := number.RowsAffected()

			if err != nil {
				log.Fatal(err)
			}

			json.NewEncoder(w).Encode(num)

		}

	}
}

func (control Controller) Withdraw(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var theTransaction transactionmodel.Transaction
		json.NewDecoder(r.Body).Decode(&theTransaction)
		var theBalance balance.Balance

		var customerId int

		err := db.QueryRow("insert into transactions(customer_id,amount,date) values($1,$2,$3) RETURNING customer_id", theTransaction.Customer_id, theTransaction.Amount, theTransaction.Date).Scan(&customerId)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(customerId)

		row := db.QueryRow("select transaction_id ,balance from balance where transaction_id = $1", theTransaction.Customer_id)

		row.Scan(&theBalance.Transaction_id, &theBalance.Balance)

		if theBalance.Balance != 0 && theBalance.Balance > theTransaction.Amount {
			// fmt.Println(theBalance.Balance,theTransaction.Amount)
			num, err := db.Exec("update balance set balance = $1, date = $2 where transaction_id = $3", theBalance.Balance-theTransaction.Amount, theTransaction.Date, theBalance.Transaction_id)

			if err != nil {
				log.Fatal(err)
			}

			numRow, err := num.RowsAffected()

			if err != nil {
				log.Fatal(err)
			}

			json.NewEncoder(w).Encode(numRow)
		} else {
			fmt.Printf("The account number %d is having  insufficient balance %d and the amount requested : %d", theBalance.Transaction_id, theBalance.Balance, theTransaction.Amount)
		}

	}
}

func (control Controller) Balance(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var theCustomer customermodel.Customer
		json.NewDecoder(r.Body).Decode(&theCustomer)

		// var accountNumber = mux.Vars(r)

		var theCustomerResult customermodel.Customer

		row := db.QueryRow("select account_number,first_name,last_name,balance from customers inner join balance on account_number = transaction_id where account_number = $1 ", theCustomer.Account_number)

		row.Scan(&theCustomerResult.Account_number, &theCustomerResult.First_name, &theCustomerResult.Last_name, &theCustomerResult.Balance.Balance)

		json.NewEncoder(w).Encode(theCustomerResult)

	}
}
