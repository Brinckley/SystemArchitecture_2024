package cache

import (
	"Gateway/internal/entity"
	"Gateway/internal/storage"
	"fmt"
)

const (
	MESSAGE_REDIS_PREFIX               = "messageId:"
	MESSAGES_WITH_ACCOUNT_REDIS_PREFIX = "accountWithMessages:"
)

func (r *RedisCache) GetMessage(key string) (message entity.Message, err error) {
	err = r.client.Get(r.ctx, MESSAGE_REDIS_PREFIX+key).Scan(&message)
	if err != nil {
		return message, storage.NewCacheError(fmt.Sprintf("cannot find message error : %s", err))
	}
	return message, nil
}

func (r *RedisCache) SetMessage(key string, message entity.Message) error {
	statusCmd := r.client.Set(r.ctx, MESSAGE_REDIS_PREFIX+key, message, r.ttl)
	return statusCmd.Err()
}

func (r *RedisCache) GetMessageCollection(accountId string) (messages entity.MessagesCollection, err error) {
	err = r.client.Get(r.ctx, MESSAGES_WITH_ACCOUNT_REDIS_PREFIX+accountId).Scan(&messages)
	if err != nil {
		return messages, storage.NewCacheError(fmt.Sprintf("cannot find messages for accountId %s error : %s", accountId, err))
	}
	return messages, nil
}

func (r *RedisCache) SetMessageCollection(accountId string, messages entity.MessagesCollection) error {
	statusCmd := r.client.Set(r.ctx, MESSAGES_WITH_ACCOUNT_REDIS_PREFIX+accountId, messages, r.ttl)
	return statusCmd.Err()
}
