package model

type Permission struct {
	ID           string  `json:"id" db:"id"`
	FeatureGroup string  `json:"featureGroup" db:"feature_group"`
	FeatureCode  string  `json:"featureCode" db:"feature_code"`
	Action       string  `json:"action" db:"action"`
	Description  *string `json:"description,omitempty" db:"description"`
}