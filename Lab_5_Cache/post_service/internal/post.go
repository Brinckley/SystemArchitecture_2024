package internal

type PostDto struct {
	AccountId string `json:"account_id" bson:"account_id"`
	Content   string `json:"content" bson:"content"`
}

type Post struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	AccountId string `json:"account_id" bson:"account_id"`
	Content   string `json:"content" bson:"content"`
}
