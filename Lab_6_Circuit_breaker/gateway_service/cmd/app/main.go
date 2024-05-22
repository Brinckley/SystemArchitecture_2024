package main

import (
	"Gateway/internal/server/router"
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
	redisClient := cache.NewRedisClient(cacheHost, cachePort, cachePassword, cacheDbInt, time.Second*60, ctx)

	userServicePort := os.Getenv("GATEWAY_SERVICE_PORT")
	accountServiceUrl := os.Getenv("ACCOUNT_SERVICE_URL")
	msgServiceUrl := os.Getenv("MSG_SERVICE_URL")
	postServiceUrl := os.Getenv("POST_SERVICE_URL")
	userApiServer := router.NewUserApiServer(userServicePort, accountServiceUrl, msgServiceUrl, postServiceUrl, redisClient)
	err = userApiServer.Run()
	if err != nil {
		log.Fatalf("can't start the userService %s", err)
	}
}
