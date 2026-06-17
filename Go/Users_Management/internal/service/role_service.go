package service

import (
	"database/sql"
	"errors"

	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type RoleService interface {
	GetByID(id string) (*dto.RoleResponse, error)

	List() ([]dto.RoleResponse, error)

	Create(
		req *dto.CreateRoleRequest,
	) error

	Update(
		id string,
		req *dto.UpdateRoleRequest,
	) error

	Delete(id string) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(
	roleRepo repository.RoleRepository,
) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func mapRoleResponse(
	role *model.Role,
) *dto.RoleResponse {

	description := ""
	if role.Description != nil {
		description = *role.Description
	}

	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		DisplayName: role.DisplayName,
		Description: description,
	}
}

func (s *roleService) GetByID(
	id string,
) (*dto.RoleResponse, error) {

	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mapRoleResponse(role), nil
}

func (s *roleService) List() (
	[]dto.RoleResponse,
	error,
) {

	roles, err := s.roleRepo.List()
	if err != nil {
		return nil, err
	}

	result := make(
		[]dto.RoleResponse,
		0,
		len(roles),
	)

	for _, role := range roles {

		description := ""
		if role.Description != nil {
			description = *role.Description
		}

		result = append(
			result,
			dto.RoleResponse{
				ID:          role.ID,
				Name:        role.Name,
				DisplayName: role.DisplayName,
				Description: description,
			},
		)
	}

	return result, nil
}

func (s *roleService) Create(
	req *dto.CreateRoleRequest,
) error {

	existing, err := s.roleRepo.GetByName(req.Name)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == nil && existing != nil {
		return errors.New("role name already exists")
	}

	description := &req.Description

	role := &model.Role{
		ID:          utils.NewUUID(),
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: description,
	}

	return s.roleRepo.Create(role)
}

func (s *roleService) Update(
	id string,
	req *dto.UpdateRoleRequest,
) error {

	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		return err
	}

	role.DisplayName = req.DisplayName
	role.Description = &req.Description

	return s.roleRepo.Update(role)
}

func (s *roleService) Delete(
	id string,
) error {

	_, err := s.roleRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.roleRepo.Delete(id)
}