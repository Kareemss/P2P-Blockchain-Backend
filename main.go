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
			genesisBlock := Block{0, t.String(), "", "", Order{"I am the genesis block", "", "", 0, 0}, true}
			genesisBlock.Hash = calculateHash(genesisBlock)
			Blockchain = append(Blockchain, genesisBlock)
			addBlock(genesisBlock, BlockchainDatabase)
		} else {
			Blockchain = append(Blockchain, getGenesisBlockFromDb(BlockchainDatabase))
		}
	}()
	log.Fatal(run())
}

func getGenesisBlockFromDb(database *mongo.Database) Block {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var blocks []Block

	cursor, err := database.Collection("Blocks").Find(ctx, bson.M{"is-genesis": bson.D{{"$eq", true}}})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &blocks); err != nil {
		panic(err)
	}

	return blocks[0]
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
