package messagemodel

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	COLLECTION_NAME = "messages"
)

type MessageDAO struct {
	collection *mongo.Collection
}

func NewMessageDAO(client *mongo.Client) *MessageDAO {
	return &MessageDAO{
		collection: client.Database(viper.GetString("mongo.database")).Collection(COLLECTION_NAME),
	}
}

func NewMessageDAOwithName(client *mongo.Client, dbName, collName string) *MessageDAO {
	return &MessageDAO{
		collection: client.Database(dbName).Collection(collName),
	}
}

func (dao *MessageDAO) Create(message *MessageDTO) (*mongo.InsertOneResult, error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	result, err := dao.collection.InsertOne(context.Background(), message)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *MessageDAO) GetByID(id string) (*MessageDTO, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	var user MessageDTO
	err = dao.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *MessageDAO) Update(message *MessageDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": message.ID}
	update := bson.M{"$set": bson.M{}}
	result, err := dao.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *MessageDAO) Delete(id string) (*mongo.DeleteResult, error) {
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
