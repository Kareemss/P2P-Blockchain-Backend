//Time to put everything together and test
package main

import (
	// We will need these libraries:
	//      // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	//"crypto/sha256" //crypto library to hash the data
	//"strconv"       // for conversion
	"time" // the time for our timestamp

	"github.com/davecgh/go-spew/spew"

	//"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//"encoding/hex"
	//"encoding/json"
	//"io"
	"log"
	//"net/http"
	//"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), "", "", "bruh"}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	log.Fatal(run())

}
