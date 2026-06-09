package utils

import (
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(secret string, userId int) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
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

func ParseToken(secret string, tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// kiểm tra token gửi lên có đúng là được ký bằng thuật toán mã hóa đối xứng HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}
	// Ép kiểu dữ liệu token.Claims về dạng jwt.MapClaims để có thể truy cập các trường dữ liệu bên trong payload dưới dạng Key-Value.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid userId")
	}

	return int(userIdFloat), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost,)
	return string(bytes), err
}