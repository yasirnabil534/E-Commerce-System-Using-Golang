package databases

import (
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://yasirnabil534:1234567890@node-training-cluster.vdd97td.mongodb.net/goEcommerce?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeOut(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil);
	if err != nil {
		log.Println("Failed to connect to mongoDB server")
	}

	fmt.Println("Connected to mongoDB server successfully")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
} 

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}