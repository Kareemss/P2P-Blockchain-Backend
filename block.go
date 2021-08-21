package main

import (
	// We will need these libraries:
	//"bytes"         // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	"crypto/sha256" //crypto library to hash the data
	//"strconv"       // for conversion
	"time" // the time for our timestamp
	//"github.com/davecgh/go-spew/spew"
	//"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
	"encoding/hex"
	//"encoding/json"
	//"io"
	//"log"
	//"net/http"
	//"os"
)

// Now let's create a method for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.AllData.Issuer + block.AllData.Seller + block.AllData.Buyer + string(block.AllData.Amount) + string(block.AllData.Price) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Create a function for new block generation and return that block
func generateBlock(oldBlock Block, AllData Order) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.AllData = AllData
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

/* let's now create the genesis block function that will return the first block. The genesis block is the first block on the chain */

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
