package main

import (
	"net/http"
	"stocksipcalculator/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var recipies []model.Recipe

func init() {
	recipies = make([]model.Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe model.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = uuid.NewString()
	recipe.PublishedAt = time.Now().UTC()

	recipies = append(recipies, recipe)
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run()
}
