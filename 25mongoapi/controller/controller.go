package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectioString = ""
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectioString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB connection secure")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready.")
}
