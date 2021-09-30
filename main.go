package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"time" // the time for our timestamp

	"github.com/davecgh/go-spew/spew"

	"github.com/joho/godotenv"

	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	mongoconnect()
	go func() {
		t := time.Now()

		// Connecting to MongoDB
		spew.Dump("Connecting to Server")
		BlockchainDatabase := connectToDb("Blockchain")

		if !presentGenesisBlockInDb(BlockchainDatabase) {
			genesisBlock := Block{0, t.String(), "", "", Order{0, "I am the genesis block", "", "", 0, 0}, true}
			genesisBlock.Hash = calculateHash(genesisBlock)
			Blockchain = append(Blockchain, genesisBlock)
			addBlock(genesisBlock, BlockchainDatabase)
		} else {
			Blockchain = getBlockchainFromDb(BlockchainDatabase)
		}
		Market = getMarketFromDB()
	}()
	log.Fatal(run())
}

func getBlockchainFromDb(database *mongo.Database) []Block {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var blocks []Block

	cursor, err := database.Collection("Blocks").Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &blocks); err != nil {
		panic(err)
	}

	return blocks
}
func getMarketFromDB() []Order {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MDatabase := connectToDb("Market")
	MCollection := MDatabase.Collection("Orders")
	var Orders []Order

	cursor, err := MCollection.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &Orders); err != nil {
		panic(err)
	}

	return Orders
}

func presentGenesisBlockInDb(database *mongo.Database) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var blocks []Block

	cursor, err := database.Collection("Blocks").Find(ctx, bson.M{"is-genesis": bson.D{{"$exists", true}}})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &blocks); err != nil {
		panic(err)
	}

	// Only for documentation purposes
	if blocks != nil {
		spew.Dump("Genesis Block is present. So a new one will not be added to the database")
	} else {
		spew.Dump("Genesis block added to the database")
	}

	return blocks != nil
}
