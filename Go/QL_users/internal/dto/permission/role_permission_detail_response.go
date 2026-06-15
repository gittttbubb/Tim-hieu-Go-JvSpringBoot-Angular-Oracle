package dto

type RolePermissionDetailResponse struct {
	PermissionID string `json:"permissionId"`
	FeatureCode  string `json:"featureCode"`
	Granted      bool   `json:"granted"`
	DataScope    string `json:"dataScope"`
}