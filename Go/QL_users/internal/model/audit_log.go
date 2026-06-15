package model

import "time"

type AuditLog struct {
	ID              string    `json:"id"`
	TenantID        string    `json:"tenantId"`
	Actor           string    `json:"actor"`
	ActorID         string    `json:"actorId"`
	Action          string    `json:"action"`
	EntityType      string    `json:"entityType"`
	EntityID        string    `json:"entityId"`
	BeforeData      string    `json:"beforeData"`
	AfterData       string    `json:"afterData"`
	IPAddress       string    `json:"ipAddress"`
	EventTimestamp  time.Time `json:"eventTimestamp"`
	Reason          string    `json:"reason"`
	EventType       string    `json:"eventType"`
	ActorIdentifier string    `json:"actorIdentifier"`
	TargetUserID    string    `json:"targetUserId"`
	Metadata        string    `json:"metadata"`
}