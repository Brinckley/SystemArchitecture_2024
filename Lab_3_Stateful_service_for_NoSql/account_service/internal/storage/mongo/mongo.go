package mongo

import (
	"account_service/internal"
	"account_service/internal/storage"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	Collection *mongo.Collection
}

func NewStorage(database *mongo.Database, collection string) storage.Storage {
	return &Db{
		Collection: database.Collection(collection),
	}
}

func (db *Db) Create(ctx context.Context, account internal.AccountDto) (string, error) {
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

func (db *Db) GetAll(ctx context.Context) (accounts []internal.Account, err error) {
	allFilter := bson.M{}
	found, _ := db.Collection.Find(ctx, allFilter)
	if found.Err() != nil {
		return accounts, fmt.Errorf("fail to get all accounts error %v", err)
	}

	if err = found.All(ctx, &accounts); err != nil {
		return accounts, fmt.Errorf("fail to handle fetched data error %v", err)
	}
	return accounts, nil
}

func (db *Db) GetById(ctx context.Context, hexId string) (account internal.Account, err error) {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return account, err
	}

	filterById := bson.M{
		"_id": oid,
	}

	result := db.Collection.FindOne(ctx, filterById)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return account, fmt.Errorf("cannot find the document")
		}
		return account, fmt.Errorf("cannot find user with id %s in database err: %v", hexId, err)
	}

	if err := result.Decode(&account); err != nil {
		return account, fmt.Errorf("cannot decode user from result error : %v", err)
	}
	return account, nil
}

func (db *Db) Update(ctx context.Context, account internal.Account) error {
	objectId, err := primitive.ObjectIDFromHex(account.Id)
	if err != nil {
		return fmt.Errorf("cannot get user id error %v", err)
	}

	accountBytes, err := bson.Marshal(account)
	if err != nil {
		return fmt.Errorf("fail to marshall account data error %v", err)
	}

	var updateAccountObj bson.M
	err = bson.Unmarshal(accountBytes, &updateAccountObj)
	if err != nil {
		return fmt.Errorf("fail to unmarshall account data as bytes error %v", err)
	}
	delete(updateAccountObj, "_id")

	filterById := bson.M{
		"_id": objectId,
	}
	updateAcc := bson.M{
		"$set": updateAccountObj,
	}

	result, err := db.Collection.UpdateOne(ctx, filterById, updateAcc)
	if err != nil {
		return fmt.Errorf("fail to update the account error %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("fail to find the account for update error %v", err)
	}

	return nil
}

func (db *Db) Delete(ctx context.Context, hexId string) error {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return fmt.Errorf("fail to handle id error %v", err)
	}

	filterById := bson.M{
		"_id": oid,
	}

	deleteResult, err := db.Collection.DeleteOne(ctx, filterById)
	if err != nil {
		return fmt.Errorf("fail to delete account error %v", err)
	}
	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("no accounts deleted")
	}

	return nil
}
