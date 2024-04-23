package repository

import "account_service/internal"

type Storage interface {
	SignUpAccount(account internal.SignUpAccount) (int, error)
	GetPasswordByUsername(username string) (int, string, error)

	GetAllAccounts() ([]internal.Account, error)
	GetAccountById(id int) (internal.Account, error)
	DeleteById(id int) error
	UpdateById(accountUpd internal.Account) error
	GetAccountsByMask(search internal.AccountSearch) ([]internal.Account, error)
}
