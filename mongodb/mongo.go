package mongo

import (
	"context"
	"log"
	"sync"
	"time"

	"goback/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var once sync.Once

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

		log.Println("Connected to mongodb ...")
	})
}

func Client() *mongo.Client {
	return client
}
