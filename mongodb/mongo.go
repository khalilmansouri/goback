package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"goback/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Get().DB_URI))
	log.Println(config.Get().DB_URI)
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

	fmt.Println("Connected to mongodb ...")

}

func Client() *mongo.Client {
	return client
}
