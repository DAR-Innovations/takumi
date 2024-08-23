package user

import (
	"errors"
	"takumi/internal/database"
	"takumi/internal/modules/user/types"

	"gorm.io/gorm"
)

func getUserByID(id int, handler database.DBHandler) (*types.User, error) {
	user := types.User{}
	query := types.User{ID: id}

	if err := handler.DB.First(&user, &query).Error; err == gorm.ErrRecordNotFound {
		return nil, errors.New("user is not found")
	}

	return &user, nil
}

func DeleteUserByID(id int, handler database.DBHandler) error {
	user := types.User{}
	query := types.User{ID: id}

	if err := handler.DB.First(&user, &query).Error; err != nil {
		return err
	}

	return handler.DB.Delete(&user).Error
}

func UpdateUserParams(update types.User, handler database.DBHandler) (*types.User, error) {
	user := types.User{}

	if err := handler.DB.First(&user, update.ID).Error; err != nil {
		return nil, err
	}

	if err := handler.DB.Model(&user).Updates(update).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
