package main

import (
	"crypto/sha256" //crypto library to hash the data
	"encoding/hex"
	"fmt"
	"strconv"
	"time" // the time for our timestamp
	//"github.com/davecgh/go-spew/spew"
)

// Now let's create a method for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp +
		strconv.Itoa(block.AllData.OrderID) +
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
func generateBlock(oldBlock Block, AllData Order) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.AllData = AllData
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	// // res, _ := isBlockValid(newBlock, Blockchain[len(Blockchain)-1])
	// // spew.Dump(res)
	// if true {
	// 	Blockchain = append(Blockchain, newBlock)
	// 	// replaceChain(newBlockchain)
	// 	// spew.Dump(Blockchain)

	// 	BlockchainDatabase := connectToDb("Blockchain")
	// 	addBlock(newBlock, BlockchainDatabase)
	// 	// spew.Dump(Blockchain)
	// }

	return newBlock
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

// func call(){
// 	for i:= 0; i < 1000; i++ {
// 		dosomething()

// 	}
// }

// func dosomething(){
// 	// Lengthy process
// }
