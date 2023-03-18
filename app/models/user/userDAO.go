package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {
	collection *mongo.Collection
}

func NewUserDAO(client *mongo.Client, dbName, collName string) *UserDAO {
	return &UserDAO{
		collection: client.Database(dbName).Collection(collName),
	}
}

func (dao *UserDAO) CreateUser(user *UserDTO) (*mongo.InsertOneResult, error) {
	result, err := dao.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *UserDAO) GetUserByID(id string) (*UserDTO, error) {
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

func (dao *UserDAO) UpdateUser(user *UserDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}}
	result, err := dao.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *UserDAO) DeleteUser(id string) (*mongo.DeleteResult, error) {
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
