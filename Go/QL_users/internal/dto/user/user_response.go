package dto

import "time"

type UserResponse struct {
	ID                 string    `json:"id"`
	TenantID           string    `json:"tenantId"`
	FullName           string    `json:"fullName"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	Phone              string    `json:"phone"`
	RoleID             string    `json:"roleId"`
	Status             string    `json:"status"`
	MustChangePassword bool      `json:"mustChangePassword"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	PasswordChangedAt  time.Time `json:"passwordChangedAt"`
}