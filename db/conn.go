package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbName   = "postgres"
)

func ConectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(" host=%s port=%d user=%s "+
		" password=%s dbname=%s sslmode=disable ",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to" + dbName + " database")
	return db, nil
}
