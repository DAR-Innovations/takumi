package user

import (
	"errors"
	"net/http"
	"takumi/internal/database"
	"takumi/internal/modules/user/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	Handler database.DBHandler
}

func InitUserService(handler database.DBHandler) (*Service, error) {
	service := &Service{
		Handler: handler,
	}
	return service, nil
}

func (s *Service) GetUserByID(c *gin.Context, userID int) (*types.User, error) {
	user, err := getUserByID(userID, s.Handler)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, errors.New("user not found"))
			return nil, errors.New("user not found")
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, errors.New("failed to retrieve user")
	}
	return user, nil
}

func (s *Service) DeleteUserByID(c *gin.Context, userID int) error {
	if err := DeleteUserByID(userID, s.Handler); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, errors.New("user not found"))
			return errors.New("user not found")
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return errors.New("failed to delete user")
	}
	return nil
}

func (s *Service) UpdateUserParams(c *gin.Context, update types.User) (*types.User, error) {
	user, err := UpdateUserParams(update, s.Handler)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithError(http.StatusNotFound, errors.New("user not found"))
			return nil, errors.New("user not found")
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, errors.New("failed to update user")
	}
	return user, nil
}

func (s *Service) UploadProfilePicture(c *gin.Context, userID int, imageData []byte) (*types.ProfilePicture, error) {
	picture := types.ProfilePicture{
		UserID:    userID,
		ImageData: imageData,
	}

	savedPicture, err := UpsertProfilePicture(&picture, s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, errors.New("could not upload profile picture")
	}

	return savedPicture, nil
}

func (s *Service) GetProfilePicture(c *gin.Context, userID int) (*types.ProfilePicture, error) {
	picture, err := GetProfilePictureByUserID(userID, s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return nil, errors.New("profile picture not found")
	}

	return picture, nil
}

func (s *Service) DeleteProfilePicture(c *gin.Context, userID int) error {
	err := DeleteProfilePicture(userID, s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return errors.New("could not delete profile picture")
	}
	return nil
}
