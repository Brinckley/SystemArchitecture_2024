package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisCache struct {
	ctx    context.Context
	ttl    time.Duration
	cbTtl  time.Duration
	client *redis.Client
}

func NewRedisClient(host, port, password string, db int, ttl, cbTtl time.Duration, ctx context.Context) *RedisCache {
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
		cbTtl:  cbTtl,
		client: client,
	}
}
