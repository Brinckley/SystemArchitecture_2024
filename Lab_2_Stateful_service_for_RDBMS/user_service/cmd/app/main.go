package main

import (
	"log"
	"os"
	"user_service/internal/db"
	"user_service/internal/server"
)

func main() {
	envPort := os.Getenv("USER_SERVICE_PORT")
	accountTableName := os.Getenv("DB_ACCOUNT_NAME")
	postgresDb, err := db.NewPostgresStorage(accountTableName)
	if err != nil {
		log.Fatalf("unable to connect to postgres %s", err)
	}
	log.Println("---------------CONNECTED TO POSTGRES---------------")
	apiServer := server.NewApiServer(envPort, postgresDb)
	apiServer.Run()
}
