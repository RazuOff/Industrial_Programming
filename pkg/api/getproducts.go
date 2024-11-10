package api

import (
	"net/http"

	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetProducts())
}
