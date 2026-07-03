package service

import (
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
)

type PermissionService interface {
	GetByID(id string) (*dto.PermissionResponse, error)
	GetByFeatureCode(featureCode string,) (*dto.PermissionResponse, error)
	List(keyword string, page int, pageSize int,) ([]dto.PermissionResponse, int64, error)
	ListAll() ([]dto.PermissionResponse, error)
}

type permissionService struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionService(permissionRepo repository.PermissionRepository) PermissionService {
	return &permissionService{permissionRepo: permissionRepo}
}

func mapPermissionResponse(permission *model.Permission,) *dto.PermissionResponse {
	description := ""
	if permission.Description != nil {
		description = *permission.Description
	}
	return &dto.PermissionResponse{
		ID:           permission.ID,
		FeatureGroup: permission.FeatureGroup,
		FeatureCode:  permission.FeatureCode,
		Action:       permission.Action,
		Description:  description,
	}
}

func (s *permissionService) GetByID(id string,) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return mapPermissionResponse(permission), nil
}

func (s *permissionService) GetByFeatureCode(featureCode string,) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepo.GetByFeatureCode(featureCode)
	if err != nil {
		return nil, err
	}
	return mapPermissionResponse(permission), nil
}

func (s *permissionService) List(keyword string, page int, pageSize int,) ([]dto.PermissionResponse, int64, error) {
    offset := (page - 1) * pageSize
    permissions, total, err := s.permissionRepo.List(
        keyword,
        offset,
        pageSize,
    )
    if err != nil {
        return nil, 0, err
    }
    result := make([]dto.PermissionResponse, 0, len(permissions))
    for _, permission := range permissions {
        description := ""
        if permission.Description != nil {
            description = *permission.Description
        }
        result = append(result, dto.PermissionResponse{
            ID:           permission.ID,
            FeatureGroup: permission.FeatureGroup,
            FeatureCode:  permission.FeatureCode,
            Action:       permission.Action,
            Description:  description,
        })
    }
    return result, total, nil
}

func (s *permissionService) ListAll() ([]dto.PermissionResponse, error) {
	permissions, err := s.permissionRepo.ListAll()
	if err != nil {
		return nil, err
	}
	result := make(
		[]dto.PermissionResponse,
		0,
		len(permissions),
	)
	for _, permission := range permissions {
		description := ""
		if permission.Description != nil {
			description = *permission.Description
		}
		result = append(
			result,
			dto.PermissionResponse{
				ID:           permission.ID,
				FeatureGroup: permission.FeatureGroup,
				FeatureCode:  permission.FeatureCode,
				Action:       permission.Action,
				Description:  description,
			},
		)
	}
	return result, nil
}