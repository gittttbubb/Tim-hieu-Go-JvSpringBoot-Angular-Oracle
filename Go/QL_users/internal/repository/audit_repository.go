package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type AuditRepository interface {
	Create(log *model.AuditLog,) error
	GetByID(id string,) (*model.AuditLog,error)
	List() ([]model.AuditLog,error)
}

type auditRepository struct {
	db *sql.DB
}

func NewAuditRepository(db *sql.DB) AuditRepository {
	return &auditRepository{
		db: db,
	}
}

func (r *auditRepository) Create(logData *model.AuditLog,) error {
	query := `INSERT INTO audit_logs(
		id,
		tenant_id,
		actor,
		actor_id,
		action,
		entity_type,
		entity_id,
		before_data,
		after_data,
		ip_address,
		reason,
		event_type,
		actor_identifier,
		target_user_id,
		metadata
	)
	VALUES(
		:1,:2,:3,:4,:5,
		:6,:7,:8,:9,:10,
		:11,:12,:13,:14,:15
	)`
	_, err := r.db.Exec(
		query,
		logData.ID,
		logData.TenantID,
		logData.Actor,
		logData.ActorID,
		logData.Action,
		logData.EntityType,
		logData.EntityID,
		logData.BeforeData,
		logData.AfterData,
		logData.IPAddress,
		logData.Reason,
		logData.EventType,
		logData.ActorIdentifier,
		logData.TargetUserID,
		logData.Metadata,
	)
	return err
}

func (r *auditRepository) GetByID(id string,) (*model.AuditLog,error) {
	query := `
	SELECT
		id,
		tenant_id,
		actor,
		actor_id,
		action,
		entity_type,
		entity_id,
		before_data,
		after_data,
		ip_address,
		event_timestamp,
		reason,
		event_type,
		actor_identifier,
		target_user_id,
		metadata
	FROM audit_logs WHERE id = :1`

	var audit model.AuditLog
	err := r.db.QueryRow(query, id,
	).Scan(
		&audit.ID,
		&audit.TenantID,
		&audit.Actor,
		&audit.ActorID,
		&audit.Action,
		&audit.EntityType,
		&audit.EntityID,
		&audit.BeforeData,
		&audit.AfterData,
		&audit.IPAddress,
		&audit.EventTimestamp,
		&audit.Reason,
		&audit.EventType,
		&audit.ActorIdentifier,
		&audit.TargetUserID,
		&audit.Metadata,
	)

	if err != nil {
		return nil, err
	}
	return &audit,nil
}

func (r *auditRepository) List() ([]model.AuditLog, error,) {
	rows, err := r.db.Query(`
		SELECT
			id,
			tenant_id,
			actor,
			action,
			entity_type,
			entity_id,
			event_timestamp
		FROM audit_logs
		ORDER BY event_timestamp DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.AuditLog
	for rows.Next() {
		var audit model.AuditLog
		err := rows.Scan(
			&audit.ID,
			&audit.TenantID,
			&audit.Actor,
			&audit.Action,
			&audit.EntityType,
			&audit.EntityID,
			&audit.EventTimestamp,)
		if err != nil {
			return nil, err
		}
		result = append(result,audit)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result,nil
}