package service

import (
	"database/sql"

	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
)

type RolePermissionService interface {
	GetByRoleID(roleID string,) ([]dto.RolePermissionResponse, error)
	Assign(req *dto.AssignRolePermissionRequest,) error
	Remove(roleID string, permissionID string,) error
}

type rolePermissionService struct {
	rolePermissionRepo repository.RolePermissionRepository
	roleRepo           repository.RoleRepository
	permissionRepo     repository.PermissionRepository
}

func NewRolePermissionService(
	rolePermissionRepo repository.RolePermissionRepository,
	roleRepo repository.RoleRepository,
	permissionRepo repository.PermissionRepository,
) RolePermissionService {
	return &rolePermissionService{
		rolePermissionRepo: rolePermissionRepo,
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
	}
}

func (s *rolePermissionService) GetByRoleID(
	roleID string,
) ([]dto.RolePermissionResponse, error) {

	_, err := s.roleRepo.GetByID(roleID)
	if err != nil {
		return nil, err
	}

	items, err := s.rolePermissionRepo.ListByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	result := make(
		[]dto.RolePermissionResponse,
		0,
		len(items),
	)

	for _, item := range items {

		result = append(
			result,
			dto.RolePermissionResponse{
				RoleID:       item.RoleID,
				PermissionID: item.PermissionID,
				Granted:      item.Granted,
				DataScope:    item.DataScope,
			},
		)
	}

	return result, nil
}

func (s *rolePermissionService) Assign(
	req *dto.AssignRolePermissionRequest,
) error {

	_, err := s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return err
	}

	_, err = s.permissionRepo.GetByID(req.PermissionID)
	if err != nil {
		return err
	}

	existing, err := s.rolePermissionRepo.GetByRoleAndPermission(
		req.RoleID,
		req.PermissionID,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == nil && existing != nil {

		existing.Granted = req.Granted
		existing.DataScope = req.DataScope

		return s.rolePermissionRepo.Update(existing)
	}

	item := &model.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
		Granted:      req.Granted,
		DataScope:    req.DataScope,
	}

	return s.rolePermissionRepo.Assign(item)
}

func (s *rolePermissionService) Remove(
	roleID string,
	permissionID string,
) error {

	_, err := s.rolePermissionRepo.GetByRoleAndPermission(
		roleID,
		permissionID,
	)

	if err != nil {
		return err
	}

	err = s.rolePermissionRepo.Delete(
		roleID,
		permissionID,
	)

	if err != nil {
		return err
	}

	return nil
}