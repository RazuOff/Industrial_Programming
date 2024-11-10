package api

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProd models.Product

	if err := c.BindJSON(&updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, item := range *repository.GetProducts() {
		if item.Id == id {
			(*repository.GetProducts())[i] = updatedProd
			c.JSON(http.StatusOK, updatedProd)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
