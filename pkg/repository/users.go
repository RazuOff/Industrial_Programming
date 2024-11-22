package repository

import (
	"errors"
	"strconv"

	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/models"
)

var users = []models.User{
	{
		Id:       "1",
		Username: "Admin",
		Password: "1234",
		Role:     "Admin",
	},
	{
		Id:       "2",
		Username: "user",
		Password: "1234",
	},
}

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
