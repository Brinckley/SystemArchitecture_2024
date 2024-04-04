package mongo

import (
	"account_service/internal"
	"account_service/internal/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	Collection *mongo.Collection
}

func NewStorage(database mongo.Database, collection string) storage.Storage {
	return &Db{
		Collection: database.Collection(collection),
	}
}

func (db *Db) Create(ctx context.Context, account internal.Account) (string, error) {
	result, err := db.Collection.InsertOne(ctx, account)
	if err != nil {
		return "", fmt.Errorf("failed to create account error: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectId to hex error: %v", err)
}

func (db *Db) GetAll(ctx context.Context) ([]internal.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Db) GetById(ctx context.Context, hexId string) (account internal.Account, err error) {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return internal.Account{}, err
	}
	filter := bson.M{
		"_id": oid,
	}
	result := db.Collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return account, fmt.Errorf("cannot find user with id %s in database err: %v", hexId, err)
	}
	if err := result.Decode(&account); err != nil {
		return account, fmt.Errorf("cannot decode user from result error : %v", err)
	}
	return account, nil
}

func (db *Db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (db *Db) Update(ctx context.Context, account internal.Account) error {
	//TODO implement me
	panic("implement me")
}
