package internal

type MessageDto struct {
	SenderId   string `json:"sender_id" bson:"sender_id"`
	ReceiverId string `json:"receiver_id" bson:"receiver_id"`
	Content    string `json:"content" bson:"content"`
}

type Message struct {
	Id         string `json:"id" bson:"_id,omitempty"`
	SenderId   string `json:"sender_id" bson:"sender_id"`
	ReceiverId string `json:"receiver_id" bson:"receiver_id"`
	Content    string `json:"content" bson:"content"`
}
