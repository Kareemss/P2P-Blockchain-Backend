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
type UpdateBalanceQuery struct {
	Email   string
	Asset   string
	Balance int
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
func UpdateFromDB(Database string, Collection string,
	Query string, Condition interface{}, Field string,
	NewValue interface{}) *mongo.UpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DDatabase := connectToDb(Database)
	DCollection := DDatabase.Collection(Collection)
	result, err := DCollection.UpdateOne(
		ctx,
		bson.M{Query: Condition},
		bson.D{
			{"$set", bson.D{{Field, NewValue}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func AddBalance(Email string, Asset string, Balance int) *mongo.UpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	UserDatabase := connectToDb("Users")
	UserCollection := UserDatabase.Collection("Users")
	filterCursor, err := UserCollection.Find(ctx, bson.M{"email": Email})
	if err != nil {
		log.Fatal(err)
	}
	var Profiles []User
	if err = filterCursor.All(ctx, &Profiles); err != nil {
		log.Fatal(err)
	}
	Profile := Profiles[0]
	if Asset == "energy-balance" {
		Balance = Balance + Profile.EnergyBalance
	} else if Asset == "currency-balance" {
		Balance = Balance + Profile.CurrencyBalance
	}
	result := UpdateFromDB("Users", "Users", "email", Email, Asset, Balance)
	return result
}
