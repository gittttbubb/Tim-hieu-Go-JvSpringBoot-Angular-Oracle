package handler

import (

	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/service"
	"go-rbac-system/internal/validator"
	"go-rbac-system/pkg/response"
)

type RoleHandler struct {
	roleService           service.RoleService
	rolePermissionService service.RolePermissionService
}

func NewRoleHandler(
	roleService service.RoleService,
	rolePermissionService service.RolePermissionService,
) *RoleHandler {
	return &RoleHandler{
		roleService:           roleService,
		rolePermissionService: rolePermissionService,
	}
}
func (h *RoleHandler) List(c *fiber.Ctx) error {
	data, err := h.roleService.List()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, data)
}

func (h *RoleHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := h.roleService.GetByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}

	return response.Success(c, data)
}

func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateRoleRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	err := h.roleService.Create(&req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "role created",
	})
}

func (h *RoleHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateRoleRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	err := h.roleService.Update(id, &req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "role updated",
	})
}

func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.roleService.Delete(id)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "role deleted",
	})
}

func (h *RoleHandler) GetPermissions(c *fiber.Ctx) error {
	roleID := c.Params("id")

	data, err := h.rolePermissionService.GetByRoleID(roleID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, data)
}

func (h *RoleHandler) AssignPermission(c *fiber.Ctx) error {
	
	var req dto.AssignRolePermissionRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	if err := validator.Validate.Struct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	err := h.rolePermissionService.Assign(&req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "permission assigned to role",
	})
}

func (h *RoleHandler) RemovePermission(c *fiber.Ctx) error {
	roleID := c.Params("id")
	permissionID := c.Params("permissionId")

	err := h.rolePermissionService.Remove(roleID, permissionID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "permission removed from role",
	})
}

