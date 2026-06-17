package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)

type PasswordResetService interface {
	ForgotPassword(req *dto.ForgotPasswordRequest, ipAddress *string, userAgent *string,) (string, error)
	ResetPassword(req *dto.ResetPasswordRequest,) error
	ChangePassword(userID string, req *dto.ChangePasswordRequest,) error
}

type passwordResetService struct {
	userRepo          repository.UserRepository
	passwordResetRepo repository.PasswordResetRepository
}

func NewPasswordResetService(
	userRepo repository.UserRepository,
	passwordResetRepo repository.PasswordResetRepository,
) PasswordResetService {
	return &passwordResetService{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
	}
}

func hashToken(
	token string,
) string {

	hash := sha256.Sum256(
		[]byte(token),
	)

	return hex.EncodeToString(
		hash[:],
	)
}

func (s *passwordResetService) ForgotPassword(
	req *dto.ForgotPasswordRequest,
	ipAddress *string,
	userAgent *string,
) (string, error) {

	user, err := s.userRepo.GetByUsername(
		req.Username,
	)
	if err != nil {
		return "", err
	}

	rawToken := utils.NewUUID()

	now := time.Now()

	token := &model.PasswordResetToken{
		ID:         utils.NewUUID(),
		UserID:     user.ID,
		TokenHash:  hashToken(rawToken),
		ExpiresAt:  now.Add(1 * time.Hour),
		CreatedIP:  ipAddress,
		UserAgent:  userAgent,
		CreatedAt:  now,
	}

	err = s.passwordResetRepo.Create(token)
	if err != nil {
		return "", err
	}

	return rawToken, nil
}

func (s *passwordResetService) ResetPassword(
	req *dto.ResetPasswordRequest,
) error {

	tokenHash := hashToken(
		req.Token,
	)

	token, err := s.passwordResetRepo.GetByTokenHash(
		tokenHash,
	)
	if err != nil {
		return err
	}

	if token.UsedAt != nil {
		return errors.New("reset token already used")
	}

	if token.RevokedAt != nil {
		return errors.New("reset token revoked")
	}

	if time.Now().After(token.ExpiresAt) {
		return errors.New("reset token expired")
	}

	passwordHash, err := utils.HashPassword(
		req.NewPassword,
	)
	if err != nil {
		return err
	}

	now := time.Now()

	err = s.userRepo.UpdatePassword(
		token.UserID,
		passwordHash,
		now,
		false,
	)
	if err != nil {
		return err
	}

	err = s.passwordResetRepo.MarkUsed(
		token.ID,
		now,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *passwordResetService) ChangePassword(
	userID string,
	req *dto.ChangePasswordRequest,
) error {

	user, err := s.userRepo.GetByID(
		userID,
	)
	if err != nil {
		return err
	}

	if !utils.CheckPassword(
		user.PasswordHash,
		req.OldPassword,
	) {
		return errors.New("old password is incorrect")
	}

	passwordHash, err := utils.HashPassword(
		req.NewPassword,
	)
	if err != nil {
		return err
	}

	now := time.Now()

	return s.userRepo.UpdatePassword(
		user.ID,
		passwordHash,
		now,
		false,
	)
}