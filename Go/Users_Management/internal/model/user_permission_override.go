package model

import "time"

type UserPermissionOverride struct {
	ID           string    `json:"id" db:"id"`
	UserID       string    `json:"userId" db:"user_id"`
	PermissionID string    `json:"permissionId" db:"permission_id"`
	Granted      bool      `json:"granted" db:"granted"`
	Reason       *string   `json:"reason,omitempty" db:"reason"`
	CreatedBy    string    `json:"createdBy" db:"created_by"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	DataScope    string    `json:"dataScope" db:"data_scope"`
}