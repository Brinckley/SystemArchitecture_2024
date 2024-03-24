package main

import (
	"log"
	"message_service/internal/repository/postgres"
	"message_service/internal/server"
	"os"
)

func main() {
	envPort := os.Getenv("MSG_SERVICE_PORT")
	messageTableName := os.Getenv("DB_MESSAGE_TABLE_NAME")
	postgresDb, err := postgres.NewPostgresStorage(messageTableName)
	if err != nil {
		log.Fatalf("unable to connect to postgres %s", err)
	}
	log.Println("---------------CONNECTED TO POSTGRES FROM MESSAGE SERVICE---------------")
	apiServer := server.NewMessageApiServer(envPort, postgresDb)
	apiServer.Run()
}
