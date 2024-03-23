package main

import (
	"log"
	"os"
	"user_service/internal/server"
)

func main() {
	userServicePort := os.Getenv("USER_SERVICE_PORT")
	accountServicePort := os.Getenv("ACCOUNT_SERVICE_PORT")
	msgServicePort := os.Getenv("MSG_SERVICE_PORT")
	postServicePort := os.Getenv("POST_SERVICE_PORT")
	userApiServer := server.NewUserApiServer(userServicePort, accountServicePort, msgServicePort, postServicePort)
	err := userApiServer.Run()
	if err != nil {
		log.Fatalf("can't start the userService %s", err)
	}
}
