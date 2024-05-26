package cache

import (
	"Gateway/internal/entity"
	"Gateway/internal/storage"
	"fmt"
)

const ACCOUNT_REDIS_PREFIX = "accountId:"

func (r *RedisCache) GetAccount(key string) (account entity.Account, err error) {
	err = r.client.Get(r.ctx, ACCOUNT_REDIS_PREFIX+key).Scan(&account)
	if err != nil {
		return account, storage.NewCacheError(fmt.Sprintf("cannot find account error : %s", err))
	}
	return account, nil
}

func (r *RedisCache) SetAccount(key string, account entity.Account) error {
	statusCmd := r.client.Set(r.ctx, ACCOUNT_REDIS_PREFIX+key, account, r.ttl)
	return statusCmd.Err()
}

func (r *RedisCache) DeleteAccount(key string) error {
	statusCmd := r.client.Del(r.ctx, ACCOUNT_REDIS_PREFIX+key)
	return statusCmd.Err()
}
