package main

import (
	"database/sql"
	"log"
	"net/http"
	"transaction/connection"
	"transaction/customercontroller"
	"transaction/transactioncontroller"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db := connection.ConnectDB()
	customer := customercontroller.Controller{}
	transaction := transactioncontroller.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/signup", customer.Createcustomer(db)).Methods("POST")
	router.HandleFunc("/delete", customer.Deletecustomer(db)).Methods("DELETE")
	router.HandleFunc("/update", customer.ChangeName(db)).Methods("PUT")
	router.HandleFunc("/changeName", customer.ChangeName(db)).Methods("PUT")

	router.HandleFunc("/transfer", transaction.Transfer(db)).Methods("POST")
	router.HandleFunc("/accountInfo/{id}", transaction.CustomersTransactions(db)).Methods("GET")
	router.HandleFunc("/withdraw", transaction.Withdraw(db)).Methods("POST")
	router.HandleFunc("/checkBalance", transaction.Balance(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
