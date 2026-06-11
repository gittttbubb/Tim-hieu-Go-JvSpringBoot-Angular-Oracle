package utils

import (
	"errors"
	"login-api/internal/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(secret string, userId int, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"role": role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func VerifyPassword(hash string, password string,) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password),)
	
	// trả về true nếu password đúng, ngược lại trả về false
	return err == nil
}

func ParseToken(secret string, tokenString string) (*dto.TokenInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("invalid userId")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role")
	}

	return &dto.TokenInfo{
		UserID: int(userIdFloat),
		Role:   role,
	}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost,)
	return string(bytes), err
}