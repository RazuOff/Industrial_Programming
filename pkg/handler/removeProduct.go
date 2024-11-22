package handler

import (
	"net/http"

	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, book := range *repository.GetProducts() {
		if book.Id == id {
			*repository.GetProducts() = append((*repository.GetProducts())[:i], (*repository.GetProducts())[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
