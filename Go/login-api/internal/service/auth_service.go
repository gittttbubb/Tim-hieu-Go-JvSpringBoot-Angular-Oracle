package service

import (
	"errors"

	"login-api/internal/repository"
	"login-api/internal/utils"
)

type AuthService interface {
	Login(
		username string,
		password string,
	) (string, error)
}

type authService struct {
	repo repository.AuthRepository

	secret string
}

func NewAuthService(
	repo repository.AuthRepository,

	secret string,
) AuthService {

	return &authService{
		repo: repo,
		secret: secret,
	}
}

func (s *authService) Login(
	username string,
	password string,
) (string, error) {

	user,
	err :=
		s.repo.FindByUsername(
			username,
		)

	if err != nil {
		return "",
			errors.New(
				"account not found",
			)
	}

	if !utils.VerifyPassword(
		user.Password,
		password,
	) {

		return "",
			errors.New(
				"wrong password",
			)
	}

	return utils.GenerateToken(
		s.secret,
		user.ID,
	)
}