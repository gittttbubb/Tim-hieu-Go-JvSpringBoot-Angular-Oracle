package handler

import (
	dto "go-user-management/internal/dto/rolepermission"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type RolePermissionHandler struct {
	rolePermissionService service.RolePermissionService
}

func NewRolePermissionHandler(
	rolePermissionService service.RolePermissionService,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		rolePermissionService: rolePermissionService,
	}
}

func (h *RolePermissionHandler) AssignPermission(
	c *fiber.Ctx,
) error {

	var req dto.AssignPermissionRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}

	err := h.rolePermissionService.AssignPermission(
		req,
	)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Created(
		c,
		"permission assigned successfully",
	)
}

func (h *RolePermissionHandler) RevokePermission(
	c *fiber.Ctx,
) error {

	var req dto.RevokePermissionRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}

	err := h.rolePermissionService.RevokePermission(
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
		"permission revoked successfully",
	)
}

func (h *RolePermissionHandler) GetByRoleID(
	c *fiber.Ctx,
) error {

	roleID := c.Params("roleId")

	if roleID == "" {
		return utils.BadRequest(
			c,
			"roleId is required",
		)
	}

	rolePermissions, err := h.rolePermissionService.GetByRoleID(
		roleID,
	)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	response := dto.RolePermissionListResponse{
		Items: rolePermissions,
		Total: len(rolePermissions),
	}

	return utils.Success(
		c,
		response,
	)
}