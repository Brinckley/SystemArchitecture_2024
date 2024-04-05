package storage

import (
	"context"
	"message_service/internal"
)

type Storage interface {
	Create(ctx context.Context, msg internal.MessageDto) (string, error)
	GetById(ctx context.Context, id string) (internal.Message, error)
	GetByDestId(ctx context.Context, id string) ([]internal.Message, error)
}
