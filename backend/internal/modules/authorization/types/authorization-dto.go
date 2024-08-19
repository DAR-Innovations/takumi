package types

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignupReq struct {
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birthDate"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	jwt.StandardClaims
}

type CurrentUser struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
