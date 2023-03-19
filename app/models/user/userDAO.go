package models

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	COLLECTION_NAME = "users"
)

type UserDAO struct {
	collection *mongo.Collection
}

func NewUserDAO(client *mongo.Client) *UserDAO {
	return &UserDAO{
		collection: client.Database(viper.GetString("mongo.database")).Collection(COLLECTION_NAME),
	}
}

func NewUserDAOwithName(client *mongo.Client, dbName, collName string) *UserDAO {
	return &UserDAO{
		collection: client.Database(dbName).Collection(collName),
	}
}

func (dao *UserDAO) Create(user *UserDTO) (*mongo.InsertOneResult, error) {
	result, err := dao.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *UserDAO) GetByID(id string) (*UserDTO, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	var user UserDTO
	err = dao.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) Update(user *UserDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"userId":        user.UserID,
		"displayName":   user.DisplayName,
		"pictureUrl":    user.PictureURL,
		"statusMessage": user.PictureURL,
		"language":      user.Language,
	}}
	result, err := dao.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *UserDAO) CreateOrUpdate(user *UserDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"userId": user.UserID}

	update := bson.M{
		"$set": user,
	}

	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/
	// SetUpsert : Applications use insert and update operations to store and modify data. Sometimes, you need to choose between an insert and an update operation depending on whether the document exist
	options := options.Update().SetUpsert(true)

	result, err := dao.collection.UpdateOne(context.Background(), filter, update, options)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dao *UserDAO) Delete(id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	result, err := dao.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
