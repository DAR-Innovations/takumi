package authorization

import (
	"errors"
	"regexp"
	"takumi/internal/database"
	"takumi/internal/modules/authorization/types"
	"takumi/pkg/utils"
	"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

func GetUserByUsernameOrEmail(login string, handler database.DBHandler) (*types.User, error) {
	user := &types.User{}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)

	var query types.User
	if emailRegex.MatchString(login) {
		query = types.User{Email: login}
	} else {
		query = types.User{Username: login}
	}
	err := handler.DB.First(user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user is not found")
	}

	return user, nil
}

func CreateUser(registration *types.User, handler *database.DBHandler) (*types.User, error) {
	if len(registration.Password) < 6 {
		return nil, errors.New("minimum password length is 6")
	}

	password := utils.HashPassword([]byte(registration.Password))
	err := checkmail.ValidateFormat(registration.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}

	user := &types.User{
		Username:  registration.Username,
		FirstName: registration.FirstName,
		LastName:  registration.LastName,
		Email:     registration.Email,
		Password:  password,
		Gender:    registration.Gender,
		BirthDate: registration.BirthDate,
		Role:      "USER",
		CreatedAt: time.Now(),
		Coins:     0,
	}

	err = handler.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
