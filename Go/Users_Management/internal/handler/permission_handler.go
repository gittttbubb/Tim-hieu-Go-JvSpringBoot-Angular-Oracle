package handler

import (
	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/service"
	"go-rbac-system/pkg/response"
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

func (h *PermissionHandler) List(c *fiber.Ctx) error {

	data, err := h.permissionService.List()
	if err != nil {
		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return response.Success(c, data)
}

func (h *PermissionHandler) GetByID(c *fiber.Ctx) error {

	id := c.Params("id")

	data, err := h.permissionService.GetByID(id)
	if err != nil {
		return response.Error(
			c,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return response.Success(c, data)
}

func (h *PermissionHandler) GetByFeatureCode(c *fiber.Ctx) error {

	code := c.Params("featureCode")

	data, err := h.permissionService.GetByFeatureCode(code)
	if err != nil {
		return response.Error(
			c,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return response.Success(c, data)
}

func (h *PermissionHandler) Grouped(c *fiber.Ctx) error {

	data, err := h.permissionService.List()
	if err != nil {
		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	grouped := make(map[string][]any)

	for _, p := range data {
		grouped[p.FeatureGroup] = append(grouped[p.FeatureGroup], p)
	}

	return response.Success(c, grouped)
}