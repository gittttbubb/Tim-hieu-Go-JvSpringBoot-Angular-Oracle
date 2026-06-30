package service

import (
	"database/sql"

	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
)

type RolePermissionService interface {
	GetByRoleID(roleID string) ([]dto.RolePermissionResponse, error)
	Assign(req *dto.AssignRolePermissionRequest) error
	Remove(roleID string, permissionID string) error
	GetEffectivePermissions(userID string, roleID string) ([]dto.UserPermission, error)
}

type rolePermissionService struct {
	rolePermissionRepo         repository.RolePermissionRepository
	roleRepo                   repository.RoleRepository
	permissionRepo             repository.PermissionRepository
	userPermissionOverrideRepo repository.UserPermissionOverrideRepository
}

func NewRolePermissionService(rolePermissionRepo repository.RolePermissionRepository, roleRepo repository.RoleRepository,
	permissionRepo repository.PermissionRepository, userPermissionOverrideRepo repository.UserPermissionOverrideRepository) RolePermissionService {
	return &rolePermissionService{
		rolePermissionRepo:         rolePermissionRepo,
		roleRepo:                   roleRepo,
		permissionRepo:             permissionRepo,
		userPermissionOverrideRepo: userPermissionOverrideRepo,
	}
}

func (s *rolePermissionService) GetByRoleID(roleID string,) ([]dto.RolePermissionResponse, error) {
	_, err := s.roleRepo.GetByID(roleID)
	if err != nil {
		return nil, err
	}
	items, err := s.rolePermissionRepo.ListByRoleID(roleID)
	if err != nil {
		return nil, err
	}
	result := make([]dto.RolePermissionResponse, 0, len(items),)
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

func (s *rolePermissionService) Assign(req *dto.AssignRolePermissionRequest,) error {
	_, err := s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return err
	}
	_, err = s.permissionRepo.GetByID(req.PermissionID)
	if err != nil {
		return err
	}
	existing, err := s.rolePermissionRepo.GetByRoleAndPermission(req.RoleID, req.PermissionID)
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

func (s *rolePermissionService) Remove(roleID string,permissionID string,) error {
	_, err := s.rolePermissionRepo.GetByRoleAndPermission(roleID,permissionID)
	if err != nil {
		return err
	}
	err = s.rolePermissionRepo.Delete(roleID,permissionID)
	if err != nil {
		return err
	}
	return nil
}

func (s *rolePermissionService) GetEffectivePermissions(userID string, roleID string,) ([]dto.UserPermission, error) {
	rolePermissions, err := s.rolePermissionRepo.ListByRoleID(roleID)
	if err != nil {
		return nil, err
	}
	userOverrides, err := s.userPermissionOverrideRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	type permissionItem struct {
		Granted   bool
		DataScope string
	}
	permissionMap := make(
		map[string]permissionItem,
	)
	// Role permissions
	for _, item := range rolePermissions {
		permissionMap[item.PermissionID] =
			permissionItem{
				Granted:   item.Granted,
				DataScope: item.DataScope,
			}
	}
	// User override
	for _, item := range userOverrides {
		permissionMap[item.PermissionID] =
			permissionItem{
				Granted:   item.Granted,
				DataScope: item.DataScope,
			}
	}
	result := make([]dto.UserPermission, 0, len(permissionMap),)
	for permissionID, permission := range permissionMap {
		if !permission.Granted {
			continue
		}
		permissionInfo, err := s.permissionRepo.GetByID(permissionID)
		if err != nil {
			continue
		}
		result = append(
			result,
			dto.UserPermission{
				PermissionID: permissionID,
				FeatureCode:  permissionInfo.FeatureCode,
				DataScope:    permission.DataScope,
			},
		)
	}
	return result, nil
}
