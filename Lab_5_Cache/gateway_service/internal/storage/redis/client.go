package cache

import (
	"Gateway/internal/entity"
	"Gateway/internal/storage"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisCache struct {
	ctx    context.Context
	ttl    time.Duration
	client *redis.Client
}

func NewRedisClient(host, port, password string, db int, ttl time.Duration, ctx context.Context) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})
	if pong := client.Ping(ctx); pong.String() != "ping: PONG" {
		log.Println("-------------Error connection redis ----------:", pong)
	} else {
		log.Println("-------------CONNECTED TO REDIS ------------- : ", pong)
	}

	return &RedisCache{
		ctx:    ctx,
		ttl:    ttl,
		client: client,
	}
}

func (r *RedisCache) Get(key string) (account entity.Account, err error) {
	err = r.client.Get(r.ctx, key).Scan(&account)
	if err != nil {
		return account, storage.NewCacheError(fmt.Sprintf("cannot find account error : %s", err))
	}
	return account, nil
}

func (r *RedisCache) Set(key string, account entity.Account) error {
	statusCmd := r.client.Set(r.ctx, "accountId:"+key, account, r.ttl)
	return statusCmd.Err()
}
