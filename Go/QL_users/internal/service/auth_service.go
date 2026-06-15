package service

import (
	"database/sql"
	"errors"
	"strings"

	"go-user-management/internal/config"
	"go-user-management/internal/dto/auth"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
	"go-user-management/internal/utils"
)

type AuthService interface {
	Login(req dto.LoginRequest,) (*dto.LoginResponse, error)
}

type authService struct {
	authRepo           repository.AuthRepository
	permissionResolver PermissionResolverService
	cfg                *config.Config
}

func NewAuthService(authRepo repository.AuthRepository, permissionResolver PermissionResolverService, cfg *config.Config,) AuthService {
	return &authService{
		authRepo: authRepo,
		permissionResolver: permissionResolver,
		cfg: cfg,
	}
}

func (s *authService) Login(req dto.LoginRequest,) (*dto.LoginResponse, error) {
	req.Username = strings.TrimSpace(req.Username)
	if req.Username == "" {
		return nil, errors.New("username is required")
	}
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := s.authRepo.FindByUsername(req.Username,)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(
				"invalid username or password",
			)
		}
		return nil, err
	}
	if user.Status == "LOCKED" {
		return nil, errors.New(
			"account is locked",
		)
	}

	err = utils.ComparePassword(user.PasswordHash, req.Password,)
	if err != nil {
		return nil, errors.New(
			"invalid username or password",
		)
	}

	permissions, err := s.permissionResolver.ResolvePermissions(user.ID,)
	if err != nil {
		return nil, err
	}
	claims := model.JWTClaims{
		UserID:      user.ID,
		TenantID:    user.TenantID,
		Username:    user.Username,
		RoleID:      user.RoleID,
		Permissions: permissions,
	}
	token, err := utils.GenerateJWT(claims, s.cfg.JWT.Secret, s.cfg.JWT.ExpireHours,)
	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		Token:              token,
		MustChangePassword: user.MustChangePassword,
		Permissions:        permissions,
	}
	return response, nil
}