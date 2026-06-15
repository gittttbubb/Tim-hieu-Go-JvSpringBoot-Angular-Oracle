package service

import (
	"go-user-management/internal/dto/permission"
	"go-user-management/internal/repository"
)

type PermissionService interface {
	GetAll() ([]dto.PermissionResponse, error)
	GetByID(id string,) (*dto.PermissionResponse, error)
	GetByFeatureCode(featureCode string,) (*dto.PermissionResponse, error)
	GetPermissionsByRole(roleID string,) ([]string, error)
	GetPermissionsByRoleWithScope(roleID string,) ([]dto.RolePermissionDetailResponse, error)
}
type permissionService struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionService(permissionRepo repository.PermissionRepository,) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

func (s *permissionService) GetAll() ([]dto.PermissionResponse, error,) {
	permissions, err := s.permissionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	result := make([]dto.PermissionResponse, 0,)
	for _, permission := range permissions {
		result = append(
			result,
			dto.PermissionResponse{
				ID:           permission.ID,
				FeatureGroup: permission.FeatureGroup,
				FeatureCode:  permission.FeatureCode,
				Action:       permission.Action,
				Description:  permission.Description,
			},
		)
	}
	return result, nil
}

func (s *permissionService) GetByID(id string,) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.PermissionResponse{
		ID:           permission.ID,
		FeatureGroup: permission.FeatureGroup,
		FeatureCode:  permission.FeatureCode,
		Action:       permission.Action,
		Description:  permission.Description,
	}, nil
}

func (s *permissionService) GetByFeatureCode(featureCode string,) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepo.GetByFeatureCode(featureCode)
	if err != nil {
		return nil, err
	}
	return &dto.PermissionResponse{
		ID:           permission.ID,
		FeatureGroup: permission.FeatureGroup,
		FeatureCode:  permission.FeatureCode,
		Action:       permission.Action,
		Description:  permission.Description,
	}, nil
}

func (s *permissionService) GetPermissionsByRole(roleID string,) ([]string, error) {
	return s.permissionRepo.GetPermissionsByRole(roleID,)
}

func (s *permissionService) GetPermissionsByRoleWithScope(roleID string,) ([]dto.RolePermissionDetailResponse, error) {
	details, err := s.permissionRepo.GetPermissionsByRoleWithScope(roleID,)
	if err != nil {
		return nil, err
	}
	result := make([]dto.RolePermissionDetailResponse, 0,)
	for _, item := range details {
		result = append(
			result,
			dto.RolePermissionDetailResponse{
				PermissionID: item.PermissionID,
				FeatureCode:  item.FeatureCode,
				Granted:      item.Granted,
				DataScope:    item.DataScope,
			},
		)
	}
	return result, nil
}