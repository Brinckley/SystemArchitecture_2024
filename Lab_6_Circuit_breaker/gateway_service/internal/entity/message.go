package entity

import "encoding/json"

type Message struct {
	Id         string `json:"id" bson:"_id,omitempty"`
	SenderId   string `json:"sender_id" bson:"sender_id"`
	ReceiverId string `json:"receiver_id" bson:"receiver_id"`
	Content    string `json:"content" bson:"content"`
}

func (a Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

type MessagesCollection struct {
	Messages map[string]Message
}
