package main

import (
	"context"
	"stocksipcalculator/handlers"
	"stocksipcalculator/repository"

	"github.com/gin-gonic/gin"
)

var stockHandler *handlers.StockHandler
var ruleRepository *repository.RuleRepository

func init() {
	ctx := context.Background()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	// if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Connected to MongoDB")
	// database := client.Database(os.Getenv("MONGO_DATABASE"))
	stockHandler = handlers.NewStockHandler(ctx)
	// ruleRepository = repository.NewRuleRepository(ctx, database)
}

func main() {
	router := gin.Default()
	router.POST("/rule", stockHandler.Rule)
	router.Run()
}
