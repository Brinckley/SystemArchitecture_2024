package storage

import (
	"account_service/internal"
	"context"
)

type Storage interface {
	Create(ctx context.Context, account internal.Account) (string, error)
	GetAll(ctx context.Context) ([]internal.Account, error)
	GetById(ctx context.Context, id string) (internal.Account, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, account internal.Account) error
}
