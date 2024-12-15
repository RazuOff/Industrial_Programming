package products

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary      Create a product
// @Description  Creating product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body   models.Product  true  "Add product"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  map[string]string  "Invalid request or error message"
// @Router       /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := repository.AddProducts([]models.Product{newProduct}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}
