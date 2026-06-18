package service

import (
	"database/sql"
	"time"

	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type UserPermissionOverrideService interface {
	GetByID(id string,) (*dto.UserPermissionOverrideResponse, error)
	GetByUserID(userID string,) ([]dto.UserPermissionOverrideResponse, error)
	Assign(req *dto.UserPermissionOverrideRequest, createdBy string,) error
	Delete(id string,) error
}

type userPermissionOverrideService struct {
	overrideRepo   repository.UserPermissionOverrideRepository
	userRepo       repository.UserRepository
	permissionRepo repository.PermissionRepository
}

func NewUserPermissionOverrideService(
	overrideRepo repository.UserPermissionOverrideRepository,
	userRepo repository.UserRepository,
	permissionRepo repository.PermissionRepository,
) UserPermissionOverrideService {
	return &userPermissionOverrideService{
		overrideRepo:   overrideRepo,
		userRepo:       userRepo,
		permissionRepo: permissionRepo,
	}
}

func mapUserPermissionOverrideResponse(
	item *model.UserPermissionOverride,
) *dto.UserPermissionOverrideResponse {

	reason := ""
	if item.Reason != nil {
		reason = *item.Reason
	}

	return &dto.UserPermissionOverrideResponse{
		ID:           item.ID,
		UserID:       item.UserID,
		PermissionID: item.PermissionID,
		Granted:      item.Granted,
		Reason:       reason,
		CreatedBy:    item.CreatedBy,
		DataScope:    item.DataScope,
	}
}

func (s *userPermissionOverrideService) GetByID(
	id string,
) (*dto.UserPermissionOverrideResponse, error) {

	item, err := s.overrideRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mapUserPermissionOverrideResponse(item), nil
}

func (s *userPermissionOverrideService) GetByUserID(
	userID string,
) ([]dto.UserPermissionOverrideResponse, error) {
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	items, err := s.overrideRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	result := make(
		[]dto.UserPermissionOverrideResponse,
		0,
		len(items),
	)

	for _, item := range items {

		reason := ""
		if item.Reason != nil {
			reason = *item.Reason
		}

		result = append(
			result,
			dto.UserPermissionOverrideResponse{
				ID:           item.ID,
				UserID:       item.UserID,
				PermissionID: item.PermissionID,
				Granted:      item.Granted,
				Reason:       reason,
				CreatedBy:    item.CreatedBy,
				DataScope:    item.DataScope,
			},
		)
	}

	return result, nil
}

func (s *userPermissionOverrideService) Assign(
	req *dto.UserPermissionOverrideRequest,
	createdBy string,
) error {

	_, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return err
	}

	_, err = s.permissionRepo.GetByID(req.PermissionID)
	if err != nil {
		return err
	}

	existing, err := s.overrideRepo.GetByUserAndPermission(
		req.UserID,
		req.PermissionID,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	reason := req.Reason

	if err == nil && existing != nil {

		existing.Granted = req.Granted
		existing.Reason = &reason
		existing.DataScope = req.DataScope

		return s.overrideRepo.Update(existing)
	}

	item := &model.UserPermissionOverride{
		ID:           utils.NewUUID(),
		UserID:       req.UserID,
		PermissionID: req.PermissionID,
		Granted:      req.Granted,
		Reason:       &reason,
		CreatedBy:    createdBy,
		CreatedAt:    time.Now(),
		DataScope:    req.DataScope,
	}

	return s.overrideRepo.Create(item)
}

func (s *userPermissionOverrideService) Delete(
	id string,
) error {

	_, err := s.overrideRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.overrideRepo.Delete(id)
}