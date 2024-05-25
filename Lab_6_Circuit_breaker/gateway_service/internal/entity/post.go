package entity

import "encoding/json"

type Post struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	AccountId string `json:"account_id" bson:"account_id"`
	Content   string `json:"content" bson:"content"`
}

func (a Post) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *Post) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

type PostCollection struct {
	Posts map[string]Post
}
