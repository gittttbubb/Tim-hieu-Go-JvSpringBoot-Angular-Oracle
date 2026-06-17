package model

type Role struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	DisplayName string  `json:"displayName" db:"display_name"`
	Description *string `json:"description,omitempty" db:"description"`
}