package postgres

import (
	"account_service/internal"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type PostgresStorage struct {
	db        *sql.DB
	tableName string
}

func NewPostgresStorage(table string) (*PostgresStorage, error) {
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
	return &PostgresStorage{
		db:        postgresDB,
		tableName: table,
	}, nil
}

func (p *PostgresStorage) CreateAccount(account *internal.CreateAccountRequest) (int, error) {
	queryInsertAccount := fmt.Sprintf(
		"insert into %s (username, password, first_name, last_name, email) values ($1, $2, $3, $4, $5) returning id;", p.tableName)
	log.Println(queryInsertAccount)
	var accountId int
	err := p.db.QueryRow(queryInsertAccount, account.Username, account.Password, account.FirstName, account.LastName, account.Email).Scan(&accountId)
	if err != nil {
		return -1, err
	}
	return accountId, nil
}

func (p *PostgresStorage) GetAccounts() ([]internal.Account, error) {
	selectAllQuery := fmt.Sprintf("SELECT * FROM %s;", p.tableName)
	log.Println(selectAllQuery)
	rows, err := p.db.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}

	var accounts []internal.Account
	for rows.Next() {
		account := new(internal.Account)
		err := rows.Scan(
			&account.Id,
			&account.Username,
			&account.Password,
			&account.FirstName,
			&account.LastName,
			&account.Email,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, *account)
	}
	return accounts, nil
}

func (p *PostgresStorage) GetAccountById(accountId int) (*internal.Account, error) {
	var account internal.Account
	selectById := fmt.Sprintf("SELECT * FROM %s WHERE id=%d;", p.tableName, accountId)
	err := p.db.QueryRow(selectById).Scan(
		&account.Id, &account.Username, &account.Password, &account.FirstName, &account.LastName, &account.Email)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (p *PostgresStorage) UpdateAccount(account *internal.Account) (*internal.Account, error) {
	updateQuery := fmt.Sprintf("UPDATE %s SET username=$2, password=$3, first_name=$4, last_name=$5, email=$6 WHERE id=$1;", p.tableName)
	_, err := p.db.Exec(updateQuery, account.Id, account.Username, account.Password, account.FirstName, account.LastName, account.Email)
	if err != nil {
		return nil, err
	}
	return p.GetAccountById(account.Id)
}

func (p *PostgresStorage) DeleteAccount(id int) (int, error) {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", p.tableName)
	_, err := p.db.Exec(deleteQuery, id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
