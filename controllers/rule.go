package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"stocksipcalculator/mappers"
	"stocksipcalculator/model"
	"stocksipcalculator/repository"

	"github.com/piquette/finance-go/quote"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ValidateStockName(stockRule []model.StockRule) model.StocksError {
	var stockNames []string
	var invlidStocks []string

	stockNames = mappers.GenerateStockSlice(stockRule)

	//create map from slice
	stockNamesMap := make(map[string]bool)
	for _, stockName := range stockNames {
		stockNamesMap[stockName] = false
	}

	//fetch data from yahoo finance
	iter := quote.List(stockNames)
	// Iterate over results. Will exit upon any error.
	for iter.Next() {
		q := iter.Quote()
		stockNamesMap[q.Symbol] = true
	}
	// Catch an error, if there was one.
	if iter.Err() != nil {
		// Uh-oh!
		return model.StocksError{
			ErrorMsg: "iter error",
		}
	}

	for stockName, value := range stockNamesMap {
		if !value {
			invlidStocks = append(invlidStocks, stockName)
		}
	}

	if len(invlidStocks) == 0 {
		return model.StocksError{}
	}

	return model.StocksError{
		ErrorMsg:      "invalid stock name",
		InvalidStocks: invlidStocks,
	}
}

func FetchStockLTP(stockNames []string) map[string]float64 {
	stockNamesMap := make(map[string]float64)

	//fetch data from yahoo finance
	iter := quote.List(stockNames)
	// Iterate over results. Will exit upon any error.
	for iter.Next() {
		q := iter.Quote()
		stockNamesMap[q.Symbol] = q.RegularMarketPrice
	}
	// Catch an error, if there was one.
	if iter.Err() != nil {
		// Uh-oh!
		return nil
	}
	return stockNamesMap
}

func AddRule(rules model.Rule) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}
	log.Println("Connected to MongoDB")
	database := client.Database(os.Getenv("MONGO_DATABASE"))

	repo := repository.NewRuleRepository(ctx, database)
	dbRules := getDBRules(rules)
	repo.AddRule(dbRules)

	return nil
}

func getDBRules(rules model.Rule) []interface{} {
	fmt.Println(rules)

	var dBRules []interface{}

	stockNameSlice := mappers.GenerateStockSlice(rules.StockRules)
	stockNamesMap := FetchStockLTP(stockNameSlice)

	fmt.Println(stockNameSlice)
	fmt.Println(stockNamesMap)

	for _, rule := range rules.StockRules {
		fmt.Println(rule)
		var targetPrice float64
		var upDown string
		if rule.TargetPercentage < 0 || rule.TargetPrice < 0 {
			upDown = "down"
		} else {
			upDown = "up"
		}

		if rule.TargetPercentage == 0 {
			targetPrice = rule.TargetPrice
		} else {
			LTP := stockNamesMap[mappers.GenerateStockTicker(rule.StockName)]
			targetPrice = LTP + (LTP/100)*rule.TargetPercentage
		}

		fmt.Println("ltp ------", stockNamesMap[rule.StockName])
		dbRule := model.DBRule{
			ID:          xid.New().String(),
			StockName:   rule.StockName,
			TargetPrice: targetPrice,
			UPBelow:     upDown,
			Alerted:     false,
			Email:       rules.Email,
		}
		dBRules = append(dBRules, dbRule)
	}

	return dBRules
}
