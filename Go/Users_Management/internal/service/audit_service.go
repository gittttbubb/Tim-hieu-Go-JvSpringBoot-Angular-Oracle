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
	Log(audit *model.AuditLog) error
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

func (s *auditService) Log(audit *model.AuditLog,) error {
    if audit.ID == "" {
        audit.ID = utils.NewUUID()
    }
    if audit.EventTimestamp.IsZero() {
        audit.EventTimestamp = time.Now()
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