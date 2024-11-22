package handler

import (
	"net/http"

	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

func Registrate(c *gin.Context) {
	var inputForm auth.Credentials

	if err := c.BindJSON(&inputForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input error"})
		return
	}

	if _, err := repository.GetUserByLogin(inputForm.Username); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage": "user already exists"})
		return
	}
	repository.AddUser(inputForm.Username, inputForm.Password)

	c.JSON(http.StatusCreated, gin.H{"message": "account created"})
}
