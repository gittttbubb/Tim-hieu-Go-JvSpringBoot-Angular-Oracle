package service

import (

	"go-user-management/internal/model"
	"go-user-management/internal/repository"
)

type PermissionResolverService interface {
	ResolvePermissions(userID string,) ([]model.EffectivePermission, error)
}

type permissionResolverService struct {
	userRepo       repository.UserRepository
	permissionRepo repository.PermissionRepository
	overrideRepo   repository.UserPermissionOverrideRepository
}

func NewPermissionResolverService(
	userRepo repository.UserRepository,
	permissionRepo repository.PermissionRepository,
	overrideRepo repository.UserPermissionOverrideRepository,
) PermissionResolverService {

	return &permissionResolverService{
		userRepo:       userRepo,
		permissionRepo: permissionRepo,
		overrideRepo:   overrideRepo,
	}
}

func (s *permissionResolverService) ResolvePermissions(userID string,) ([]model.EffectivePermission, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	rolePermissions, err := s.permissionRepo.GetPermissionsByRoleWithScope(user.RoleID,)
	if err != nil {
		return nil, err
	}
	
	permissionMap :=make(map[string]model.EffectivePermission,)
	for _, rp := range rolePermissions {
		if !rp.Granted {
			continue
		}
		permissionMap[rp.FeatureCode] = model.EffectivePermission{
			FeatureCode: rp.FeatureCode, 
			DataScope: rp.DataScope,
		}
	}

	overrides, err := s.overrideRepo.GetByUserID(userID,)
	if err != nil {
		return nil, err
	}
	for _, override := range overrides {
		permission, err := s.permissionRepo.GetByID(override.PermissionID,)
		if err != nil {
			return nil, err
		}
		featureCode := permission.FeatureCode
		if !override.Granted {
			delete(permissionMap, featureCode,)
			continue
		}
		permissionMap[featureCode] =
			model.EffectivePermission{
				FeatureCode: featureCode,
				DataScope:   override.DataScope,
			}
	}

	result := make([]model.EffectivePermission, 0, len(permissionMap),)
	for _, permission := range permissionMap {
		result = append(result, permission,)
	}
	return result, nil
}