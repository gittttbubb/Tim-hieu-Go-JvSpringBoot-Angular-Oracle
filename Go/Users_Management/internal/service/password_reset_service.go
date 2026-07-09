package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/internal/validator"
	"go-rbac-system/pkg/utils"
)

type PasswordResetService interface {
	ForgotPassword(req *dto.ForgotPasswordRequest, ipAddress *string, userAgent *string,) error
	AdminResetPassword(userID string, actorID string, actor string,) (string, error)
	ChangePassword(userID string, req *dto.ChangePasswordRequest,) error
    ResetPassword(req *dto.ResetPasswordRequest,) error
}

type passwordResetService struct {
	userRepo          repository.UserRepository
	passwordResetRepo repository.PasswordResetRepository
	auditRepo repository.AuditRepository
    emailService EmailService
    frontendURL string
}

func NewPasswordResetService(userRepo repository.UserRepository, passwordResetRepo repository.PasswordResetRepository,
	auditRepo repository.AuditRepository, emailService EmailService, frontendURL string) PasswordResetService {
	return &passwordResetService{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
		auditRepo: auditRepo,
        emailService: emailService,
        frontendURL: frontendURL,
	}
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (s *passwordResetService) ForgotPassword(req *dto.ForgotPasswordRequest, ipAddress *string, userAgent *string) error {
    user, err := s.userRepo.GetByEmail(req.Email)
    if err != nil {
        // tránh user enumeration
        return nil
    }
    rawToken, err := utils.GenerateResetToken()
    if err != nil {
        return err
    }
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
        return err
    }
    resetLink := fmt.Sprintf(
        "%s/reset-password?token=%s",
        s.frontendURL,
        rawToken,
    )
    body := fmt.Sprintf(`
        <h2>Yêu cầu đặt lại mật khẩu</h2>
        <p>Xin chào %s,</p>
        <p>Chúng tôi nhận được yêu cầu đặt lại mật khẩu cho tài khoản của bạn.</p>
        <p><a href="%s">Đặt lại mật khẩu</a></p>
        <p>Liên kết này sẽ hết hạn trong vòng 1 giờ.</p>
        <p>Nếu bạn không yêu cầu đặt lại mật khẩu, bạn có thể bỏ qua email này.</p>
    `,
        user.Username,
        resetLink,
    )
    err = s.emailService.Send(
        user.Email,
        "Yêu cầu đặt lại mật khẩu",
        body,
    )
    if err != nil {
        return err
    }
    return nil
}

func (s *passwordResetService) ResetPassword(req *dto.ResetPasswordRequest,) error {
	if req.NewPassword != req.ConfirmPassword {
		return errors.New("password confirmation does not match")
	}
	tokenHash := hashToken(req.Token)
	token, err := s.passwordResetRepo.GetByTokenHash(tokenHash)
	if err != nil {
		return errors.New("invalid token")
	}
	if token.UsedAt != nil {
		return errors.New("token already used")
	}
	if token.RevokedAt != nil {
		return errors.New("token revoked")
	}
	if time.Now().After(token.ExpiresAt) {
		return errors.New("token expired")
	}
	user, err := s.userRepo.GetByID(token.UserID)
	if err != nil {
		return err
	}
	// validate password
	if err := validator.ValidatePassword(req.NewPassword); err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	now := time.Now()
	err = s.userRepo.UpdatePassword(
		user.ID,
		hashedPassword,
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

func (s *passwordResetService) AdminResetPassword(userID string, actorID string, actor string,) (string, error) {
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
        Actor:          actor,
		ActorID:        &actorID,
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
    if err := validator.ValidatePassword(req.NewPassword); err != nil {
        return err
    }
    // optional: enforce first login logic
    if user.MustChangePassword == false {
        // normal flow
        if !utils.CheckPassword(user.PasswordHash, req.OldPassword) {
            return errors.New("Old password is incorrect")
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

    return nil
}