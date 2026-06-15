package dto


type AuditLogResponse struct {
	ID              string `json:"id"`
	TenantID        string `json:"tenantId"`
	Actor           string `json:"actor"`
	Action          string `json:"action"`
	EntityType      string `json:"entityType"`
	EntityID        string `json:"entityId"`
	EventType       string `json:"eventType"`
	EventTimestamp  string `json:"eventTimestamp"`
}