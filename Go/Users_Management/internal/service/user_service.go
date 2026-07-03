package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type UserService interface {
	GetByID(id string,) (*dto.UserDetailResponse, error)
	List(keyword string, page int, pageSize int,) ([]dto.UserListResponse, int64, error)
	Create(req *dto.CreateUserRequest, actorID string, actor string,) (string, error)
	Update(id string, req *dto.UpdateUserRequest,actorID string, actor string,) error
	Delete(id string,actorID string, actor string,) error
	LockUser(id string,actorID string, actor string,) error
	UnlockUser(id string,actorID string, actor string,) error
	UpdateRole(userID string, roleID string,actorID string, actor string,) error
}

type userService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
	auditRepo repository.AuditRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository,
	auditRepo repository.AuditRepository) UserService {
	return &userService{
		userRepo: userRepo,
		roleRepo: roleRepo,
		auditRepo: auditRepo,
	}
}

func mapUserListResponse(user *model.User,) dto.UserListResponse {
	return dto.UserListResponse{
		ID:                 user.ID,
		TenantID:           user.TenantID,
		FullName:           user.FullName,
		Username:           user.Username,
		Email:              user.Email,
		Phone:              user.Phone,
		RoleID:             user.RoleID,
		Status:             user.Status,
		MustChangePassword: user.MustChangePassword,
	}
}

func mapUserDetailResponse(user *model.User) *dto.UserDetailResponse {
	response := &dto.UserDetailResponse{
		ID:                 user.ID,
		TenantID:           user.TenantID,
		FullName:           user.FullName,
		Username:           user.Username,
		Email:              user.Email,
		Phone:              user.Phone,
		RoleID:             user.RoleID,
		Status:             user.Status,
		MustChangePassword: user.MustChangePassword,
	}
	response.CreatedAt = user.CreatedAt.Format(time.RFC3339)
	response.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)
	return response
}

func (s *userService) GetByID(id string) (*dto.UserDetailResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return mapUserDetailResponse(user), nil
}

func (s *userService) List(keyword string, page int, pageSize int,) ([]dto.UserListResponse, int64, error) {
	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.List(keyword,offset,pageSize,)
	if err != nil {
		return nil, 0, err
	}
	result := make([]dto.UserListResponse, 0, len(users))
	for _, user := range users {
		result = append(result, mapUserListResponse(&user))
	}
	return result, total, nil
}

func (s *userService) Create(req *dto.CreateUserRequest, actorID string, actor string,) (string, error) {
	_, err := s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return "", errors.New("username already exists")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	_, err = s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return "", errors.New("email already exists")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	_, err = s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return "", err
	}
	tempPassword := utils.GenerateTempPassword()
	passwordHash, err := utils.HashPassword(tempPassword)
	if err != nil {
		return "", err
	}
	now := time.Now()
	user := &model.User{
		ID:                 utils.NewUUID(),
		TenantID:           req.TenantID,
		FullName:           req.FullName,
		Username:           req.Username,
		Email:              req.Email,
		Phone:              req.Phone,
		PasswordHash:       passwordHash,
		RoleID:             req.RoleID,
		Status:             constants.UserStatusPendingPasswordChange,
		MustChangePassword: true,
		CreatedAt:          now,
		UpdatedAt:          now,
		CreatedBy:          &actorID,
		PasswordChangedAt:  now,
	}
	err = s.userRepo.Create(user)
	if err != nil {
		return "", err
	}
	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		Actor:          actor,
    	ActorID:        &actorID,
		Action:         "USER_CREATE",
		EntityType:     "USER",
		EntityID:       user.ID,
		TargetUserID:   &user.ID,
		EventTimestamp: time.Now(),
		AfterData: utils.StringPtr(
			fmt.Sprintf(
				`{"username":"%s","email":"%s","roleId":"%s","status":"%s"}`,
				user.Username,
				user.Email,
				user.RoleID,
				user.Status,
			),
		),
	}
	// _ = s.auditRepo.Create(audit)
	err = s.auditRepo.Create(audit)
	if err != nil {
		fmt.Println("AUDIT ERROR:", err)
	}
	return tempPassword, nil
}

func (s *userService) Update(id string, req *dto.UpdateUserRequest, actorID string, actor string,) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	existingEmail, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil && existingEmail.ID != id {
		return errors.New("email already exists")
	}
	_, err = s.roleRepo.GetByID(req.RoleID)
	if err != nil {
		return err
	}
	before := fmt.Sprintf(
		`{"fullName":"%s","userName":"%s","email":"%s","roleId":"%s","status":"%s"}`, user.FullName, user.Username, user.Email, user.RoleID, user.Status,
	)

	user.FullName = req.FullName
	user.Username = req.Username
	user.Email = req.Email
	user.Phone = req.Phone
	user.RoleID = req.RoleID
	user.Status = req.Status
	user.UpdatedAt = time.Now()

	after := fmt.Sprintf(
		`{"fullName":"%s","userName":"%s","email":"%s","roleId":"%s","status":"%s"}`, req.FullName, req.Username, req.Email, req.RoleID, req.Status,
	)

	 err = s.userRepo.Update(user)
    if err != nil {
        return err
    }

	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		Actor:          actor,
		ActorID:        &actorID,
		TargetUserID:   &user.ID,
		Action:         "USER_UPDATE",
		EntityType:     "USER",
		EntityID:       user.ID,
		BeforeData:     &before,
		AfterData:      &after,
		EventTimestamp: time.Now(),
	}
	// _ = s.auditRepo.Create(audit)
	err = s.auditRepo.Create(audit)
	if err != nil {
		fmt.Println("AUDIT ERROR:", err)
	}

	return nil
}

func (s *userService) Delete(id string, actorID string, actor string,) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	before := fmt.Sprintf(
		`{"username":"%s","email":"%s","roleId":"%s"}`, user.Username, user.Email, user.RoleID,
	)
	err = s.userRepo.Delete(id)
    if err != nil {
        return err
    }
	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		Actor:          actor,
		ActorID:        &actorID,
		TargetUserID:   &user.ID,
		Action:         "USER_DELETE",
		EntityType:     "USER",
		EntityID:       user.ID,
		BeforeData:     &before,
		EventTimestamp: time.Now(),
	}
	// _ = s.auditRepo.Create(audit)
	err = s.auditRepo.Create(audit)
	if err != nil {
		fmt.Println("AUDIT ERROR:", err)
	}

	return nil
}

func (s *userService) LockUser(id string, actorID string, actor string,) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	err = s.userRepo.UpdateStatus(
        id,
        constants.UserStatusLocked,
    )
    if err != nil {
        return err
    }
	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		Actor:          actor,
		ActorID:        &actorID,
		TargetUserID:   &user.ID,
		Action:         "USER_LOCK",
		EntityType:     "USER",
		EntityID:       user.ID,
		Reason:         utils.StringPtr("manual lock"),
		EventTimestamp: time.Now(),
	}
	// _ = s.auditRepo.Create(audit)
	err = s.auditRepo.Create(audit)
	if err != nil {
		fmt.Println("AUDIT ERROR:", err)
	}

	return nil
}

func (s *userService) UnlockUser(id string, actorID string, actor string,) error {
    user, err := s.userRepo.GetByID(id)
    if err != nil {
        return err
    }
    status := constants.UserStatusActive
    if user.MustChangePassword {
        status = constants.UserStatusPendingPasswordChange
    }
    err = s.userRepo.UpdateStatus(
        id,
        status,
    )
    if err != nil {
        return err
    }

    audit := &model.AuditLog{
        ID:             utils.NewUUID(),
        TenantID:       user.TenantID,
		Actor:          actor,
		ActorID:        &actorID,
        TargetUserID:   &user.ID,
        Action:         "USER_UNLOCK",
        EntityType:     "USER",
        EntityID:       user.ID,
        EventTimestamp: time.Now(),
    }
    // _ = s.auditRepo.Create(audit)
	err = s.auditRepo.Create(audit)
	if err != nil {
		fmt.Println("AUDIT ERROR:", err)
	}

    return nil
}

func (s *userService) UpdateRole(userID string, roleID string, actorID string, actor string,) error {
	// kiểm tra user tồn tại
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	// kiểm tra role tồn tại
	_, err = s.roleRepo.GetByID(roleID)
	if err != nil {
		return err
	}
	// tránh update vô nghĩa
	if user.RoleID == roleID {
		return errors.New("user already has this role")
	}

	oldRoleID := user.RoleID
	err = s.userRepo.UpdateRole(
        userID,
        roleID,
    )
    if err != nil {
        return err
    }

	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		Actor:          actor,
		ActorID:        &actorID,
		TargetUserID:   &user.ID,
		Action:         "USER_ROLE_CHANGE",
		EntityType:     "USER",
		EntityID:       user.ID,
		BeforeData:     utils.StringPtr(oldRoleID),
		AfterData:      utils.StringPtr(roleID),
		EventTimestamp: time.Now(),
	}
	// _ = s.auditRepo.Create(audit)
	err = s.auditRepo.Create(audit)
	if err != nil {
		fmt.Println("AUDIT ERROR:", err)
	}

	return nil
}