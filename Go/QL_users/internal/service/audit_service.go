package service

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"go-user-management/internal/dto/audit"
	"go-user-management/internal/model"
	"go-user-management/internal/repository"
)

type AuditService interface {
	Create(req dto.CreateAuditLogRequest,) error
	GetByID(id string,) (*dto.AuditLogResponse, error)
	List() ([]dto.AuditLogResponse, error)
}

type auditService struct {
	auditRepo repository.AuditRepository
}

func NewAuditService(auditRepo repository.AuditRepository,) AuditService {
	return &auditService{
		auditRepo: auditRepo,
	}
}

func (s *auditService) Create(req dto.CreateAuditLogRequest,) error {
	if strings.TrimSpace(req.TenantID) == "" {
		return errors.New("tenant id is required")
	}
	if strings.TrimSpace(req.Action) == "" {
		return errors.New("action is required")
	}
	if strings.TrimSpace(req.EntityType) == "" {
		return errors.New("entity type is required")
	}

	logData := &model.AuditLog{
		ID:              uuid.NewString(),
		TenantID:        req.TenantID,
		Actor:           req.Actor,
		ActorID:         req.ActorID,
		Action:          req.Action,
		EntityType:      req.EntityType,
		EntityID:        req.EntityID,
		BeforeData:      req.BeforeData,
		AfterData:       req.AfterData,
		IPAddress:       req.IPAddress,
		Reason:          req.Reason,
		EventType:       req.EventType,
		ActorIdentifier: req.ActorIdentifier,
		TargetUserID:    req.TargetUserID,
		Metadata:        req.Metadata,
	}
	return s.auditRepo.Create(logData)
}

func (s *auditService) GetByID(id string,) (*dto.AuditLogResponse, error) {
	audit, err := s.auditRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.AuditLogResponse{
		ID:             audit.ID,
		TenantID:       audit.TenantID,
		Actor:          audit.Actor,
		Action:         audit.Action,
		EntityType:     audit.EntityType,
		EntityID:       audit.EntityID,
		EventType:      audit.EventType,
		EventTimestamp: audit.EventTimestamp.Format(time.RFC3339),
	}, nil
}

func (s *auditService) List() ([]dto.AuditLogResponse,error,) {
	items, err := s.auditRepo.List()
	if err != nil {
		return nil, err
	}
	result := make([]dto.AuditLogResponse, 0, len(items),)
	for _, item := range items {
		result = append(
			result,
			dto.AuditLogResponse{
				ID:             item.ID,
				TenantID:       item.TenantID,
				Actor:          item.Actor,
				Action:         item.Action,
				EntityType:     item.EntityType,
				EntityID:       item.EntityID,
				EventType:      item.EventType,
				EventTimestamp: item.EventTimestamp.Format(time.RFC3339),
			},
		)
	}
	return result, nil
}