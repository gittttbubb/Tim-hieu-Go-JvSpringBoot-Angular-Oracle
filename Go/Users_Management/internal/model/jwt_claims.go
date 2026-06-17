package model

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID   string `json:"userId"`
	TenantID string `json:"tenantId"`
	Username string `json:"username"`
	RoleID   string `json:"roleId"`

	jwt.RegisteredClaims
}