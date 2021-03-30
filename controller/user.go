package controller

import (
	"context"
	"fmt"
	db "goback/mongodb"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"fistName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
}

func Create() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := db.Client().Database("dbTest").Collection("users")
	u := User{FirstName: "John", LastName: "doe"}

	ok, err := userCollection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ok.InsertedID)
}
