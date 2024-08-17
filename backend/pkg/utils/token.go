package utils

import (
	"errors"
	"takumi/internal/services/authorization/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user types.User, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &types.Claims{
		ID:       user.UserID,
		Email:    user.Email,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}

func ParseToken(tokenString string, jwtKey string) (*types.Claims, error) {
	claims := &types.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if err != nil {
		return nil, err
	}

	return claims, nil
}
