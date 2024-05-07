package postgres

import (
	"database/sql"
	"fmt"
	"os"
)

type Storage struct {
	db        *sql.DB
	tableName string
}

func NewPostgresStorage(table string) (*Storage, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password='%s' sslmode=disable search_path=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SCHEMA"))
	postgresDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := postgresDB.Ping(); err != nil {
		return nil, err
	}
	return &Storage{
		db:        postgresDB,
		tableName: table,
	}, nil
}
