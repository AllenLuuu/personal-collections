package model

import (
	"context"
	"personal-collections/database"

	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id	   string `json:"id" bson:"_id,omitempty"` // bson:"_id,omitempty" is used to ignore the id field when inserting data
	Username string `json:"username" bson:"username"` // bson:"username" is used to specify the field name in the database
	Password string `json:"password" bson:"password"`
}

func (u *User) GetByUsername(username string) error {
	err := DB.Collection(database.COL_User).FindOne(context.TODO(), M{"username": username}).Decode(u)
	return err
}

func (u *User) GetByID(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	err := database.DB.Collection(database.COL_User).FindOne(context.TODO(), M{"_id": oid}).Decode(u)
	return err
}
