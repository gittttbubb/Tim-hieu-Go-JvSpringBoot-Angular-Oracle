package service

import ( 
	"database/sql"
	"errors"
	"strings"
	"github.com/google/uuid"

	"go-user-management/internal/dto/user_override"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
)

type UserPermissionOverrideService interface {
	GrantOverride(req dto.GrantOverrideRequest,createdBy string,) error
	RevokeOverride(req dto.RevokeOverrideRequest,) error
	GetByUserID(userID string,) ([]dto.UserPermissionOverrideResponse, error)
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

func (s *userPermissionOverrideService) GrantOverride(req dto.GrantOverrideRequest, createdBy string,) error {
	req.UserID = strings.TrimSpace(req.UserID)
	req.PermissionID = strings.TrimSpace(req.PermissionID)
	req.DataScope = strings.ToUpper(
		strings.TrimSpace(req.DataScope),
	)
	if req.UserID == "" {
		return errors.New("user id is required")
	}
	if req.PermissionID == "" {
		return errors.New("permission id is required")
	}
	switch req.DataScope {
	case "OWN", "ALL":
	default:
		return errors.New("invalid data scope")
	}

	_, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return errors.New("user not found")
	}
	_, err = s.permissionRepo.GetByID(
		req.PermissionID,
	)
	if err != nil {
		return errors.New("permission not found")
	}

	existing, err := s.overrideRepo.GetByUserAndPermission(req.UserID, req.PermissionID,)
	if err == nil && existing != nil {
		return errors.New(
			"override already exists",
		)
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	override := &model.UserPermissionOverride{
		ID:           uuid.NewString(),
		UserID:       req.UserID,
		PermissionID: req.PermissionID,
		Granted:      req.Granted,
		Reason:       req.Reason,
		CreatedBy:    createdBy,
		DataScope:    req.DataScope,
	}
	return s.overrideRepo.Create(override,)
}

func (s *userPermissionOverrideService) RevokeOverride(req dto.RevokeOverrideRequest,) error {
	_, err := s.overrideRepo.GetByUserAndPermission(req.UserID, req.PermissionID,)
	if err != nil {
		return err
	}
	return s.overrideRepo.Delete(
		req.UserID,
		req.PermissionID,
	)
}

func (s *userPermissionOverrideService) GetByUserID(userID string,) ([]dto.UserPermissionOverrideResponse, error) {
	items, err := s.overrideRepo.GetByUserID(userID,)
	if err != nil {
		return nil, err
	}
	result :=make([]dto.UserPermissionOverrideResponse,0,)
	for _, item := range items {
		result = append(
			result,
			dto.UserPermissionOverrideResponse{
				ID:           item.ID,
				UserID:       item.UserID,
				PermissionID: item.PermissionID,
				Granted:      item.Granted,
				Reason:       item.Reason,
				CreatedBy:    item.CreatedBy,
				DataScope:    item.DataScope,
			},
		)
	}
	return result, nil
}