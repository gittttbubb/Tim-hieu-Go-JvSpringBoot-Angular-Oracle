package dto

type PermissionResponse struct {
	ID           string `json:"id"`
	FeatureGroup string `json:"featureGroup"`
	FeatureCode  string `json:"featureCode"`
	Action       string `json:"action"`
	Description  string `json:"description"`
}