package repository

import "account_service/internal"

type Storage interface {
	CreateAccount(*internal.CreateAccountRequest) (int, error)
	GetAccounts() ([]internal.Account, error)
	GetAccountById(int) (*internal.Account, error)
	DeleteAccount(int) error
	UpdateAccount(*internal.Account) (*internal.Account, error)
}
