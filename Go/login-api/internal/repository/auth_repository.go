package repository

import (
	"login-api/internal/model"
)

type AuthRepository interface {
	FindByUsername(username string) *model.User
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) FindByUsername(username string) *model.User {
	return &model.User{
		ID: 1,
		Username: "admin",
		Password: "$2a$10$N9qo8uLOickgx2ZMRZoMye",
		Role: "ADMIN",
	}
}