package main

import "user_service/internal/server"

func main() {
	//HOST = os.Getenv("USER_SERVICE_HOST")
	//PORT = os.Getenv("USER_SERVICE_PORT")
	apiServer := server.NewApiServer("", ":8080")
	apiServer.Run()
}
