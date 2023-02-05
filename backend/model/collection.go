package model

import (
	"context"
	"personal-collections/database"

	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collection struct {
	Id      string   `json:"id" bson:"_id,omitempty"`
	Content string   `json:"content" bson:"content"`
	Author  string   `json:"author" bson:"author"`
	Book    string   `json:"book" bson:"book"`
	Tags    []string `json:"tags" bson:"tags"`
}

func (c *Collection) GetByID(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	err := database.DB.Collection(database.COL_Collection).FindOne(context.TODO(), M{"_id": oid}).Decode(c)
	return err
}

func GetCollectionsByIDs(ids []string) ([]Collection, error) {
	var collections []Collection
	for _, id := range ids {
		oid, _ := primitive.ObjectIDFromHex(id)
		var collection Collection
		err := database.DB.Collection(database.COL_Collection).FindOne(context.TODO(), M{"_id": oid}).Decode(&collection)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

type Filter struct {
	Keyword string   `json:"keyword"`
	Author  string   `json:"author"`
	Book    string   `json:"book"`
	Tags    []string `json:"tags"`
}

func ListCollections(filter Filter) ([]Collection, error) {
	var collections []Collection
	var filterMap = M{}
	if filter.Keyword != "" {
		filterMap["content"] = M{"text": M{"$search": filter.Keyword}}
	}
	if filter.Author != "" {
		filterMap["author"] = filter.Author
	}
	if filter.Book != "" {
		filterMap["book"] = filter.Book
	}
	if len(filter.Tags) != 0 {
		filterMap["tags"] = M{"$elemMatch": M{"$in": filter.Tags}}
	}
	cursor, err := database.DB.Collection(database.COL_Collection).Find(context.TODO(), filterMap)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var collection Collection
		err := cursor.Decode(&collection)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (c *Collection) Create() error {
	result, err := database.DB.Collection(database.COL_Collection).InsertOne(context.TODO(), c)
	if err == nil {
		c.Id = result.InsertedID.(primitive.ObjectID).Hex()
	}
	return err
}

func (c *Collection) Update() error {
	oid, _ := primitive.ObjectIDFromHex(c.Id)
	_, err := database.DB.Collection(database.COL_Collection).UpdateOne(context.TODO(), M{"_id": oid}, M{"$set": D{
		{Key: "content", Value: c.Content},
		{Key: "author", Value: c.Author},
		{Key: "book", Value: c.Book},
		{Key: "tags", Value: c.Tags}},
	})
	return err
}

func DeleteCollection(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := database.DB.Collection(database.COL_Collection).DeleteOne(context.TODO(), M{"_id": oid})
	return err
}
