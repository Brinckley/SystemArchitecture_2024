package repository

import "message_service/internal"

type Storage interface {
	CreateMessage(*internal.CreateMessageRequest) (int, error)
	GetMessagesByReceiverId(int) ([]internal.Message, error)
	GetMessageById(int, int) (*internal.Message, error)
}
