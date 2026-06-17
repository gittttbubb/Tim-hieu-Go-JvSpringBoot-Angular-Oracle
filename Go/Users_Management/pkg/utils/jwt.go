package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   string `json:"user_id"`
	TenantID string `json:"tenant_id"`
	Username string `json:"username"`
	RoleID   string `json:"role_id"`

	jwt.RegisteredClaims
}

func GenerateToken(userID string, tenantID string, username string, roleID string, secret string, expiry time.Duration,
) (string, error) {
	claims := Claims{
		UserID:   userID,
		TenantID: tenantID,
		Username: username,
		RoleID:   roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(expiry),
			),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims,)
	return token.SignedString([]byte(secret),)
}

func ParseToken(tokenString string, secret string,) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}