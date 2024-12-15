package products

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Updates an existing product in the repository
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Updated product data"
// @Success      200      {object}  map[string]string  "Product updated successfully"
// @Failure      400      {object}  map[string]string  "Invalid request or update failed"
// @Router       /products [put]
func UpdateProduct(c *gin.Context) {

	var updatedProd models.Product

	if err := c.BindJSON(&updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := repository.UpdateProduct(updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}
