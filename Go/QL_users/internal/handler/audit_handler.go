package handler

import (
	"github.com/gofiber/fiber/v2"

	dto "go-user-management/internal/dto/audit"
	"go-user-management/internal/service"
	"go-user-management/internal/utils"
)

type AuditHandler struct {
	auditService service.AuditService
}

func NewAuditHandler(
	auditService service.AuditService,
) *AuditHandler {
	return &AuditHandler{
		auditService: auditService,
	}
}

func (h *AuditHandler) Create(
	c *fiber.Ctx,
) error {

	var req dto.CreateAuditLogRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request payload",
		)
	}

	err := h.auditService.Create(req)

	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Created(
		c,
		"audit log created successfully",
	)
}

func (h *AuditHandler) GetByID(
	c *fiber.Ctx,
) error {

	id := c.Params("id")

	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	audit, err := h.auditService.GetByID(id)

	if err != nil {
		return utils.NotFound(
			c,
			"audit log not found",
		)
	}

	return utils.Success(
		c,
		audit,
	)
}

func (h *AuditHandler) List(
	c *fiber.Ctx,
) error {

	items, err := h.auditService.List()

	if err != nil {
		return utils.InternalServerError(
			c,
			err.Error(),
		)
	}

	response := dto.AuditLogListResponse{
		Items: items,
		Total: len(items),
	}

	return utils.Success(
		c,
		response,
	)
}
