package internal

type AccountDto struct {
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
}

type Account struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
}

type AccountSearch struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}
