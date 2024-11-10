package api

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func createProduct(c *gin.Context) {
	var newBook models.Product

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	*repository.GetProducts() = append(*repository.GetProducts(), newBook)
	c.JSON(http.StatusCreated, newBook)
}
