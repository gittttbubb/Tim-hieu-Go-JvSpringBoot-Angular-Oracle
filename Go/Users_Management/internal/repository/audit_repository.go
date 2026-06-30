package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"go-rbac-system/internal/model"
)

type AuditRepository interface {
	Create(log *model.AuditLog) error
	GetAll() ([]model.AuditLog, error)
	GetByID(id string) (*model.AuditLog, error)
	ListByActorID(actorID string) ([]model.AuditLog, error)
	ListByEntity(entityType string, entityID string) ([]model.AuditLog, error)
	ListByTenant(tenantID string) ([]model.AuditLog, error)
	ListWithFilter(tenantID string, actorID string, entityType string, entityID string, from time.Time, to time.Time, limit int,
		offset int) ([]model.AuditLog, error)
}

type auditRepository struct {
	db *sql.DB
}

func NewAuditRepository(db *sql.DB) AuditRepository {
	return &auditRepository{
		db: db,
	}
}

func scanAuditLog(scanner interface{ Scan(dest ...any) error }) (*model.AuditLog, error) {
	var item model.AuditLog
	err := scanner.Scan(
		&item.ID,
		&item.TenantID,
		&item.Actor,
		&item.ActorID,
		&item.Action,
		&item.EntityType,
		&item.EntityID,
		&item.BeforeData,
		&item.AfterData,
		&item.IPAddress,
		&item.EventTimestamp,
		&item.Reason,
		&item.EventType,
		&item.ActorIdentifier,
		&item.TargetUserID,
		&item.Metadata,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *auditRepository) GetByID(
	id string,
) (*model.AuditLog, error) {
	query := `SELECT id, tenant_id, actor, actor_id, action, entity_type, entity_id, before_data, after_data,
			ip_address, event_timestamp, reason, event_type, actor_identifier, target_user_id, metadata
		FROM audit_logs WHERE id = :1`
	row := r.db.QueryRow(query, id)
	return scanAuditLog(row)
}

func (r *auditRepository) Create(log *model.AuditLog) error {
	query := `INSERT INTO audit_logs (id, tenant_id, actor, actor_id, action, entity_type, entity_id, before_data,
			after_data, ip_address, event_timestamp, reason, event_type, actor_identifier, target_user_id, metadata)
		VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)`
	_, err := r.db.Exec(
		query,
		log.ID,
		log.TenantID,
		log.Actor,
		log.ActorID,
		log.Action,
		log.EntityType,
		log.EntityID,
		log.BeforeData,
		log.AfterData,
		log.IPAddress,
		log.EventTimestamp,
		log.Reason,
		log.EventType,
		log.ActorIdentifier,
		log.TargetUserID,
		log.Metadata,
	)
	return err
}
func (r *auditRepository) GetAll() ([]model.AuditLog, error) {
	query := `SELECT id, tenant_id, actor, actor_id, action, entity_type, entity_id, before_data, after_data,
			ip_address, event_timestamp, reason, event_type, actor_identifier, target_user_id,metadata
		FROM audit_logs ORDER BY event_timestamp DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.AuditLog, 0)
	for rows.Next() {
		item, err := scanAuditLog(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, rows.Err()
}

func (r *auditRepository) ListByActorID(actorID string) ([]model.AuditLog, error) {
	query := `SELECT id, tenant_id, actor, actor_id, action, entity_type, entity_id, before_data, after_data, ip_address,
			event_timestamp, reason, event_type, actor_identifier, target_user_id, metadata
			FROM audit_logs WHERE actor_id = :1 ORDER BY event_timestamp DESC, id DESC`
	rows, err := r.db.Query(query, actorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.AuditLog, 0)
	for rows.Next() {
		item, err := scanAuditLog(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, rows.Err()
}

func (r *auditRepository) ListByEntity(entityType string, entityID string) ([]model.AuditLog, error) {
	query := `SELECT id, tenant_id, actor, actor_id, action, entity_type, entity_id, before_data, after_data, ip_address,
			event_timestamp, reason, event_type, actor_identifier, 	target_user_id, metadata
			FROM audit_logs WHERE entity_type = :1 AND entity_id = :2 ORDER BY event_timestamp DESC`
	rows, err := r.db.Query(query, entityType, entityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.AuditLog, 0)
	for rows.Next() {
		item, err := scanAuditLog(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, rows.Err()
}

func (r *auditRepository) ListByTenant(tenantID string) ([]model.AuditLog, error) {
	query := `SELECT id, tenant_id, actor, actor_id, action, entity_type, entity_id, before_data, after_data, ip_address,
			event_timestamp, reason, event_type, actor_identifier, target_user_id, metadata
			FROM audit_logs WHERE tenant_id = :1 ORDER BY event_timestamp DESC`
	rows, err := r.db.Query(query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.AuditLog, 0)
	for rows.Next() {
		item, err := scanAuditLog(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, rows.Err()
}

func (r *auditRepository) ListWithFilter(tenantID string, actorID string, entityType string, entityID string,
	from time.Time, to time.Time, limit int, offset int) ([]model.AuditLog, error) {
	query := `
		SELECT id, tenant_id, actor, actor_id, action, entity_type, entity_id,
		       before_data, after_data, ip_address, event_timestamp, reason,
		       event_type, actor_identifier, target_user_id, metadata
		FROM audit_logs
		WHERE tenant_id = :1
	`

	args := []interface{}{tenantID}
	idx := 2
	if actorID != "" {
		query += " AND actor_id = :" + strconv.Itoa(idx)
		args = append(args, actorID)
		idx++
	}

	if entityType != "" {
		query += " AND entity_type = :" + strconv.Itoa(idx)
		args = append(args, entityType)
		idx++
	}
	if entityID != "" {
		query += " AND entity_id = :" + strconv.Itoa(idx)
		args = append(args, entityID)
		idx++
	}
	if !from.IsZero() {
		query += " AND event_timestamp >= :" + strconv.Itoa(idx)
		args = append(args, from)
		idx++
	}
	if !to.IsZero() {
		query += " AND event_timestamp <= :" + strconv.Itoa(idx)
		args = append(args, to)
		idx++
	}
	query += " ORDER BY event_timestamp DESC"
	// Oracle pagination chuẩn
	if limit > 0 {
		query += fmt.Sprintf(
			" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY",
			offset,
			limit,
		)
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.AuditLog, 0)
	for rows.Next() {
		item, err := scanAuditLog(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, rows.Err()
}
