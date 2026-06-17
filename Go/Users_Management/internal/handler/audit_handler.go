package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

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

func (h *AuditHandler) ListWithFilter(c *fiber.Ctx) error {
	tenantID := c.Query("tenantId")

	if tenantID == "" {
		return response.Error(c, fiber.StatusBadRequest, "tenantId is required")
	}

	actorID := c.Query("actorId")
	entityType := c.Query("entityType")
	entityID := c.Query("entityId")

	fromStr := c.Query("from")
	toStr := c.Query("to")

	var from, to time.Time
	var err error

	if fromStr != "" {
		from, err = time.Parse(time.RFC3339, fromStr)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "invalid from date")
		}
	}

	if toStr != "" {
		to, err = time.Parse(time.RFC3339, toStr)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "invalid to date")
		}
	}

	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	logs, err := h.auditService.ListWithFilter(
		tenantID,
		actorID,
		entityType,
		entityID,
		from,
		to,
		limit,
		offset,
	)

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, logs)
}