package repository

import (
	"ginexample.com/pkg/db/postgre"
	"ginexample.com/pkg/models"
)

func GetUserByLogin(login string) (models.User, error) {
	var user models.User
	err := postgre.DB.Where("username = ?", login).First(&user).Error
	return user, err
}
func GetUsers() ([]models.User, error) {
	var users []models.User
	err := postgre.DB.Find(&users).Error
	return users, err
}
func AddUser(login string, password string) error {

	return postgre.DB.Create(&models.User{Username: login, Password: password}).Error
}
