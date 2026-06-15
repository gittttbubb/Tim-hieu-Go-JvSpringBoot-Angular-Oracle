package model

type RolePermissionDetail struct {
	PermissionID string
	FeatureCode string
	Granted     bool
	DataScope   string
}