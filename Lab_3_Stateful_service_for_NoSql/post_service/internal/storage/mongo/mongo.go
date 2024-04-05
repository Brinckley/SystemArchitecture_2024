package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"post_service/internal"
	"post_service/internal/storage"
)

type Db struct {
	Collection *mongo.Collection
}

func NewStorage(database *mongo.Database, collection string) storage.Storage {
	return &Db{
		Collection: database.Collection(collection),
	}
}

func (db *Db) Create(ctx context.Context, postDto internal.PostDto) (string, error) {
	result, err := db.Collection.InsertOne(ctx, postDto)
	if err != nil {
		return "", fmt.Errorf("failed to create post error: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectId to hex error: %v", err)
}

func (db *Db) GetById(ctx context.Context, hexId string) (post internal.Post, err error) {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return post, err
	}

	filterById := bson.M{
		"_id": oid,
	}

	result := db.Collection.FindOne(ctx, filterById)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return post, fmt.Errorf("cannot find the document")
		}
		return post, fmt.Errorf("cannot find post with id %s in database err: %v", hexId, err)
	}

	if err := result.Decode(&post); err != nil {
		return post, fmt.Errorf("cannot decode post from result error : %v", err)
	}
	return post, nil
}

func (db *Db) GetByAccountId(ctx context.Context, id string) ([]internal.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Db) Update(ctx context.Context, post internal.Post) error {
	objectId, err := primitive.ObjectIDFromHex(post.Id)
	if err != nil {
		return fmt.Errorf("cannot get post id error %v", err)
	}

	accountBytes, err := bson.Marshal(post)
	if err != nil {
		return fmt.Errorf("fail to marshall post data error %v", err)
	}

	var updateAccountObj bson.M
	err = bson.Unmarshal(accountBytes, &updateAccountObj)
	if err != nil {
		return fmt.Errorf("fail to unmarshall post data as bytes error %v", err)
	}
	delete(updateAccountObj, "_id")

	filterById := bson.M{
		"_id": objectId,
	}
	updateMsg := bson.M{
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

func (db *Db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
