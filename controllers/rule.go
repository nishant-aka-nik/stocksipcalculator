package controllers

import (
	"bytes"
	"stocksipcalculator/model"
	"strings"

	"github.com/piquette/finance-go/quote"
)

func ValidateStockName(stockRule []model.StockRule) model.StocksError {
	var stockNames []string
	var invlidStocks []string

	//generate a slice of stock names
	for _, rule := range stockRule {
		var b bytes.Buffer
		stockName := strings.ToUpper(rule.StockName)
		b.WriteString(stockName)
		b.WriteString(".NS")
		stockNames = append(stockNames, b.String())
	}

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
