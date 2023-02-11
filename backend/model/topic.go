package model

import (
	"context"
	"fmt"
	"personal-collections/database"
	"strings"

	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topic struct {
	Id         string   `json:"id" bson:"_id,omitempty"`
	Title      string   `json:"title" bson:"title"`
	Detail     string   `json:"detail" bson:"detail"`
	Collctions []string `json:"collections" bson:"collections"`
}

func (t *Topic) GetByID(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	err := database.DB.Collection(database.COL_Topic).FindOne(context.TODO(), M{"_id": oid}).Decode(t)
	return err
}

func (t *Topic) GetCollections(filter Filter) ([]Collection, error) {
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
	cursor, err := database.DB.Collection(database.COL_Collection).Find(context.TODO(), M{"_id": M{"$in": t.Collctions}, "$and": []M{filterMap}})
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &collections)
	cursor.Close(context.TODO())
	return collections, nil
}

func ListTopics() ([]Topic, error) {
	var topics []Topic = []Topic{}
	cursor, err := database.DB.Collection(database.COL_Topic).Find(context.TODO(), M{})
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &topics)
	cursor.Close(context.TODO())
	return topics, nil
}

func (t *Topic) Create() error {
	result, err := database.DB.Collection(database.COL_Topic).InsertOne(context.TODO(), t)
	if err == nil {
		t.Id = result.InsertedID.(primitive.ObjectID).Hex()
	}
	return err
}

func (t *Topic) Update() error {
	oid, _ := primitive.ObjectIDFromHex(t.Id)
	_, err := database.DB.Collection(database.COL_Topic).UpdateByID(context.TODO(), oid, M{"$set": M{
		"title":       t.Title,
		"detail":      t.Detail,
		"collections": t.Collctions},
	})
	return err
}

func (t *Topic) AddCollection(cid string) error {
	oid, _ := primitive.ObjectIDFromHex(t.Id)
	t.Collctions = append(t.Collctions, cid)
	_, err := database.DB.Collection(database.COL_Topic).UpdateByID(context.TODO(), oid, M{"$set": M{"collections": t.Collctions}})
	return err
}

func (t *Topic) RemoveCollection(cid string) error {
	oid, _ := primitive.ObjectIDFromHex(t.Id)
	collections := []string{}
	for _, c := range t.Collctions {
		if c != cid {
			collections = append(collections, c)
		}
	}
	t.Collctions = collections
	_, err := database.DB.Collection(database.COL_Topic).UpdateByID(context.TODO(), oid, M{"$set": M{"collections": t.Collctions}})
	return err
}

func DeleteTopic(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := database.DB.Collection(database.COL_Topic).DeleteOne(context.TODO(), M{"_id": oid})
	return err
}
