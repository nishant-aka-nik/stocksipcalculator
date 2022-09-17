package mappers

import (
	"bytes"
	"net/mail"
	"regexp"
	"stocksipcalculator/model"
	"strings"
)

func ValidateReq(rule model.Rule) model.StocksError {
	var stockError model.StocksError
	for _, stockRule := range rule.StockRules {
		if !validateStockName(stockRule.StockName) {
			stockError.InvalidStocks = append(stockError.InvalidStocks, stockRule.StockName)
		}
	}

	if len(stockError.InvalidStocks) != 0 {
		return model.StocksError{
			ErrorMsg:      "invalid stock name",
			InvalidStocks: stockError.InvalidStocks,
		}
	}

	if !validateEmail(rule.Email) {
		return model.StocksError{
			ErrorMsg: "invalid email address",
		}
	}

	return model.StocksError{}
}

func validateStockName(stockName string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(stockName)
}

//TODO: fix email validation make it more robust
func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func GenerateStockSlice(stockRule []model.StockRule) []string {
	var stockNames []string
	//generate a slice of stock names
	for _, rule := range stockRule {
		var b bytes.Buffer
		stockName := strings.ToUpper(rule.StockName)
		b.WriteString(stockName)
		b.WriteString(".NS")
		stockNames = append(stockNames, b.String())
	}
	return stockNames
}

func GenerateStockTicker(stockName string) string {
	var b bytes.Buffer
	stockName = strings.ToUpper(stockName)
	b.WriteString(stockName)
	b.WriteString(".NS")
	return b.String()
}
