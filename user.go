package main //Import the main package
import (
	"context"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	FullName             string  `bson:"name"`
	ID                   string  `bson:"id,omitempty"`
	PhoneNumber          int     `bson:"phone"`
	Email                string  `bson:"email"`
	UserName             string  `bson:"user-name,omitempty"`
	SmartMeterNumber     int     `bson:"smart-meter-number,omitempty"`
	PasswordHash         string  `bson:"passowrd-hash,omitempty"`
	Address              string  `bson:"address,omitempty"`
	EnergyBalance        float32 `bson:"energy-balance,omitempty"`
	CurrencyBalance      float32 `bson:"currency-balance,omitempty"`
	CompletedTransaction int     `bson:"completed-transactions"`
}

func GetUser(FieldName int, value string) (User, bool) {
	var result bool
	var Field string
	if FieldName == 1 {
		Field = "email"
	} else if FieldName == 2 {
		Field = "user-name"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	UserDatabase := connectToDb("Users")
	UserCollection := UserDatabase.Collection("Users")
	filterCursor, err := UserCollection.Find(ctx, bson.M{Field: value})
	if err != nil {
		log.Fatal(err)
		result = false
	}
	var Profiles []User
	if err = filterCursor.All(ctx, &Profiles); err != nil {
		log.Fatal(err)
		result = false
	}
	Profile := Profiles[0]
	return Profile, result
}

func ValidateUserLogin(Email string, PasswordHash string) bool {
	Profile, result := GetUser(1, Email)
	if Email == Profile.Email && PasswordHash == Profile.PasswordHash {
		result = true
	} else {
		result = false
	}

	return result
}
