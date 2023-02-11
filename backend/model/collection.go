package model

import (
	"context"
	"fmt"
	"personal-collections/database"
	"strings"

	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collection struct {
	Id      string   `json:"id" bson:"_id,omitempty"`
	Starred bool     `json:"starred" bson:"starred"`
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
	collections := []Collection{}
	var oids []primitive.ObjectID
	for _, id := range ids {
		oid, _ := primitive.ObjectIDFromHex(id)
		oids = append(oids, oid)
	}
	cursor, err := database.DB.Collection(database.COL_Collection).Find(context.TODO(), M{"_id": M{"$in": oids}})
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &collections)
	cursor.Close(context.TODO())
	return collections, nil
}

type Filter struct {
	Keyword string   `json:"keyword"`
	Author  string   `json:"author"`
	Book    string   `json:"book"`
	Tags    []string `json:"tags"`
}

func ListStarredCollections(filter Filter) ([]Collection, error) {
	collections := []Collection{}
	keywords := strings.Split(filter.Keyword, " ")
	for i, keyword := range keywords {
		keywords[i] = fmt.Sprintf("(?=.*%s)", keyword)
	}
	regexpStr := fmt.Sprintf("%s.*", strings.Join(keywords, ""))
	var filterMap = M{"starred": true}
	if filter.Keyword != "" {
		filterMap["content"] = M{"$regex": regexpStr}
	}
	if filter.Author != "" {
		filterMap["author"] = M{"$regex": filter.Author}
	}
	if filter.Book != "" {
		filterMap["book"] = M{"$regex": filter.Book}
	}
	if len(filter.Tags) != 0 {
		filterMap["tags"] = M{"$elemMatch": M{"$in": filter.Tags}}
	}
	cursor, err := database.DB.Collection(database.COL_Collection).Find(context.TODO(), filterMap)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &collections)
	cursor.Close(context.TODO())
	return collections, nil
}

func ListCollections(filter Filter) ([]Collection, error) {
	collections := []Collection{}
	keywords := strings.Split(filter.Keyword, " ")
	for i, keyword := range keywords {
		keywords[i] = fmt.Sprintf("(?=.*%s)", keyword)
	}
	regexpStr := fmt.Sprintf("%s.*", strings.Join(keywords, ""))
	var filterMap = M{}
	if filter.Keyword != "" {
		filterMap["content"] = M{"$regex": regexpStr}
	}
	if filter.Author != "" {
		filterMap["author"] = M{"$regex": filter.Author}
	}
	if filter.Book != "" {
		filterMap["book"] = M{"$regex": filter.Book}
	}
	if len(filter.Tags) != 0 {
		filterMap["tags"] = M{"$elemMatch": M{"$in": filter.Tags}}
	}
	cursor, err := database.DB.Collection(database.COL_Collection).Find(context.TODO(), filterMap)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &collections)
	cursor.Close(context.TODO())
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
		{Key: "starred", Value: c.Starred},
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
