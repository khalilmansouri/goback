package mongo

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"goback/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var once sync.Once

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"fistName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
}

func Init() {
	once.Do(func() {
		var err error = nil
		client, err = mongo.NewClient(options.Client().ApplyURI(config.Get().DB_URI))
		if err != nil {
			log.Fatal(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(client)
		fmt.Println("Connected to mongodb ...")
	})
	// u := User{FirstName: "John", LastName: "doe"}
	// fmt.Println(u)
	// userCollection := client.Database("dbTest").Collection("users")
	// ok, err := userCollection.InsertOne(ctx, u)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(ok.InsertedID)
}

func Client() *mongo.Client {
	fmt.Println("CONNECTION")
	// cl, err := mongo.NewClient(options.Client().ApplyURI(config.Get().DB_URI))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// err = cl.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = cl.Ping(ctx, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return client
}
