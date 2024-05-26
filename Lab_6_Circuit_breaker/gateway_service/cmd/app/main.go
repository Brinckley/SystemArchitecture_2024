package main

import (
	"Gateway/internal/service/circuit_breaker"
	"Gateway/internal/service/router"
	cache "Gateway/internal/storage/redis"
	"context"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()
	cacheHost := os.Getenv("CACHE_HOST")
	cachePort := os.Getenv("CACHE_PORT")
	cachePassword := os.Getenv("CACHE_PASSWORD")
	cacheDb := os.Getenv("CACHE_DB")
	cacheDbInt, err := strconv.Atoi(cacheDb)
	if err != nil {
		log.Fatalf("Cache DB Error: %v", err)
	}
	ttl := time.Second * 60
	cbTtl := time.Minute * 5
	redisClient := cache.NewRedisClient(cacheHost, cachePort, cachePassword, cacheDbInt, ttl, cbTtl, ctx)

	circuitBreakerPost := circuit_breaker.NewPostCircuitBreaker(5, redisClient)
	circuitBreakerMessage := circuit_breaker.NewMessageCircuitBreaker(5, redisClient)

	userServicePort := os.Getenv("GATEWAY_SERVICE_PORT")
	accountServiceUrl := os.Getenv("ACCOUNT_SERVICE_URL")
	msgServiceUrl := os.Getenv("MSG_SERVICE_URL")
	postServiceUrl := os.Getenv("POST_SERVICE_URL")
	userApiServer := router.NewUserApiServer(userServicePort, accountServiceUrl, msgServiceUrl, postServiceUrl,
		redisClient, circuitBreakerPost, circuitBreakerMessage)
	err = userApiServer.Run()
	if err != nil {
		log.Fatalf("can't start the userService %s", err)
	}
}
