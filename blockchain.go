package main

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
