package handler

import (
	dto "go-user-management/internal/dto/permission"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(
	permissionService service.PermissionService,
) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) GetAllPermissions(c *fiber.Ctx) error {
	permissions, err := h.permissionService.GetAll()
	if err != nil {
		return utils.InternalServerError(
			c,
			err.Error(),
		)
	}

	response := dto.PermissionListResponse{
		Items: permissions,
		Total: len(permissions),
	}

	return utils.Success(
		c,
		response,
	)
}

func (h *PermissionHandler) GetPermissionByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	permission, err := h.permissionService.GetByID(id)
	if err != nil {
		return utils.NotFound(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		permission,
	)
}

func (h *PermissionHandler) GetPermissionByFeatureCode(
	c *fiber.Ctx,
) error {

	featureCode := c.Params("featureCode")

	if featureCode == "" {
		return utils.BadRequest(
			c,
			"featureCode is required",
		)
	}

	permission, err := h.permissionService.GetByFeatureCode(
		featureCode,
	)
	if err != nil {
		return utils.NotFound(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		permission,
	)
}

func (h *PermissionHandler) GetPermissionsByRole(
	c *fiber.Ctx,
) error {

	roleID := c.Params("roleId")

	if roleID == "" {
		return utils.BadRequest(
			c,
			"roleId is required",
		)
	}

	permissions, err := h.permissionService.GetPermissionsByRole(
		roleID,
	)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		permissions,
	)
}

func (h *PermissionHandler) GetPermissionsByRoleWithScope(
	c *fiber.Ctx,
) error {

	roleID := c.Params("roleId")

	if roleID == "" {
		return utils.BadRequest(
			c,
			"roleId is required",
		)
	}

	permissions, err := h.permissionService.GetPermissionsByRoleWithScope(
		roleID,
	)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		permissions,
	)
}