package dto

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type AssignRoleRequest struct {
	RoleName string `json:"role" validate:"required"`
}
