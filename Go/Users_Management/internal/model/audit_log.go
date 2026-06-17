package model

import "time"

type AuditLog struct {
	ID              string     `json:"id" db:"id"`
	TenantID        string     `json:"tenantId" db:"tenant_id"`
	Actor           string     `json:"actor" db:"actor"`
	ActorID         *string    `json:"actorId,omitempty" db:"actor_id"`
	Action          string     `json:"action" db:"action"`
	EntityType      string     `json:"entityType" db:"entity_type"`
	EntityID        string     `json:"entityId" db:"entity_id"`
	BeforeData      *string    `json:"beforeData,omitempty" db:"before_data"`
	AfterData       *string    `json:"afterData,omitempty" db:"after_data"`
	IPAddress       *string    `json:"ipAddress,omitempty" db:"ip_address"`
	EventTimestamp  time.Time  `json:"eventTimestamp" db:"event_timestamp"`
	Reason          *string    `json:"reason,omitempty" db:"reason"`
	EventType       *string    `json:"eventType,omitempty" db:"event_type"`
	ActorIdentifier *string    `json:"actorIdentifier,omitempty" db:"actor_identifier"`
	TargetUserID    *string    `json:"targetUserId,omitempty" db:"target_user_id"`
	Metadata        *string    `json:"metadata,omitempty" db:"metadata"`
}