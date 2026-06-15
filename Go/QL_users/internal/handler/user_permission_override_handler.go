package handler

import (
	"github.com/gofiber/fiber/v2"

	dto "go-user-management/internal/dto/user_override"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"
	"go-user-management/internal/middleware"
)

type UserPermissionOverrideHandler struct {
	overrideService service.UserPermissionOverrideService
}

func NewUserPermissionOverrideHandler(
	overrideService service.UserPermissionOverrideService,
) *UserPermissionOverrideHandler {
	return &UserPermissionOverrideHandler{
		overrideService: overrideService,
	}
}

func (h *UserPermissionOverrideHandler) GrantOverride(
	c *fiber.Ctx,
) error {

	var req dto.GrantOverrideRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request payload",
		)
	}

	claims := middleware.GetClaims(
		c.Locals("claims"),
	)

	if claims == nil {
		return utils.Unauthorized(
			c,
			"invalid token",
		)
	}

	err := h.overrideService.GrantOverride(
		req,
		claims.UserID,
	)

	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Created(
		c,
		"user permission override granted successfully",
	)
}

func (h *UserPermissionOverrideHandler) RevokeOverride(
	c *fiber.Ctx,
) error {

	var req dto.RevokeOverrideRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request payload",
		)
	}

	err := h.overrideService.RevokeOverride(
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
		"user permission override revoked successfully",
	)
}

func (h *UserPermissionOverrideHandler) GetByUserID(
	c *fiber.Ctx,
) error {

	userID := c.Params("userId")

	if userID == "" {
		return utils.BadRequest(
			c,
			"user id is required",
		)
	}

	items, err := h.overrideService.GetByUserID(
		userID,
	)

	if err != nil {
		return utils.InternalServerError(
			c,
			err.Error(),
		)
	}

	response := dto.UserPermissionOverrideListResponse{
		Items: items,
		Total: len(items),
	}

	return utils.Success(
		c,
		response,
	)
}