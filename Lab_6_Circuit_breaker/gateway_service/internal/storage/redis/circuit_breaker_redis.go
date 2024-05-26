package cache

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

const CIRCUIT_BREAKER_PREFIX = "circuit_breaker:"

func (r *RedisCache) IncrementCounter(serviceName string) (int, error) {
	counter, err := r.GetCounter(serviceName)
	if err != nil {
		return 0, err
	}
	counter++
	err = r.client.Set(r.ctx, CIRCUIT_BREAKER_PREFIX+serviceName, counter, r.cbTtl).Err()
	if err != nil {
		return -1, err
	}
	return counter, nil
}

func (r *RedisCache) GetCounter(serviceName string) (int, error) {
	counterString, err := r.client.Get(r.ctx, CIRCUIT_BREAKER_PREFIX+serviceName).Result()
	if errors.Is(err, redis.Nil) {
		counterString = "0"
	} else if err != nil {
		return -1, fmt.Errorf("cannot get counter from cache err %s", err)
	}
	counterInteger, err := strconv.Atoi(counterString)
	if err != nil {
		return -1, fmt.Errorf("cannot convert to int err %s", err)
	}
	return counterInteger, nil
}

func (r *RedisCache) ClearCounter(serviceName string) error {
	err := r.client.Del(r.ctx, CIRCUIT_BREAKER_PREFIX+serviceName).Err()
	if errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}
