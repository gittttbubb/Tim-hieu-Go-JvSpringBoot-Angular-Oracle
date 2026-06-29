package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int64  `json:"expiresIn"`

	UserID   string `json:"userId"`
	Username string `json:"username"`
	RoleID   string `json:"roleId"`
	Permissions []UserPermission `json:"permissions"`

	MustChangePassword bool `json:"mustChangePassword"`
}

type UserPermission struct {
	PermissionID string `json:"permissionId"`
	FeatureCode string `json:"featureCode"`
	DataScope    string `json:"dataScope"`
}