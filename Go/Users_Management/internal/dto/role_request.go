package dto

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required,max=50"`
	DisplayName string `json:"displayName" validate:"required,max=100"`
	Description string `json:"description" validate:"max=255"`
}

type UpdateRoleRequest struct {
	DisplayName string `json:"displayName" validate:"required,max=100"`
	Description string `json:"description" validate:"max=255"`
}

type RoleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}