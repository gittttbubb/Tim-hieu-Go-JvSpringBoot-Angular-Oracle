package service

import (
	"database/sql"
	"errors"
	"time"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type UserService interface {
	GetByID(id string,) (*dto.UserDetailResponse, error)
	List() ([]dto.UserListResponse, error)
	Create(req *dto.CreateUserRequest, createdBy string,) error
	Update(id string, req *dto.UpdateUserRequest,) error
	Delete(id string,) error
}

type userService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func mapUserListResponse(
	user *model.User,
) dto.UserListResponse {

	return dto.UserListResponse{
		ID:                 user.ID,
		TenantID:           user.TenantID,
		FullName:           user.FullName,
		Username:           user.Username,
		Email:              user.Email,
		Phone:              user.Phone,
		RoleID:             user.RoleID,
		Status:             user.Status,
		MustChangePassword: user.MustChangePassword,
	}
}

func mapUserDetailResponse(
	user *model.User,
) *dto.UserDetailResponse {

	response := &dto.UserDetailResponse{
		ID:                 user.ID,
		TenantID:           user.TenantID,
		FullName:           user.FullName,
		Username:           user.Username,
		Email:              user.Email,
		Phone:              user.Phone,
		RoleID:             user.RoleID,
		Status:             user.Status,
		MustChangePassword: user.MustChangePassword,
	}

	response.CreatedAt = user.CreatedAt.Format(time.RFC3339)
	response.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)

	return response
}

func (s *userService) GetByID(
	id string,
) (*dto.UserDetailResponse, error) {

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mapUserDetailResponse(user), nil
}

func (s *userService) List() ([]dto.UserListResponse, error) {

	users, err := s.userRepo.List()
	if err != nil {
		return nil, err
	}

	result := make(
		[]dto.UserListResponse,
		0,
		len(users),
	)

	for _, user := range users {
		result = append(
			result,
			mapUserListResponse(&user),
		)
	}

	return result, nil
}

func (s *userService) Create(
	req *dto.CreateUserRequest,
	createdBy string,
) error {

	_, err := s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	_, err = s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return err
	}
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	now := time.Now()
	createdByValue := createdBy
	user := &model.User{
		ID:           utils.NewUUID(),
		TenantID:     req.TenantID,
		FullName:     req.FullName,
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: passwordHash,
		RoleID:       req.RoleID,
		Status:             constants.UserStatusActive,
		MustChangePassword: true,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: &createdByValue,
		PasswordChangedAt: now,
	}

	return s.userRepo.Create(user)
}

func (s *userService) Update(
	id string,
	req *dto.UpdateUserRequest,
) error {

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	existingEmail, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == nil && existingEmail.ID != id {
		return errors.New("email already exists")
	}

	_, err = s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return err
	}

	user.FullName = req.FullName
	user.Email = req.Email
	user.Phone = req.Phone
	user.RoleID = req.RoleID
	user.Status = req.Status
	user.UpdatedAt = time.Now()

	return s.userRepo.Update(user)
}

func (s *userService) Delete(
	id string,
) error {

	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}
