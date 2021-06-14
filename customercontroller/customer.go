package customercontroller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"transaction/customermodel"
)

type Controller struct{}

var customers []customermodel.Customer

var db *sql.DB

func (control Controller) Createcustomer(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var theCustomer customermodel.Customer

		var customerId int

		json.NewDecoder(r.Body).Decode(&theCustomer)

		err := db.QueryRow("insert into customers(account_number,first_name,last_name,email,phone_number) values($1,$2,$3,$4,$5) RETURNING account_number", theCustomer.Account_number, theCustomer.First_name, theCustomer.Last_name, theCustomer.Phone_number, theCustomer.Email).Scan(&customerId)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(customerId)

	}
}

func (control Controller) Deletecustomer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var theCustomer customermodel.Customer
		json.NewDecoder(r.Body).Decode(&theCustomer)

		number, err := db.Exec("delete from customers where account_number = $1", theCustomer.Account_number)

		if err != nil {
			log.Fatal(err)

		}
		rowsDeleted, err := number.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}

		number2, err2 := db.Exec("delete from transactions where customer_id = $1", theCustomer.Account_number)

		if err2 != nil {
			log.Fatal(err2)
		}

		rowsDeleted2, theErr := number2.RowsAffected()

		if theErr != nil {
			log.Fatal(theErr)
		}

		number3, err3 := db.Exec("delete from balance where transaction_id = $1", theCustomer.Account_number)

		if err3 != nil {
			log.Fatal(err3)
		}

		rowsDeleted3, theErr2 := number3.RowsAffected()

		if theErr2 != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(rowsDeleted)
		json.NewEncoder(w).Encode(rowsDeleted2)
		json.NewEncoder(w).Encode(rowsDeleted3)
	}
}

func (control Controller) ChangeName(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var theCustomer customermodel.Customer
		json.NewDecoder(r.Body).Decode(&theCustomer)

		row, err := db.Exec("update customers set first_name = $1, last_name = $2 where account_number = $3", theCustomer.First_name, theCustomer.Last_name, theCustomer.Account_number)

		if err != nil {
			log.Fatal(err)
		}

		rowsUpdated, err := row.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(rowsUpdated)

	}
}
