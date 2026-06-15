package handler

import (
	"go-user-management/internal/dto/auth"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(
	authService service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}

	response, err := h.authService.Login(req)
	if err != nil {
		return utils.Unauthorized(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		response,
	)
}