package repository

import (
	"errors"
	"strconv"

	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/models"
)

func GetUserByLogin(login string) (models.User, error) {
	for _, user := range users {
		if user.Username == login {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
func GetAllUsers() []models.User {
	return users
}

func AddUser(login string, password string) {
	users = append(users, models.User{Id: strconv.Itoa(len(users) + 1), Username: login, Password: password, Role: auth.USER_ROLE})
}
