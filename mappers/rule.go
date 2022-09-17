package mappers

import (
	"fmt"
	"net/mail"
	"regexp"
	"stocksipcalculator/model"
)

func ValidateReq(rule model.Rule) model.StocksError {
	var stockError model.StocksError
	for _, stockRule := range rule.StockRules {
		if !validateStockName(stockRule.StockName) {
			stockError.InvalidStocks = append(stockError.InvalidStocks, stockRule.StockName)
		}
	}

	if len(stockError.InvalidStocks) == 0 {
		return model.StocksError{
			ErrorMsg:      "invalid stock name",
			InvalidStocks: stockError.InvalidStocks,
		}
	}

	fmt.Println(rule.Email)
	fmt.Println(validateEmail(rule.Email))
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

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
