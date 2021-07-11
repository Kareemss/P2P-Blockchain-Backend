package main

// create the method that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // the previous block is needed, so let's get it
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash)        // create a new block containing the data and the hash of the previous block
	blockchain.Blocks = append(blockchain.Blocks, newBlock)      // add that block to the chain to create a chain of blocks
}

/* Create the function that returns the whole blockchain and add the genesis to it first. the genesis block is the first ever mined block, so let's create a function that will return it since it does not exist yet */
func NewBlockchain() *Blockchain { // the function is created
	return &Blockchain{[]*Block{NewGenesisBlock()}} // the genesis block is added first to the chain
}
