package main

import (
	"account_service/internal/server"
	"account_service/internal/storage/postgres"
	"log"
	"os"
)

func main() {
	envPort := os.Getenv("ACCOUNT_SERVICE_PORT")
	accountTableName := os.Getenv("DB_ACCOUNT_TABLE_NAME")
	postgresDb, err := postgres.NewPostgresStorage(accountTableName)
	if err != nil {
		log.Fatalf("unable to connect to postgres %s", err)
	}
	log.Println("---------------CONNECTED TO POSTGRES FROM ACCOUNT SERVICE---------------")
	apiServer := server.NewAccountApiServer(envPort, postgresDb)
	apiServer.Run()
}
