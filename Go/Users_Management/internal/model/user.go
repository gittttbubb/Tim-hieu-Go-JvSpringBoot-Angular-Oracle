package model

import "time"

type User struct {
	ID                   string    `json:"id" db:"id"`
	TenantID             string    `json:"tenantId" db:"tenant_id"`
	FullName             string    `json:"fullName" db:"full_name"`
	Username             string    `json:"username" db:"username"`
	Email                string    `json:"email" db:"email"`
	Phone                string    `json:"phone" db:"phone"`
	PasswordHash         string    `json:"-" db:"password_hash"`
	RoleID               string    `json:"roleId" db:"role_id"`
	Status               string    `json:"status" db:"status"`
	MustChangePassword   bool      `json:"mustChangePassword" db:"must_change_password"`
	CreatedAt            time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt            time.Time `json:"updatedAt" db:"updated_at"`
	CreatedBy            *string   `json:"createdBy,omitempty" db:"created_by"`
	PasswordChangedAt    time.Time `json:"passwordChangedAt" db:"password_changed_at"`
}