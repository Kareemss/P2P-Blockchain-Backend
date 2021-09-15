package main

import (
	"crypto/sha256" //crypto library to hash the data
	"encoding/hex"
	"fmt"
	"strconv"
	"time" // the time for our timestamp
)

// Now let's create a method for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp +
		block.AllData.Issuer + block.AllData.Seller +
		block.AllData.Buyer +
		fmt.Sprintf("%f", block.AllData.Amount) +
		fmt.Sprintf("%f", block.AllData.Price) + block.PrevHash
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

func isBlockValid(newBlock, oldBlock Block) (bool, int) {
	if oldBlock.Index+1 != newBlock.Index {
		return false, 1
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false, 2
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false, 3
	}

	return true, 0
}
