package storage

import (
	"account_service/internal"
	"context"
)

type Storage interface {
	Create(ctx context.Context, account internal.AccountDto) (string, error)
	GetAll(ctx context.Context) ([]internal.Account, error)
	GetById(ctx context.Context, id string) (internal.Account, error)
	Update(ctx context.Context, account internal.Account) error
	Delete(ctx context.Context, id string) error
	GetByMask(ctx context.Context, regex internal.AccountSearch) ([]internal.Account, error)
}
