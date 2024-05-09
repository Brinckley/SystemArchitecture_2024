package storage

import "Gateway/internal/entity"

type Cache interface {
	Get(key string) (entity.Account, error)
	Set(key string, value entity.Account) error
}
