package model

import "time"

type UserPermissionOverride struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	PermissionID string    `json:"permissionId"`
	Granted      bool      `json:"granted"`
	Reason       string    `json:"reason"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	DataScope    string    `json:"dataScope"`
}