package service

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"go-user-management/internal/dto/auth"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
	"go-user-management/internal/utils"
)

type PasswordResetService interface {
	ForgotPassword(req dto.ForgotPasswordRequest, ip string, userAgent string,) (string, error)
	ResetPassword(req dto.ResetPasswordRequest,) error
}

type passwordResetService struct {
	authRepo          repository.AuthRepository
	userRepo          repository.UserRepository
	passwordResetRepo repository.PasswordResetRepository
}

func NewPasswordResetService(authRepo repository.AuthRepository, userRepo repository.UserRepository, passwordResetRepo repository.PasswordResetRepository,
) PasswordResetService {
	return &passwordResetService{
		authRepo:          authRepo,
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
	}
}

func generateResetToken() (string, error) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func hashResetToken(token string) string {

	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}

func (s *passwordResetService) ForgotPassword(req dto.ForgotPasswordRequest, ip string, userAgent string,) (string, error) {
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		return "", errors.New("email is required")
	}

	user, err := s.authRepo.GetByEmail(req.Email,)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	rawToken, err := generateResetToken()
	if err != nil {
		return "", err
	}
	tokenHash := hashResetToken(rawToken,)
	resetToken := &model.PasswordResetToken{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(
			1 * time.Hour,
		),
		CreatedIP: ip,
		UserAgent: userAgent,
	}
	err = s.passwordResetRepo.Create(
		resetToken,
	)
	if err != nil {
		return "", err
	}
	return rawToken, nil
}

func (s *passwordResetService) ResetPassword(req dto.ResetPasswordRequest,) error {
	if req.Token == "" {
		return errors.New("token is required")
	}
	if req.NewPassword == "" {
		return errors.New("new password is required")
	}
	if len(req.NewPassword) < 8 {
		return errors.New(
			"password must be at least 8 characters",
		)
	}
	tokenHash := hashResetToken(req.Token,)
	token, err := s.passwordResetRepo.GetByTokenHash(tokenHash,)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid reset token",)
		}
		return err
	}
	if token.RevokedAt != nil {
		return errors.New("reset token has been revoked",)
	}
	if token.UsedAt != nil {
		return errors.New("reset token has already been used",)
	}
	if time.Now().After(token.ExpiresAt) {
		_ = s.passwordResetRepo.Revoke(token.ID,)
		return errors.New("reset token has expired",)
	}
	passwordHash, err := utils.HashPassword(req.NewPassword,)
	if err != nil {
		return err
	}
	err = s.userRepo.UpdatePassword(token.UserID, passwordHash,)
	if err != nil {
		return err
	}

	err = s.passwordResetRepo.MarkUsed(token.ID,)
	if err != nil {
		return err
	}
	return nil
}
