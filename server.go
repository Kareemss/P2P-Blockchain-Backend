package main

import (
	// We will need these libraries:
	// "bytes"         // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	// "crypto/sha256" //crypto library to hash the data
	// "strconv"       // for conversion
	"context"
	"time" // the time for our timestamp

	uuid "github.com/satori/go.uuid"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	// "github.com/joho/godotenv"
	// "encoding/hex"
	"crypto/sha256" //crypto library to hash the data
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("PORT")
	log.Println("Listening on ", os.Getenv("PORT"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/WriteBlock", handleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/WriteUser", HandleWriteUser).Methods("POST")
	muxRouter.HandleFunc("/UserLogin", UserLogin).Methods("POST")
	muxRouter.HandleFunc("/Market", handleGetMarket).Methods("GET")

	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	// bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// _, err = io.WriteString(w, string(bytes))
	// if err != nil {
	// 	return
	// }

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	BlockchainDatabase := connectToDb("Blockchain")
	collection := BlockchainDatabase.Collection("blocks")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	cursor.All(ctx, &Blockchain)
	defer cursor.Close(ctx)

	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.WriteString(w, string(bytes))
	if err != nil {
		return
	}
}
func handleGetMarket(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MarketDatabase := connectToDb("Market")
	Collection := MarketDatabase.Collection("Orders")

	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	cursor.All(ctx, &Market)
	defer cursor.Close(ctx)

	bytes, err := json.MarshalIndent(Market, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.WriteString(w, string(bytes))
	if err != nil {
		return
	}
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Order

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	res, _ := isBlockValid(newBlock, Blockchain[len(Blockchain)-1])
	if res {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		spew.Dump(Blockchain)

		BlockchainDatabase := connectToDb("Blockchain")
		addBlock(newBlock, BlockchainDatabase)
		//spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

func HandleWriteUser(w http.ResponseWriter, r *http.Request) {
	var NewUser User

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&NewUser); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	PasswordHash := NewUser.PasswordHash
	h := sha256.New()
	h.Write([]byte(PasswordHash))
	hashed := h.Sum(nil)
	NewUser.PasswordHash = hex.EncodeToString(hashed)

	UserDatabase := connectToDb("Users")
	AddUser(NewUser, UserDatabase)

	respondWithJSON(w, r, http.StatusCreated, NewUser)

}
func HandleWriteOrder(w http.ResponseWriter, r *http.Request) {
	var NewOrder Order

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&NewOrder); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	MarketDatabase := connectToDb("Market")
	AddOrder(NewOrder, MarketDatabase)

	respondWithJSON(w, r, http.StatusCreated, NewOrder)

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var User User

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&User); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	PasswordHash := User.PasswordHash
	h := sha256.New()
	h.Write([]byte(PasswordHash))
	hashed := h.Sum(nil)
	User.PasswordHash = hex.EncodeToString(hashed)
	res := ValidateUserLogin(User.Email, User.PasswordHash)
	sessionToken := uuid.NewV4().String()

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	respondWithJSON(w, r, http.StatusCreated, res)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
