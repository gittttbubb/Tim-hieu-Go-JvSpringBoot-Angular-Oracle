package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/model"
	"go-rbac-system/internal/service"
	"go-rbac-system/pkg/response"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(permissionService service.PermissionService,) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) ListAll(c *fiber.Ctx) error {
	data, err := h.permissionService.ListAll()
	if err != nil {
		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}
	return response.Success(c, data)

}
func (h *PermissionHandler) List(c *fiber.Ctx) error {
    page := c.QueryInt("page", 1)
    pageSize := c.QueryInt("pageSize", 10)
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }
    if pageSize > 100 {
        pageSize = 100
    }
    keyword := strings.TrimSpace(
        c.Query("keyword"),
    )
    data, total, err := h.permissionService.List(
        keyword,
        page,
        pageSize,
    )
    if err != nil {
        return response.Error(
            c,
            fiber.StatusInternalServerError,
            err.Error(),
        )
    }
    return response.Success(
        c,
        model.Pagination{
            Items:    data,
            Total:    total,
            Page:     page,
            PageSize: pageSize,
        },
    )
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
	data, err := h.permissionService.ListAll()
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