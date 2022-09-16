package mappers

import (
	"errors"
	"net/mail"
	"regexp"
	"stocksipcalculator/model"
)

func ValidateReq(rule model.Rule) error {
	for _, stockRule := range rule.StockRules {
		if !validateStockName(stockRule.StockName) {
			return errors.New("invalid stock name")
		}
	}
	if !validateEmail(rule.Email) {
		return errors.New("invalid email address")
	}
	return nil
}

func validateStockName(stockName string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(stockName)
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
