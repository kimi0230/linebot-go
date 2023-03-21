package usermodel

import (
	"context"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	COLLECTION_NAME = "users"
)

type PageQueryArgs struct {
	Keyword string `json:"keyword,omitempty" form:"keyword,omitempty" binding:"-"`
	Limit   int    `json:"limit,default=100" form:"limit,default=100" binding:"required,number"`
	Order   string `json:"order,default=desc" form:"order,default=desc" binding:"oneof=desc asc"`
	By      string `json:"by,default=updated_at" form:"by,default=updated_at" binding:"-"`
	Page    int    `json:"page,default=1" form:"page,default=1" binding:"number"`
}

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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
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

func (dao *UserDAO) GetByQuery(filter interface{}, limit int64, skip int64, order string, by string) (*[]UserDTO, error) {
	var users []UserDTO
	ctx := context.Background()
	opts := options.Find().SetSkip(skip).SetLimit(limit)
	// Handle Order
	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/sort/
	order = strings.ToLower(order)
	if order == "desc" {
		opts = opts.SetSort(bson.M{by: -1})
	} else {
		opts = opts.SetSort(bson.M{by: 1})
	}

	cur, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &users); err != nil {
		panic(err)
	}

	return &users, nil
}

func (dao *UserDAO) Update(user *UserDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": user.ID}
	// filter := bson.M{"userId": user.UserID}
	update := bson.M{"$set": bson.M{
		"userId":        user.UserID,
		"displayName":   user.DisplayName,
		"pictureUrl":    user.PictureURL,
		"statusMessage": user.PictureURL,
		"language":      user.Language,
		"updated_at":    time.Now(),
	}}
	result, err := dao.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dao *UserDAO) CreateOrUpdateByUserID(user *UserDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"userId": user.UserID}

	update := bson.M{
		"$set":         user,
		"$setOnInsert": bson.M{"created_at": time.Now()},
		"$currentDate": bson.M{"updated_at": true},
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

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/count/
func (dao *UserDAO) CountDocuments(filter interface{}) (int64, error) {
	ctx := context.Background()
	count, err := dao.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
