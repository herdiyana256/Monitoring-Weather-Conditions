package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertData(db *sql.DB, water, wind float64, status string) error {
	query := "INSERT INTO ass3 (water, wind, status) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, water, wind, status)
	if err != nil {
		return err
	}

	return nil
}