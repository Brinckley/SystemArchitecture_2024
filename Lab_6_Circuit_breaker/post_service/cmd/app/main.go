package main

import (
	"context"
	"log"
	"os"
	"post_service/internal/service"
	"post_service/internal/storage/mongo"
)

func main() {
	ctx := context.Background()
	appPort := os.Getenv("POST_SERVICE_PORT")
	mongoUri := os.Getenv("MONGO_URI")
	collectionName := os.Getenv("MONGO_COLLECTION")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("MONGO_DB")
	mongoDatabase, err := mongo.NewMongoClient(ctx, mongoUri, username, password, database)
	if err != nil {
		log.Fatalf("unable to connect to mongo error %v", err)
	}
	storage := mongo.NewStorage(mongoDatabase, collectionName)

	log.Println("---------------CONNECTED TO MONGO FROM POST SERVICE---------------")
	apiServer := service.NewPostApiServer(appPort, storage, &ctx)
	apiServer.Run()
}
