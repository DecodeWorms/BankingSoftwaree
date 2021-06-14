package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)

	if err != nil {
		logFatal(err)
	}

	if err = db.Ping(); err != nil {
		logFatal(err)
	} else {
		fmt.Println("connected")
	}

	return db

}
