package handlers

import (
	"context"
	"net/http"
	"stocksipcalculator/mappers"
	"stocksipcalculator/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type StockHandler struct {
	database *mongo.Database
	ctx      context.Context
}

func NewStockHandler(ctx context.Context, database *mongo.Database) *StockHandler {
	return &StockHandler{
		database: database,
		ctx:      ctx,
	}
}

func (handler *StockHandler) Rule(c *gin.Context) {
	var rule model.Rule

	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := mappers.ValidateReq(rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rule)
}
