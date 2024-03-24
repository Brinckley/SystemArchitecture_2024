package internal

type CreateMessageRequest struct {
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Content    string `json:"content"`
}

type Message struct {
	Id         int    `json:"id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Content    string `json:"content"`
}

func MessageFrom(id int, request *CreateMessageRequest) *Message {
	return &Message{
		Id:         id,
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Content:    request.Content,
	}
}
