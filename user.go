package main //Import the main package
// import (
// 		"crypto/sha256" //crypto library to hash the data
// )

type User struct {
	FullName         string `bson:"name"`
	ID               string `bson:"id,omitempty"`
	PhoneNumber      int    `bson:"phone"`
	Email            string `bson:"email"`
	Address          string `bson:"address,omitempty"`
	SmartMeterNumber int    `bson:"smart-meter-number,omitempty"`
	PasswordHash     string `bson:"passowrdhash"`
}

// // Prepare the Blockchain data structure :
// type Blockchain struct {
// 	Blocks []*Block // remember a blockchain is a series of blocks
// }

// type Data struct {
// 	Seller string `bson:"seller,omitempty"`
// 	Buyer  string `bson:"buyer,omitempty"`
// 	Amount int    `bson:"amount,omitempty"`
// 	Price  int    `bson:"price,omitempty"`
// }

// var Blockchain []Block
