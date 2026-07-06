package utils

import (
	"crypto/rand"
	"math/big"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost,)
	return string(hash), err
}

func CheckPassword(hash string, password string,) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password),) == nil
}

const passwordChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*"
func GenerateTempPassword() string {
	length := 8
	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(
			rand.Reader,
			big.NewInt(int64(len(passwordChars))),
		)
		if err != nil {
			panic(err)
		}
		password[i] = passwordChars[n.Int64()]
	}

	return string(password)
}