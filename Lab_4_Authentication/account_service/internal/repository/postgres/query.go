package postgres

import (
	"account_service/internal"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func (p *Storage) SignUpAccount(account internal.SignUpAccount) (accountId int, err error) {
	queryInsertAccount := fmt.Sprintf(
		"insert into %s (username, password, first_name, last_name, email) values ($1, $2, $3, $4, $5) returning id;", p.tableName)
	log.Println(queryInsertAccount)
	err = p.db.QueryRow(queryInsertAccount, account.Username, account.Password,
		account.FirstName, account.LastName, account.Email).Scan(&accountId)
	if err != nil {
		return -1, err
	}
	return accountId, nil
}

func (p *Storage) GetPasswordByUsername(username string) (password string, err error) {
	getPasswordQuery := fmt.Sprintf("SELECT password FROM %s WHERE username=$1;", p.tableName)
	err = p.db.QueryRow(getPasswordQuery, username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (p *Storage) GetAllAccounts() (accounts []internal.Account, err error) {
	selectAllQuery := fmt.Sprintf("SELECT * FROM %s;", p.tableName)
	rows, err := p.db.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}

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

func (p *Storage) DeleteById(id int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", p.tableName)
	_, err := p.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Storage) UpdateById(accountUpd internal.Account) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET username=$2, password=$3, first_name=$4, last_name=$5, email=$6 WHERE id=$1;", p.tableName)
	_, err := p.db.Exec(updateQuery, accountUpd.Id, accountUpd.Username, accountUpd.Password, accountUpd.FirstName, accountUpd.LastName, accountUpd.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *Storage) GetAccountById(accountId int) (account internal.Account, err error) {
	selectById := fmt.Sprintf("SELECT * FROM %s WHERE id=%d;", p.tableName, accountId)
	err = p.db.QueryRow(selectById).Scan(
		&account.Id, &account.Username, &account.Password, &account.FirstName, &account.LastName, &account.Email)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (p *Storage) GetAccountsByMask(search internal.AccountSearch) (accounts []internal.Account, err error) {
	likeQuery := fmt.Sprintf("SELECT * FROM %s WHERE first_name LIKE '%s' AND last_name LIKE '%s';",
		p.tableName, search.FirstName, search.LastName)
	rows, err := p.db.Query(likeQuery)
	if err != nil {
		return nil, err
	}

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
