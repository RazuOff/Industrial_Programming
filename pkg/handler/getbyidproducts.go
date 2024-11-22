package handler

import (
	"net/http"

	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func GetProductsById(c *gin.Context) {
	id := c.Param("id")

	for _, item := range *repository.GetProducts() {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "prod not found"})
}
