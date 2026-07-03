package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/model"
	"go-rbac-system/internal/service"
	"go-rbac-system/pkg/response"
)

type AuditHandler struct {
	auditService service.AuditService
}

func NewAuditHandler(auditService service.AuditService) *AuditHandler {
	return &AuditHandler{
		auditService: auditService,
	}
}

func (h *AuditHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log, err := h.auditService.GetByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "audit log not found")
	}
	return response.Success(c, log)
}

func (h *AuditHandler) GetAll(c *fiber.Ctx) error {
	logs, err := h.auditService.GetAll()
	if err != nil {
		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}
	return response.Success(c, logs)
}

func (h *AuditHandler) ListByActor(c *fiber.Ctx) error {
	actorID := c.Params("actorId")
	logs, err := h.auditService.ListByActorID(actorID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, logs)
}

func (h *AuditHandler) ListByEntity(c *fiber.Ctx) error {
	entityType := c.Query("entityType")
	entityID := c.Query("entityId")
	logs, err := h.auditService.ListByEntity(entityType, entityID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, logs)
}

func (h *AuditHandler) ListByTenant(c *fiber.Ctx) error {
	tenantID := c.Query("tenantId")
	logs, err := h.auditService.ListByTenant(tenantID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, logs)
}

func (h *AuditHandler) List(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 20)

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	keyword := strings.TrimSpace(
		c.Query("keyword"),
	)

	logs, total, err := h.auditService.List(
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
			Items:    logs,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	)
}