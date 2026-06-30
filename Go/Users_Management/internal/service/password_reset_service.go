package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type PasswordResetService interface {
	ForgotPassword(req *dto.ForgotPasswordRequest, ipAddress *string, userAgent *string,) (string, error)
	AdminResetPassword(userID string, adminID string) (string, error)
	ChangePassword(userID string, req *dto.ChangePasswordRequest,) error
}

type passwordResetService struct {
	userRepo          repository.UserRepository
	passwordResetRepo repository.PasswordResetRepository
	auditRepo repository.AuditRepository
}

func NewPasswordResetService(userRepo repository.UserRepository, passwordResetRepo repository.PasswordResetRepository,
	auditRepo repository.AuditRepository) PasswordResetService {
	return &passwordResetService{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
		auditRepo: auditRepo,
	}
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (s *passwordResetService) ForgotPassword(req *dto.ForgotPasswordRequest, ipAddress *string, userAgent *string,) (string, error) {
    user, err := s.userRepo.GetByUsername(req.Username)
    if err != nil {
        return "", nil // tránh user enumeration
    }
    rawToken := utils.NewUUID()
    now := time.Now()
    token := &model.PasswordResetToken{
        ID:        utils.NewUUID(),
        UserID:    user.ID,
        TokenHash: hashToken(rawToken),
        ExpiresAt: now.Add(1 * time.Hour),
        CreatedIP: ipAddress,
        UserAgent: userAgent,
        CreatedAt: now,
    }
    err = s.passwordResetRepo.Create(token)
    if err != nil {
        return "", err
    }

	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		TargetUserID:   &user.ID,
		Action:         "PASSWORD_RESET_REQUEST",
		EntityType:     "USER",
		EntityID:       user.ID,
		IPAddress:      ipAddress,
		EventTimestamp: time.Now(),
	}
	_ = s.auditRepo.Create(audit)

    // production: gửi email, KHÔNG return token
    return "", nil
}

func (s *passwordResetService) AdminResetPassword(userID string, adminID string) (string, error) {
    user, err := s.userRepo.GetByID(userID)
    if err != nil {
        return "", err
    }
    tempPassword := utils.GenerateTempPassword()
    passwordHash, err := utils.HashPassword(tempPassword)
    if err != nil {
        return "", err
    }
    now := time.Now()
    err = s.userRepo.UpdatePassword(
        user.ID,
        passwordHash,
        now,
        false,
    )
    if err != nil {
        return "", err
    }
    // enforce first login change
    err = s.userRepo.UpdateMustChangePassword(user.ID, true)
    if err != nil {
        return "", err
    }
	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		ActorID:        &adminID,
		TargetUserID:   &user.ID,
		Action:         "ADMIN_RESET_PASSWORD",
		EntityType:     "USER",
		EntityID:       user.ID,
		Reason:         utils.StringPtr("password reset by administrator"),
		EventTimestamp: time.Now(),
	}
	_ = s.auditRepo.Create(audit)

    return tempPassword, nil
}

func (s *passwordResetService) ChangePassword(userID string, req *dto.ChangePasswordRequest) error {
    user, err := s.userRepo.GetByID(userID)
    if err != nil {
        return err
    }
    // optional: enforce first login logic
    if user.MustChangePassword == false {
        // normal flow
        if !utils.CheckPassword(user.PasswordHash, req.OldPassword) {
            return errors.New("old password is incorrect")
        }
    }
    passwordHash, err := utils.HashPassword(req.NewPassword)
    if err != nil {
        return err
    }
    now := time.Now()
    err = s.userRepo.UpdatePassword(user.ID, passwordHash, now, false)
    if err != nil {
        return err
    }
	err = s.userRepo.UpdateStatus(user.ID, constants.UserStatusActive)
	if err != nil {
		return err
	}
    // clear flag after first login
    if user.MustChangePassword {
        _ = s.userRepo.UpdateMustChangePassword(user.ID, false)
    }

	audit := &model.AuditLog{
		ID:             utils.NewUUID(),
		TenantID:       user.TenantID,
		ActorID:        &user.ID,
		TargetUserID:   &user.ID,
		Action:         "CHANGE_PASSWORD",
		EntityType:     "USER",
		EntityID:       user.ID,
		EventTimestamp: time.Now(),
	}
	_ = s.auditRepo.Create(audit)

    return nil
}