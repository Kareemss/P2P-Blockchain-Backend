package main //Import the main package
// Create the Block data structure
// A block contains this info:
type Block struct {
	Index     int
	Timestamp string // the time when the block was created
	PrevHash  string // the hash of the previous block
	Hash      string // the hash of the current block
	AllData   string // the data or transactions (body info)
}

// // Prepare the Blockchain data structure :
// type Blockchain struct {
// 	Blocks []*Block // remember a blockchain is a series of blocks
// }

var Blockchain []Block
