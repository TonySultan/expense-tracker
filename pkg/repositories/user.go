package repositories

import (
	"fmt"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/TonySultan/expense-tracker/pkg/utils"
)

var Users = make(map[string]models.User)

func CreateUser(login string, password string) (int, error) {
	user := models.User{
		Login:    login,
		Password: password,
	}

	if !utils.CheckPassword(password) {
		return 0, fmt.Errorf("password birnarse")
	}

	Users[user.Login] = user

	return len(Users), nil
}
