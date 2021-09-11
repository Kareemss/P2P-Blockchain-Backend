package main

import (
	"context"
	_ "log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectToDb(Choice string) *mongo.Database {
	// Replace the uri string with your MongoDB deployment's connection string.
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	cluster := os.Getenv("DB_CLUSTER_ADDR")
	//uri := "mongodb+srv://" + username + ":" + password + "@" + cluster + ".bzh1l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	uri := "mongodb://" + username + ":" + password + "@" + cluster + "-shard-00-00.bzh1l.mongodb.net:27017," + cluster + "-shard-00-01.bzh1l.mongodb.net:27017," + cluster + "-shard-00-02.bzh1l.mongodb.net:27017/myFirstDatabase?ssl=true&replicaSet=atlas-hmhvdy-shard-0&authSource=admin&retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	spew.Dump("Successfully connected and pinged.")

	BlockchainDatabase := client.Database("Blockchain")
	UserDatabase := client.Database("Users")
	MarketDatabase := client.Database("Market")
	var Database *mongo.Database
	if Choice == "Blockchain" {
		Database = BlockchainDatabase
	} else if Choice == "Users" {
		Database = UserDatabase
	} else if Choice == "Market" {
		Database = MarketDatabase
	}

	return Database
}

func addBlock(block Block, database *mongo.Database) *mongo.InsertOneResult {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blocksCollection := database.Collection("Blocks")

	insertionResult, err := blocksCollection.InsertOne(ctx, block)
	if err != nil {
		panic(err)
	}

	return insertionResult
}

func AddUser(User User, database *mongo.Database) *mongo.InsertOneResult {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Users := database.Collection("Users")

	insertionResult, err := Users.InsertOne(ctx, User)
	if err != nil {
		panic(err)
	}

	return insertionResult
}

func AddOrder(Order Order, database *mongo.Database) *mongo.InsertOneResult {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Users := database.Collection("Orders")

	insertionResult, err := Users.InsertOne(ctx, Order)
	if err != nil {
		panic(err)
	}

	return insertionResult
}
