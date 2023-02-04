package model

import (
	"context"
	"personal-collections/database"
	"time"

	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionData struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	User        *User              `bson:"User"`
	HasLoggedIn bool               `bson:"HasLoggedIn"`
	CreateTime  time.Time          `bson:"CreateTime"`
}

func (session *SessionData) Create() error {
	session.User = nil
	session.HasLoggedIn = false
	session.CreateTime = time.Now()
	result, err := DB.Collection(database.COL_SessionData).InsertOne(context.TODO(), session)
	if err == nil {
		session.Id = result.InsertedID.(primitive.ObjectID)
	}
	return err
}

func (session SessionData) Update() error {
	_, err := DB.Collection(database.COL_SessionData).UpdateByID(context.TODO(), session.Id, M{"$set": session})
	return err
}

func (session *SessionData) GetById(id string) (err error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	err = DB.Collection(database.COL_SessionData).FindOne(context.TODO(), M{"_id": oid}).Decode(session)
	return err
}

func (session *SessionData) GetByUserId(id string) (err error) {
	ouid, _ := primitive.ObjectIDFromHex(id)
	err = DB.Collection(database.COL_SessionData).FindOne(context.TODO(), M{"User._id": ouid}).Decode(session)
	return err
}

func (session SessionData) DeleteById(id string) (err error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err = DB.Collection(database.COL_SessionData).DeleteOne(context.TODO(), M{"_id": oid})
	return err
}
