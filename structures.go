package main //Import the main package
// Create the Block data structure
// A block contains this info:
type Block struct {
	Timestamp         int64  // the time when the block was created
	PreviousBlockHash []byte // the hash of the previous block
	MyBlockHash       []byte // the hash of the current block
	AllData           []byte // the data or transactions (body info)
}

// Prepare the Blockchain data structure :
type Blockchain struct {
	Blocks []*Block // remember a blockchain is a series of blocks
}
