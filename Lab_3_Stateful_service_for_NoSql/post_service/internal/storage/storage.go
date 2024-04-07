package storage

import (
	"context"
	"post_service/internal"
)

type Storage interface {
	Create(ctx context.Context, postDto internal.PostDto) (string, error)
	GetById(ctx context.Context, hexId string) (internal.Post, error)
	GetByAccountId(ctx context.Context, hexId string) ([]internal.Post, error)
	Update(ctx context.Context, post internal.Post) error
	Delete(ctx context.Context, hexId string) error
}