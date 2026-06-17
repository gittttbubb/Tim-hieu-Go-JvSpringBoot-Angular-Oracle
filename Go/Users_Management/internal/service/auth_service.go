package service

import (
	"errors"
	"time"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type AuthService interface {
	Login(req *dto.LoginRequest, jwtSecret string, jwtExpiry time.Duration,) (*dto.LoginResponse, error)
	LoadAuthorizationData(userID string,mroleID string,) (
		[]model.RolePermission,
		[]model.UserPermissionOverride,
		error,
	)
}

type authService struct {
	authRepo repository.AuthRepository
	roleRepo repository.RoleRepository
}

func NewAuthService(
	authRepo repository.AuthRepository,
	roleRepo repository.RoleRepository,
) AuthService {
	return &authService{
		authRepo: authRepo,
		roleRepo: roleRepo,
	}
}

func (s *authService) Login(
	req *dto.LoginRequest,
	jwtSecret string,
	jwtExpiry time.Duration,
) (*dto.LoginResponse, error) {

	user, err := s.authRepo.GetUserByUsername(
		req.Username,
	)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(
		user.PasswordHash,
		req.Password,
	) {
		return nil, errors.New(
			"invalid username or password",
		)
	}

	if user.Status == constants.UserStatusLocked {
		return nil, errors.New(
			"user account is locked",
		)
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.TenantID,
		user.Username,
		user.RoleID,
		jwtSecret,
		jwtExpiry,
	)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int64(jwtExpiry.Seconds()),

		UserID:   user.ID,
		Username: user.Username,
		RoleID:   user.RoleID,

		MustChangePassword: user.MustChangePassword,
	}, nil
}

func (s *authService) LoadAuthorizationData(
	userID string,
	roleID string,
) (
	[]model.RolePermission,
	[]model.UserPermissionOverride,
	error,
) {

	rolePermissions, err := s.authRepo.GetRolePermissions(
		roleID,
	)
	if err != nil {
		return nil, nil, err
	}

	userOverrides, err := s.authRepo.GetUserPermissionOverrides(
		userID,
	)
	if err != nil {
		return nil, nil, err
	}

	return rolePermissions, userOverrides, nil
}

