package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash)
}

func ComparePasswords(
	hashedPwd string,
	plainPwd string,
) bool {
	byteHash := []byte(hashedPwd)
	bytePlainPsw := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPsw)

	return err == nil
}
