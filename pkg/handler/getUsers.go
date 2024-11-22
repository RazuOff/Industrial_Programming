package handler

import (
	"net/http"

	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := repository.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
