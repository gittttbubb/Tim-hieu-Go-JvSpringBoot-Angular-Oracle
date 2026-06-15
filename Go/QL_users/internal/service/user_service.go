package service

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"go-user-management/internal/dto/user"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest, createdBy string) error
	GetByID(id string) (*dto.UserResponse, error)
	List() ([]dto.UserResponse, error)
	UpdateUser(id string, req dto.UpdateUserRequest) error
	DeleteUser(id string) error
}

type userService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository,) UserService {
	return &userService{userRepo: userRepo, roleRepo: roleRepo,}
}

func (s *userService) CreateUser(req dto.CreateUserRequest, createdBy string,) error {
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)
	if req.Username == "" {
		return errors.New("username is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	if req.RoleID == "" {
		return errors.New("role id is required")
	}

	existingUser, err := s.userRepo.GetByUsername(req.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	existingEmail, err := s.userRepo.GetByEmail(req.Email)
	if err == nil && existingEmail != nil {
		return errors.New("email already exists")
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	_, err = s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return errors.New("role not found")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	user := &model.User{
		ID:                 uuid.NewString(),
		TenantID:           req.TenantID,
		FullName:           req.FullName,
		Username:           req.Username,
		Email:              req.Email,
		Phone:              req.Phone,
		PasswordHash:       string(hash),
		RoleID:             req.RoleID,
		Status:             "ACTIVE",
		MustChangePassword: req.MustChangePassword,
		CreatedBy:          createdBy,
	}
	return s.userRepo.Create(user)
}

func (s *userService) GetByID(id string,) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:                 user.ID,
		TenantID:           user.TenantID,
		FullName:           user.FullName,
		Username:           user.Username,
		Email:              user.Email,
		Phone:              user.Phone,
		RoleID:             user.RoleID,
		Status:             user.Status,
		MustChangePassword: user.MustChangePassword,
		CreatedAt:          user.CreatedAt,
		UpdatedAt:          user.UpdatedAt,
		PasswordChangedAt:  user.PasswordChangedAt,
	}, nil
}

func (s *userService) List() ([]dto.UserResponse, error) {
	users, err := s.userRepo.List()
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserResponse, 0)
	for _, user := range users {
		result = append(result, dto.UserResponse{
			ID:                 user.ID,
			TenantID:           user.TenantID,
			FullName:           user.FullName,
			Username:           user.Username,
			Email:              user.Email,
			Phone:              user.Phone,
			RoleID:             user.RoleID,
			Status:             user.Status,
			MustChangePassword: user.MustChangePassword,
			CreatedAt:          user.CreatedAt,
			UpdatedAt:          user.UpdatedAt,
			PasswordChangedAt:  user.PasswordChangedAt,
		})
	}
	return result, nil
}

func (s *userService) UpdateUser(id string, req dto.UpdateUserRequest,) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	_, err = s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return errors.New("role not found")
	}
	user.FullName = req.FullName
	user.Email = req.Email
	user.Phone = req.Phone
	user.RoleID = req.RoleID
	user.Status = req.Status
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id string) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.userRepo.Delete(id)
}