package main

// create the method that adds a new block to a blockchain
// func (blockchain *Blockchain) AddBlock(data string) {
// 	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // the previous block is needed, so let's get it
// 	newBlock := generateBlock(data, PreviousBlock.Hash)        // create a new block containing the data and the hash of the previous block
// 	blockchain.Blocks = append(blockchain.Blocks, newBlock)      // add that block to the chain to create a chain of blocks
// }

/* Create the function that returns the whole blockchain and add the genesis to it first. the genesis block is the first ever mined block, so let's create a function that will return it since it does not exist yet */
type Error struct {
	Index       int
	NewBlock    Block
	OldBlock    Block
	TypeOfError int
	// Types of errors:
	// 1: Unmatching index values
	// 2: Unmatching Hashes
	// 3: NewBlock hash is different when calculated,
	// meaning data has been tampered
}

func ValidateBlockchain(Blockchain []Block) (bool, []Error) {
	var ErrorBlocks []Error
	for i := 0; i < len(Blockchain); i++ {
		Result, ErrorType := isBlockValid(Blockchain[i+1], Blockchain[i])
		if Result == false {
			Error := Error{len(ErrorBlocks), Blockchain[i+1], Blockchain[i], ErrorType}
			ErrorBlocks = append(ErrorBlocks, Error)
		}

	}
	if len(ErrorBlocks) != 0 {
		return false, ErrorBlocks
	}
	return true, ErrorBlocks
}

func FixBlockchains(Blockchains [][]Block) bool {
	for i := 0; i < len(Blockchains); i++ {
		ValidateBlockchain(Blockchains[i])
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
