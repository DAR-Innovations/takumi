package utils

import (
	"errors"
	"takumi/internal/services/authorization/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user types.User, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &types.Claims{
		ID:       user.ID,
		FullName: user.FirstName + " " + user.LastName,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
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
