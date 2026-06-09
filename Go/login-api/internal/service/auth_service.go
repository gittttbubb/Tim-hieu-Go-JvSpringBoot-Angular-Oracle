package service

import (
	"errors"

	"login-api/internal/model"
	"login-api/internal/repository"
	"login-api/internal/utils"
)

type AuthService interface {
	Login(username string,password string,) (string, error)
	GetUserByID(id int) (*model.User, error)
	Register(username string, password string) error
}

type authService struct {
	repo repository.AuthRepository
	secret string
}

func NewAuthService(repo repository.AuthRepository,secret string,) AuthService {
	return &authService{repo: repo, secret: secret,}
}
// thuộc về con trỏ của authService
func (s *authService) Login(username string, password string,) (string, error) {

	user,err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("account not found")
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("wrong password",)
	}
	return utils.GenerateToken(s.secret, user.ID)
}
func (s *authService) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *authService) Register(username string, password string) error {
	_, err := s.repo.FindByUsername(username)
	if err == nil {
		return errors.New("username already exists")
	}
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &model.User{
		Username: username,
		Password: hash,
		Role:     "USER",
	}
	return s.repo.Register(user)
}