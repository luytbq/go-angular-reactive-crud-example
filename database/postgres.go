package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/luytbq/go-angular-reactive-crud-example/config"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	// connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.App.PG_USER,
		config.App.PG_PASSWORD,
		config.App.PG_HOST,
		config.App.PG_PORT,
		config.App.PG_DB_NAME,
		config.App.PG_SSL_MODE,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = initPostgres(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initPostgres(db *sql.DB) error {
	err := db.Ping()

	if err != nil {
		return err
	}

	log.Println("Postgres DB init successfully")

	return nil
}
