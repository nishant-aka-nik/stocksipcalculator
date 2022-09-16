package main

import (
	"context"
	"log"
	"os"
	"stocksipcalculator/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var stockHandler *handlers.StockHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	database := client.Database(os.Getenv("MONGO_DATABASE"))
	stockHandler = handlers.NewStockHandler(ctx, database)
}

func main() {
	router := gin.Default()
	router.POST("/rule", stockHandler.Rule)
	router.Run()
}
