package database

import (
	"context"
	"fmt"
	"personal-collections/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func initMongo() error {
	cfg := config.Mongo
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetConnectTimeout(time.Second * 2)
	clientOptions.SetSocketTimeout(time.Second * 2)
	clientOptions.SetServerSelectionTimeout(time.Second * 2)

	Client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}

	DB = Client.Database(cfg.Database)
	return nil

}
