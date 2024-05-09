package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
		log.Printf("INSERTED ID %v", oid)
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectId to hex error: %v", err)
}

func (db *Db) GetById(ctx context.Context, hexId string) (post internal.Post, err error) {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return post, fmt.Errorf("failed to convert id error %v", err)
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

func (db *Db) GetByAccountId(ctx context.Context, hexId string) (posts []internal.Post, err error) {
	filterReceiver := bson.D{{"account_id", hexId}}

	found, _ := db.Collection.Find(ctx, filterReceiver)
	if found.Err() != nil {
		return posts, fmt.Errorf("fail to get all messages for user error %v", err)
	}

	if err = found.All(ctx, &posts); err != nil {
		return posts, fmt.Errorf("fail to handle fetched data error %v", err)
	}
	return posts, nil
}

func (db *Db) Update(ctx context.Context, post internal.Post) error {
	objectId, err := primitive.ObjectIDFromHex(post.Id)
	if err != nil {
		return fmt.Errorf("cannot get post id error %v", err)
	}

	postBytes, err := bson.Marshal(post)
	if err != nil {
		return fmt.Errorf("fail to marshall post data error %v", err)
	}

	var updatePostObj bson.M
	err = bson.Unmarshal(postBytes, &updatePostObj)
	if err != nil {
		return fmt.Errorf("fail to unmarshall post data as bytes error %v", err)
	}
	delete(updatePostObj, "_id")

	filterById := bson.M{
		"_id": objectId,
	}
	updatePost := bson.M{
		"$set": updatePostObj,
	}

	result, err := db.Collection.UpdateOne(ctx, filterById, updatePost)
	if err != nil {
		return fmt.Errorf("fail to update the post error %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("fail to find the post for update error %v", err)
	}

	return nil
}

func (db *Db) Delete(ctx context.Context, accountId, hexId string) error {
	oid, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return fmt.Errorf("fail to handle id error %v", err)
	}

	filterById := bson.M{
		"_id": oid,
	}

	deleteResult, err := db.Collection.DeleteOne(ctx, filterById)
	if err != nil {
		return fmt.Errorf("fail to delete post error %v", err)
	}
	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("no posts deleted")
	}

	return nil
}
