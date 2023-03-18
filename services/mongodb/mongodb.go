package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	POOLSIZE = 10
)

func ConnectMongoDB(ip, port, username, password, poolsize, database string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, ip, port, database)
	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(uint64(POOLSIZE))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewMongoClient() (*mongo.Client, error) {
	ip := viper.GetString("mongo.ip")
	port := viper.GetString("mongo.port")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")
	poolsize := viper.GetString("mongo.poolsize")
	database := viper.GetString("mongo.database")
	return ConnectMongoDB(ip, port, username, password, poolsize, database)
}
