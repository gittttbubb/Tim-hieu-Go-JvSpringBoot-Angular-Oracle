package service

import (
	"time"

	"go-rbac-system/internal/model"
	"go-rbac-system/internal/repository"
	"go-rbac-system/pkg/utils"
)
type AuditContext struct {
	UserID   string
	TenantID string
	Actor    string
	Method   string
	Path     string
	IP       string
}

type AuditService interface {
	Log(action string, ctx *AuditContext) error
	GetByID(id string,) (*model.AuditLog, error)
	ListByActorID(actorID string,) ([]model.AuditLog, error)
	ListByEntity(entityType string, entityID string,) ([]model.AuditLog, error)
	ListByTenant(tenantID string,) ([]model.AuditLog, error)
	ListWithFilter(
	tenantID string,
	actorID string,
	entityType string,
	entityID string,
	from time.Time,
	to time.Time,
	limit int,
	offset int,
) ([]model.AuditLog, error)
}

type auditService struct {
	auditRepo repository.AuditRepository
}

func NewAuditService(auditRepo repository.AuditRepository,) AuditService {
	return &auditService{
		auditRepo: auditRepo,
	}
}

func (s *auditService) Log(action string, ctx *AuditContext) error {
	audit := &model.AuditLog{
		ID:       utils.NewUUID(),
		TenantID: ctx.TenantID,
		Actor:    ctx.Actor,
		Action:   normalizeAction(action),
		EventTimestamp: time.Now(),
		ActorID: utils.StringPtr(ctx.UserID),
		IPAddress: utils.StringPtr(ctx.IP),
		// optional (Phase 9 có thể mở rộng)
		EntityType: "",
		EntityID:   "",
	}

	return s.auditRepo.Create(audit)
}

func (s *auditService) GetByID(
	id string,
) (*model.AuditLog, error) {

	return s.auditRepo.GetByID(id)
}

func (s *auditService) ListByActorID(
	actorID string,
) ([]model.AuditLog, error) {

	return s.auditRepo.ListByActorID(actorID)
}

func (s *auditService) ListByEntity(
	entityType string,
	entityID string,
) ([]model.AuditLog, error) {

	return s.auditRepo.ListByEntity(
		entityType,
		entityID,
	)
}
func normalizeAction(action string) string {
	switch action {
	case "login":
		return "USER_LOGIN"
	case "create_user":
		return "USER_CREATE"
	case "update_user":
		return "USER_UPDATE"
	case "delete_user":
		return "USER_DELETE"
	default:
		return action
	}
}
func (s *auditService) ListByTenant(
	tenantID string,
) ([]model.AuditLog, error) {

	return s.auditRepo.ListByTenant(
		tenantID,
	)
}

func (s *auditService) ListWithFilter(
	tenantID string,
	actorID string,
	entityType string,
	entityID string,
	from time.Time,
	to time.Time,
	limit int,
	offset int,
) ([]model.AuditLog, error) {

	return s.auditRepo.ListWithFilter(
		tenantID,
		actorID,
		entityType,
		entityID,
		from,
		to,
		limit,
		offset,
	)
}