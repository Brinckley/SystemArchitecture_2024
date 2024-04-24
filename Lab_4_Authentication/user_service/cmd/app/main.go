package main

import (
	"log"
	"os"
	"user_service/internal/server/handler"
)

func main() {
	userServicePort := os.Getenv("USER_SERVICE_PORT")
	accountServiceUrl := os.Getenv("ACCOUNT_SERVICE_URL")
	msgServiceUrl := os.Getenv("MSG_SERVICE_URL")
	postServiceUrl := os.Getenv("POST_SERVICE_URL")
	userApiServer := handler.NewUserApiServer(userServicePort, accountServiceUrl, msgServiceUrl, postServiceUrl)
	err := userApiServer.Run()
	if err != nil {
		log.Fatalf("can't start the userService %s", err)
	}
}
