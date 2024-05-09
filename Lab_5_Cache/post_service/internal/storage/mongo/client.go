package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, mongoUri, username, password, database string) (*mongo.Database, error) {
	mongoDbUri := fmt.Sprintf(mongoUri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongoDbUri).SetServerAPIOptions(serverAPI)
	clientOptions.SetAuth(options.Credential{
		AuthSource: database,
		Username:   username,
		Password:   password,
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo error %v", err)
	}

	var result bson.M
	if err := client.Database(database).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client.Database(database), nil
}
