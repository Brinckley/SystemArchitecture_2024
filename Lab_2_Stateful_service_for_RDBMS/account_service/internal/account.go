package internal

import "account_service/internal/util"

type CreateAccountRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Account struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func AccountFrom(id int, request *CreateAccountRequest) *Account {
	return &Account{
		Id:        id,
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	}
}

func (a *CreateAccountRequest) SetPassword(password string) error {
	var err error
	a.Password, err = util.HashPassword(password)
	if err != nil {
		return err
	}
	return nil
}
