package internal

type MessageDto struct {
	SenderId   int    `json:"sender_id" bson:"sender_id"`
	ReceiverId int    `json:"receiver_id" bson:"receiver_id"`
	Content    string `json:"content" bson:"content"`
}

type Message struct {
	Id         int    `json:"id" bson:"_id,omitempty"`
	SenderId   int    `json:"sender_id" bson:"sender_id"`
	ReceiverId int    `json:"receiver_id" bson:"receiver_id"`
	Content    string `json:"content" bson:"content"`
}
