package main

import (
	"log"
	"os"
	"post_service/internal/repository/postgres"
	"post_service/internal/server"
)

func main() {
	envPort := os.Getenv("POST_SERVICE_PORT")
	postTableName := os.Getenv("DB_POST_TABLE_NAME")
	postgresDb, err := postgres.NewPostgresStorage(postTableName)
	if err != nil {
		log.Fatalf("unable to connect to postgres %s", err)
	}
	log.Println("---------------CONNECTED TO POSTGRES FROM POST SERVICE---------------")
	apiServer := server.NewPostApiServer(envPort, postgresDb)
	apiServer.Run()
}
