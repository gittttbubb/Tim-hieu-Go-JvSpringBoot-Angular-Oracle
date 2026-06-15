package dto

import (
	"go-user-management/internal/model"
)

type LoginResponse struct {
	Token               string                      `json:"token"`
	MustChangePassword  bool                        `json:"mustChangePassword"`
	Permissions         []model.EffectivePermission `json:"permissions"`
}