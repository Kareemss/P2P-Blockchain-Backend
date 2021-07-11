//Time to put everything together and test
package main

import (
	"fmt" // just for printing something on the screen
)

func main() {
	newblockchain := NewBlockchain() // Initialize the blockchain
	// create 2 blocks and add 2 transactions
	newblockchain.AddBlock("first transaction")  // first block containing one tx
	newblockchain.AddBlock("Second transaction") // second block containing one tx
	// Now print all the blocks and their contents
	for i, block := range newblockchain.Blocks { // iterate on each block
		fmt.Printf("Block ID : %d \n", i)                                        // print the block ID
		fmt.Printf("Timestamp : %d \n", block.Timestamp+int64(i))                // print the timestamp of the block, to make them different, we just add a value i
		fmt.Printf("Hash of the block : %x\n", block.MyBlockHash)                // print the hash of the block
		fmt.Printf("Hash of the previous Block : %x\n", block.PreviousBlockHash) // print the hash of the previous block
		fmt.Printf("All the transactions : %s\n", block.AllData)                 // print the transactions
	} // our blockchain will be printed
}
