//Time to put everything together and test
package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// We will need these libraries:
	//      // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	//"crypto/sha256" //crypto library to hash the data
	//"strconv"       // for conversion
	"time" // the time for our timestamp

	"github.com/davecgh/go-spew/spew"

	//"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//"encoding/hex"
	//"encoding/json"
	//"io"
	"log"
	//"net/http"
	//"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()

		// Connecting to MongoDB
		spew.Dump("Connecting to Server")
		database := connectToDb()

		if !presentGenesisBlockInDb(database) {
			genesisBlock := Block{0, t.String(), "", "", "I am the genesis block", true}
			Blockchain = append(Blockchain, genesisBlock)
			addBlock(genesisBlock, database)
		} else {
			Blockchain = append(Blockchain, getGenesisBlockFromDb(database))
		}
	}()
	log.Fatal(run())
}

func getGenesisBlockFromDb(database *mongo.Database) Block {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var blocks []Block

	cursor, err := database.Collection("blocks").Find(ctx, bson.M{"is-genesis": bson.D{{"$eq", true}}})

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

	cursor, err := database.Collection("blocks").Find(ctx, bson.M{"is-genesis": bson.D{{"$exists", true}}})

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
