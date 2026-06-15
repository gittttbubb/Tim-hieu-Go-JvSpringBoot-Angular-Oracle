package service

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"

	"go-user-management/internal/dto/role"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
)

type RoleService interface {
	CreateRole(req dto.CreateRoleRequest) error
	GetByID(id string) (*dto.RoleResponse, error)
	List() ([]dto.RoleResponse, error)
	UpdateRole(id string, req dto.UpdateRoleRequest) error
	DeleteRole(id string) error
}	
type roleService struct {
	roleRepo repository.RoleRepository
}
func NewRoleService(roleRepo repository.RoleRepository,) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (s *roleService) CreateRole(req dto.CreateRoleRequest,) error {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return errors.New("role name is required")
	}

	existingRole, err := s.roleRepo.GetByName(req.Name)
	if err == nil && existingRole != nil {
		return errors.New("role already exists")
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	role := &model.Role{
		ID:          uuid.NewString(),
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: req.Description,
	}
	return s.roleRepo.Create(role)
}

func (s *roleService) GetByID(id string,) (*dto.RoleResponse, error) {
	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		DisplayName: role.DisplayName,
		Description: role.Description,
	}, nil
}

func (s *roleService) List() ([]dto.RoleResponse, error) {
	roles, err := s.roleRepo.List()
	if err != nil {
		return nil, err
	}
	result := make([]dto.RoleResponse, 0)
	for _, role := range roles {
		result = append(result, dto.RoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			DisplayName: role.DisplayName,
			Description: role.Description,
		})
	}
	return result, nil
}

func (s *roleService) UpdateRole(id string, req dto.UpdateRoleRequest,) error {
	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		return err
	}
	role.DisplayName = req.DisplayName
	role.Description = req.Description
	return s.roleRepo.Update(role)
}

func (s *roleService) DeleteRole(id string,) error {
	_, err := s.roleRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.roleRepo.Delete(id)
}