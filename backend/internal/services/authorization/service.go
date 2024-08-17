package authorization

import (
	"errors"
	"net/http"
	"takumi/internal/config"
	"takumi/internal/database"
	"takumi/internal/middleware"
	"takumi/internal/services/authorization/types"
	"takumi/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Handler database.DBHandler
}

func InitAuthService(handler database.DBHandler) (*Service, error) {
	service := &Service{
		Handler: handler,
	}

	return service, nil
}

func (s *Service) Login(c *gin.Context, req *types.LoginReq) error {
	found, err := GetUserByUsernameOrEmail(req.Login, s.Handler)
	if err != nil && found == nil {
		c.AbortWithError(http.StatusNotFound, err)
		return errors.New("such user was not found")
	}

	if !utils.ComparePasswords(found.Password, req.Password) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("incorrect password"))
		return errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(*found, config.JWTSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return errors.New("apologies for inconvenience, internal error occurred")
	}

	middleware.SetCookieHandler(c, token)
	return nil
}

func (s *Service) Signup(c *gin.Context, req *types.SignupReq) error {
	found, err := GetUserByUsernameOrEmail(req.Username, s.Handler)
	if err == nil && found != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("username is taken"))
		return errors.New("username is taken")
	}

	found, err = GetUserByUsernameOrEmail(req.Email, s.Handler)
	if err == nil && found != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("user with this email already exists"))
		return errors.New("user with this email already exists")
	}

	newUser := types.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Gender:    req.Gender,
		BirthDate: req.BirthDate,
		Password:  req.Password,
	}

	created, err := CreateUser(&newUser, &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return errors.New("could not create an account")
	}

	token, err := utils.GenerateToken(*created, config.JWTSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return errors.New("apologies for inconvenience, internal error occurred")
	}

	middleware.SetCookieHandler(c, token)
	return nil
}
