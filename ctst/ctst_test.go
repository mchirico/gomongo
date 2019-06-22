package ctst

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"testing"
	"time"
)

func Test_SimpleConnect(t *testing.T) {

	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	dbName := os.Getenv("MONGO_DATABASE")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())

	collection := client.Database(dbName).Collection("numbers")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		// do something with result....
		fmt.Printf("id: %v\nResult: %v\n", id, result)

		if result["name"] != "pi" {
			t.FailNow()
		}

		if result["value"] != 3.14159 {
			t.FailNow()
		}

		for k, v := range result {
			fmt.Printf("k: %v, v: %v\n", k, v)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	collection.Drop(ctx)
}

func TestExample(t *testing.T) {
	Example()
}
