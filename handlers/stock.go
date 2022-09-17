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

	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, model.StocksError{
			ErrorMsg: err.Error(),
		})
		return
	}

	if invalidStocks := mappers.ValidateReq(rule); invalidStocks.InvalidStocks != nil {
		c.JSON(http.StatusBadRequest, invalidStocks)
		return
	}

	if invalidStocks := controllers.ValidateStockName(rule.StockRules); invalidStocks.InvalidStocks != nil {
		c.JSON(http.StatusBadRequest, invalidStocks)
		return
	}

	c.JSON(http.StatusOK, rule)
}
