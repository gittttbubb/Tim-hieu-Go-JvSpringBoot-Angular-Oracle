package service

import (
	"database/sql"
	"errors"
	"strings"

	"go-user-management/internal/dto/rolepermission"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
)

type RolePermissionService interface {
	AssignPermission(req dto.AssignPermissionRequest,) error
	RevokePermission(req dto.RevokePermissionRequest,) error
	GetByRoleID(roleID string,) ([]dto.RolePermissionResponse, error)
}

type rolePermissionService struct {
	rolePermissionRepo repository.RolePermissionRepository
	roleRepo           repository.RoleRepository
	permissionRepo     repository.PermissionRepository
}

func NewRolePermissionService(rolePermissionRepo repository.RolePermissionRepository, 
	roleRepo repository.RoleRepository,
	permissionRepo repository.PermissionRepository,
) RolePermissionService {
	return &rolePermissionService{
		rolePermissionRepo: rolePermissionRepo,
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
	}
}

func (s *rolePermissionService) AssignPermission(req dto.AssignPermissionRequest,) error {
	req.RoleID = strings.TrimSpace(req.RoleID)
	req.PermissionID = strings.TrimSpace(req.PermissionID)
	req.DataScope = strings.ToUpper(
		strings.TrimSpace(req.DataScope),
	)

	if req.RoleID == "" {
		return errors.New("role id is required")
	}
	if req.PermissionID == "" {
		return errors.New("permission id is required")
	}
	if req.DataScope == "" {
		return errors.New("data scope is required")
	}

	switch req.DataScope {
	case "OWN", "ALL":
	default:
		return errors.New("invalid data scope")
	}

	_, err := s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return errors.New("role not found")
	}
	_, err = s.permissionRepo.GetByID(
		req.PermissionID,
	)
	if err != nil {
		return errors.New("permission not found")
	}

	existing, err := s.rolePermissionRepo.GetByRoleAndPermission(req.RoleID, req.PermissionID,)
	if err == nil && existing != nil {
		return errors.New(
			"permission already assigned to role",
		)
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	rp := &model.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
		Granted:      true,
		DataScope:    req.DataScope,
	}
	return s.rolePermissionRepo.Create(rp)
}

func (s *rolePermissionService) RevokePermission(req dto.RevokePermissionRequest,) error {
	_, err := s.rolePermissionRepo.GetByRoleAndPermission(req.RoleID, req.PermissionID,)
	if err != nil {
		return err
	}
	return s.rolePermissionRepo.Delete(req.RoleID, req.PermissionID,)
}

func (s *rolePermissionService) GetByRoleID(roleID string,) ([]dto.RolePermissionResponse, error) {
	rolePermissions,err := s.rolePermissionRepo.GetByRoleID(roleID,)
	if err != nil {
		return nil, err
	}
	result := make([]dto.RolePermissionResponse, 0,)
	for _, rp := range rolePermissions {
		result = append(
			result,
			dto.RolePermissionResponse{
				RoleID:       rp.RoleID,
				PermissionID: rp.PermissionID,
				Granted:      rp.Granted,
				DataScope:    rp.DataScope,
			},
		)
	}
	return result, nil
}