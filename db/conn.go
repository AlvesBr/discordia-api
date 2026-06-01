package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectDB() (*sql.DB, error) {
	var dsn string

	if url := os.Getenv("DATABASE_URL"); url != "" {
		dsn = url
	} else {
		dsn = fmt.Sprintf(
			"host=localhost port=5432 user=postgres password=1234 dbname=api-go sslmode=disable",
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco: %w", err)
	}

	fmt.Println("Connected to database")
	return db, nil
}
