package entity

import (
	"encoding/json"
)

type Account struct {
	Id        int    `json:"id" redis:"id"`
	Username  string `json:"username" redis:"username"`
	Password  string `json:"password" redis:"password"`
	FirstName string `json:"first_name" redis:"first_name"`
	LastName  string `json:"last_name" redis:"last_name"`
	Email     string `json:"email" redis:"email"`
}

func (a Account) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *Account) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

type SignUpAccount struct {
	Username  string `json:"username" redis:"username"`
	Password  string `json:"password" redis:"password"`
	FirstName string `json:"first_name" redis:"first_name"`
	LastName  string `json:"last_name" redis:"last_name"`
	Email     string `json:"email" redis:"email"`
}

func (a SignUpAccount) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *SignUpAccount) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

type AccountCollection struct {
	Accounts map[string]Account
}
