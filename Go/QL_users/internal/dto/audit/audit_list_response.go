package dto

type AuditLogListResponse struct {
	Items []AuditLogResponse `json:"items"`
	Total int                `json:"total"`
}