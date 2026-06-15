package handler

import (
	dto "go-user-management/internal/dto/role"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(
	roleService service.RoleService,
) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var req dto.CreateRoleRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}

	err := h.roleService.CreateRole(req)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Created(
		c,
		"role created successfully",
	)
}

func (h *RoleHandler) GetRoleByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	role, err := h.roleService.GetByID(id)
	if err != nil {
		return utils.NotFound(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		role,
	)
}

func (h *RoleHandler) ListRoles(c *fiber.Ctx) error {
	roles, err := h.roleService.List()
	if err != nil {
		return utils.InternalServerError(
			c,
			err.Error(),
		)
	}
	response := dto.RoleListResponse{
		Items: roles,
		Total: len(roles),
	}
	return utils.Success(
		c,
		response,
	)
}

func (h *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	var req dto.UpdateRoleRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}

	err := h.roleService.UpdateRole(
		id,
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
		"role updated successfully",
	)
}

func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	err := h.roleService.DeleteRole(id)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.OK(
		c,
		"role deleted successfully",
	)
}