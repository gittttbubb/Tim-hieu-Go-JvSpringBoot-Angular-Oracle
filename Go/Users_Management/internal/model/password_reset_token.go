package model

import "time"

type PasswordResetToken struct {
	ID         string     `json:"id" db:"id"`
	UserID     string     `json:"userId" db:"user_id"`
	TokenHash  string     `json:"tokenHash" db:"token_hash"`
	ExpiresAt  time.Time  `json:"expiresAt" db:"expires_at"`
	UsedAt     *time.Time `json:"usedAt,omitempty" db:"used_at"`
	RevokedAt  *time.Time `json:"revokedAt,omitempty" db:"revoked_at"`
	CreatedIP  *string    `json:"createdIp,omitempty" db:"created_ip"`
	UserAgent  *string    `json:"userAgent,omitempty" db:"user_agent"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
}