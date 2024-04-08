package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"message_service/internal"
	"message_service/internal/storage"
)

type Db struct {
	Collection *mongo.Collection
}

func NewStorage(database *mongo.Database, collection string) storage.Storage {
	return &Db{
		Collection: database.Collection(collection),
	}
}

func (db *Db) Create(ctx context.Context, msg internal.MessageDto) (string, error) {
	result, err := db.Collection.InsertOne(ctx, msg)
	if err != nil {
		return "", fmt.Errorf("failed to create message error: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectId to hex error: %v", err)
}

func (db *Db) GetById(ctx context.Context, hexId string) (message internal.Message, err error) {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return message, fmt.Errorf("cannot handle message id error %v", err)
	}

	filterById := bson.M{
		"_id": oid,
	}

	result := db.Collection.FindOne(ctx, filterById)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return message, fmt.Errorf("cannot find the document")
		}
		return message, fmt.Errorf("cannot find message with id %s in database err: %v", hexId, err)
	}

	if err := result.Decode(&message); err != nil {
		return message, fmt.Errorf("cannot decode user from result error : %v", err)
	}
	return message, nil
}

func (db *Db) GetByDestId(ctx context.Context, hexId string) (messages []internal.Message, err error) {
	filterReceiver := bson.D{{"receiver_id", hexId}}

	found, _ := db.Collection.Find(ctx, filterReceiver)
	if found.Err() != nil {
		return messages, fmt.Errorf("fail to get all messages for user error %v", err)
	}

	if err = found.All(ctx, &messages); err != nil {
		return messages, fmt.Errorf("fail to handle fetched data error %v", err)
	}
	return messages, nil
}
