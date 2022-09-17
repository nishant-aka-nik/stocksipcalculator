package handlers

import (
	"context"
	"net/http"
	"stocksipcalculator/controllers"
	"stocksipcalculator/mappers"
	"stocksipcalculator/model"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	ctx context.Context
}

func NewStockHandler(ctx context.Context) *StockHandler {
	return &StockHandler{
		ctx: ctx,
	}
}

func (handler *StockHandler) Rule(c *gin.Context) {
	var rule model.Rule
	var invalidStocks model.StocksError

	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, model.StocksError{
			ErrorMsg: err.Error(),
		})
		return
	}

	invalidStocks = mappers.ValidateReq(rule)
	if invalidStocks.ErrorMsg != "" {
		c.JSON(http.StatusBadRequest, invalidStocks)
		return
	}

	invalidStocks = controllers.ValidateStockName(rule.StockRules)
	if invalidStocks.ErrorMsg != "" {
		c.JSON(http.StatusBadRequest, invalidStocks)
		return
	}

	err := controllers.AddRule(rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, rule)
}
