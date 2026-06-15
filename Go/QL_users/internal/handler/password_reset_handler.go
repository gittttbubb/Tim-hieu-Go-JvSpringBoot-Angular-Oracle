package handler

import (
	"github.com/gofiber/fiber/v2"

	dto "go-user-management/internal/dto/auth"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"
)

type PasswordResetHandler struct {
	passwordResetService service.PasswordResetService
}

func NewPasswordResetHandler(
	passwordResetService service.PasswordResetService,
) *PasswordResetHandler {
	return &PasswordResetHandler{
		passwordResetService: passwordResetService,
	}
}

func (h *PasswordResetHandler) ForgotPassword(
	c *fiber.Ctx,
) error {

	var req dto.ForgotPasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request payload",
		)
	}

	ip := c.IP()

	userAgent := c.Get(
		"User-Agent",
	)

	token, err := h.passwordResetService.ForgotPassword(
		req,
		ip,
		userAgent,
	)

	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		fiber.Map{
			"resetToken": token,
		},
	)
}

func (h *PasswordResetHandler) ResetPassword(
	c *fiber.Ctx,
) error {

	var req dto.ResetPasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request payload",
		)
	}

	err := h.passwordResetService.ResetPassword(
		req,
	)

	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.OK(
		c,
		"password reset successfully",
	)
}
