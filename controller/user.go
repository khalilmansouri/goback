package controller

import (
	"context"
	db "goback/mongodb"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"fistName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
}

func Create(u User) *mongo.InsertOneResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := db.Client().Database("dbTest").Collection("users")
	// u := User{FirstName: "John", LastName: "doe"}

	ok, err := userCollection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}
	return ok
}
