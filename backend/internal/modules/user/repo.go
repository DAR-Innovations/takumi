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

func UpsertProfilePicture(picture *types.ProfilePicture, handler database.DBHandler) (*types.ProfilePicture, error) {
	existing := types.ProfilePicture{}
	err := handler.DB.Where("user_id = ?", picture.UserID).First(&existing).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existing.ID != 0 {
		existing.ImageData = picture.ImageData
		err = handler.DB.Save(&existing).Error
		if err != nil {
			return nil, err
		}
		return &existing, nil
	} else {
		err = handler.DB.Create(picture).Error
		if err != nil {
			return nil, err
		}
		return picture, nil
	}
}

func GetProfilePictureByUserID(userID int, handler database.DBHandler) (*types.ProfilePicture, error) {
	var picture types.ProfilePicture
	err := handler.DB.Where("user_id = ?", userID).First(&picture).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("profile picture not found")
		}
		return nil, err
	}
	return &picture, nil
}

func DeleteProfilePicture(userID int, handler database.DBHandler) error {
	err := handler.DB.Where("user_id = ?", userID).Delete(&types.ProfilePicture{}).Error
	if err != nil {
		return err
	}
	return nil
}
