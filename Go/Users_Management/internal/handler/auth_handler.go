package handler

import (
	"time"
	v "go-rbac-system/internal/validator"
	"go-rbac-system/internal/config"
	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/service"
	"go-rbac-system/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
	passwordResetService service.PasswordResetService
	cfg         *config.Config
}

func NewAuthHandler(authService service.AuthService, passwordResetService service.PasswordResetService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		passwordResetService: passwordResetService,
		cfg:         cfg,
	}
}
func (h *AuthHandler) validateStruct(s interface{}) error {
	return v.Validate.Struct(s)
}
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if err := h.validateStruct(&req); err != nil {
		return response.ValidationError(c, err)
	}
	res, err := h.authService.Login(
		&req,
		h.cfg.JWT.Secret,
		time.Duration(h.cfg.JWT.AccessExpiryMinutes)*time.Second,
	)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, err.Error())
	}
	return response.Success(c, res)
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	var req dto.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	userID := c.Locals(constants.ContextUserID)
	if userID == nil {
		return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}
	err := h.passwordResetService.ChangePassword(userID.(string), &req,)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.Map{"message": "password changed successfully",})
}

func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
    var req dto.ForgotPasswordRequest
    if err := c.BodyParser(&req); err != nil {
        return response.Error(
            c,
            fiber.StatusBadRequest,
            "invalid request body",
        )
    }
    ip := c.IP()
    ua := c.Get("User-Agent")
    err := h.passwordResetService.ForgotPassword(
        &req,
        &ip,
        &ua,
    )
    if err != nil {
        return response.Error(
            c,
            fiber.StatusInternalServerError,
            err.Error(),
        )
    }
    return response.Success(c, fiber.Map{
        "message": "Email khôi phục mật khẩu đã được gửi.",
    })
}

func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req dto.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}
	err := h.passwordResetService.ResetPassword(&req)
	if err != nil {
		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}
	return response.Success(c, fiber.Map{
		"message": "password reset successfully",
	})
}

func (h *AuthHandler) AdminResetPassword(c *fiber.Ctx) error {
    userID := c.Params("id")
    actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)
    tempPassword, err := h.passwordResetService.AdminResetPassword(userID, actorID, actor)
    if err != nil {
        return response.Error(c, fiber.StatusBadRequest, err.Error())
    }
    return response.Success(c, fiber.Map{
        "message": "password reset successfully",
        "temporaryPassword": tempPassword,
    })
}