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

func (db *Db) GetById(ctx context.Context, userId, hexMsgId string) (message internal.Message, err error) {
	oid, err := primitive.ObjectIDFromHex(hexMsgId)
	if err != nil {
		return message, fmt.Errorf("cannot handle message id error %v", err)
	}

	filterByIdReceiverById := bson.M{
		"_id":         oid,
		"receiver_id": userId,
	}

	result := db.Collection.FindOne(ctx, filterByIdReceiverById)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			filterByIdSenderById := bson.M{
				"_id":       oid,
				"sender_id": userId,
			}
			result = db.Collection.FindOne(ctx, filterByIdSenderById)
			if result.Err() != nil {
				return message, fmt.Errorf("failed to find message by id error %v", err)
			}
			if err := result.Decode(&message); err != nil {
				return message, fmt.Errorf("cannot decode user from result error : %v", err)
			}
			return message, nil
		}
		return message, fmt.Errorf("cannot find message with id %s in database err: %v", hexMsgId, err)
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