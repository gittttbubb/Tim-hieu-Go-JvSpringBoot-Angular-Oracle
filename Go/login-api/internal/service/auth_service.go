package service

import (
	"errors"

	"login-api/internal/repository"
	"login-api/internal/utils"
)

type AuthService interface {
	Login(username string, password string) (string, error)
}

type authService struct {
	repo repository.AuthRepository
	secret string
}

func NewAuthService(repo repository.AuthRepository, secret string) AuthService {
	return &authService{
		repo: repo,
		secret: secret,
	}
}

func (s *authService) Login(username string, password string) (string, error) {
	user := s.repo.FindByUsername(username)

	if user == nil {
		return "", errors.New("invalid account")
	}

	return utils.GenerateToken(s.secret, user.ID)
}