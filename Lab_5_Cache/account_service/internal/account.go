package internal

type Account struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type SignUpAccount struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type AccountSearch struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SignInAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AccountFrom(id int, request *SignUpAccount) *Account {
	return &Account{
		Id:        id,
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	}
}
