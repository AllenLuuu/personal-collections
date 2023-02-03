package model

import (
	"personal-collections/database"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func Init() error {
	logrus.Info("[model] Init...")
	DB = database.DB
	logrus.Info("[model] Init Success")
	return nil
}