package main

import (
	"context"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteQuery struct {
	Database     string
	Collection   string
	Query        string
	Condition    interface{}
	DeletionType int
	// Types of Deletions:
	// 1: Document
	// 2: Collection
}

func DeleteOneFromDB(Database string, Collection string, Query string, Condition interface{}) *mongo.DeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DDatabase := connectToDb(Database)
	DCollection := DDatabase.Collection(Collection)
	result, err := DCollection.DeleteMany(ctx, bson.M{Query: Condition})
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func DeleteCollection(Database string, Collection string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DDatabase := connectToDb(Database)
	DCollection := DDatabase.Collection(Collection)
	if err := DCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}

}
