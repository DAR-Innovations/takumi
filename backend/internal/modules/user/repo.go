package user

import (
	"errors"
	"log"
	"takumi/internal/database"
	"takumi/internal/modules/user/types"
	"takumi/pkg/utils"

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

func UpdateUserProfilePicture(userID int, profilePictureURL string, handler database.DBHandler) (*types.User, error) {
	user := types.User{}
	tx := handler.DB.Begin()

	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if user.ProfilePicture != "" && user.ProfilePicture != "default.jpg" {
		if err := utils.DeleteFile("profile-pictures", user.ProfilePicture); err != nil {
			log.Printf("Failed to delete: %s, error: %v", user.ProfilePicture, err)
			tx.Rollback()
			return nil, err
		}
	}

	user.ProfilePicture = profilePictureURL
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}

func GetProfilePictureByUserID(userID int, handler database.DBHandler) (string, error) {
	user := types.User{}

	if err := handler.DB.Select("profile_picture").First(&user, userID).Error; err != nil {
		return "", err
	}

	if user.ProfilePicture == "" {
		return "", errors.New("no profile picture found for this user")
	}

	return user.ProfilePicture, nil
}

func DeleteUserProfilePicture(userID int, handler database.DBHandler) (*types.User, error) {
	user := types.User{}
	if err := handler.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	user.ProfilePicture = ""
	if err := handler.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
