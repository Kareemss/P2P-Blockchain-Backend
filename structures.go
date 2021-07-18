package main //Import the main package
// Create the Block data structure
// A block contains this info:
type Block struct {
	Index     int    `bson:"_id"`
	Timestamp string `bson:"timestamp,omitempty"` // the time when the block was created
	PrevHash  string `bson:"prev-hash"`           // the hash of the previous block
	Hash      string `bson:"hash"`                // the hash of the current block
	AllData   string `bson:"all-data,omitempty"`  // the data or transactions (body info)
	IsGenesis bool   `bson:"is-genesis,omitempty"`
}

// // Prepare the Blockchain data structure :
// type Blockchain struct {
// 	Blocks []*Block // remember a blockchain is a series of blocks
// }

var Blockchain []Block
