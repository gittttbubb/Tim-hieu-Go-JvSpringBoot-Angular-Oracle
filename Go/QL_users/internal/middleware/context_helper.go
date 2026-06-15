package middleware

import (
	"go-user-management/internal/model"
)

func GetClaims(
	value interface{},
) *model.JWTClaims {

	claims, ok := value.(*model.JWTClaims)
	if !ok {
		return nil
	}

	return claims
}