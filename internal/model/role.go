package model

import "time"

type RoleCreateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type RoleUpdateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type RoleResponse struct {
	ID        int        `json:"id"`
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Alias     string     `json:"alias"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	CreatedBy int        `json:"created_by"`
	UpdatedBy int        `json:"updated_by"`
	DeletedBy int        `json:"deleted_by"`
}
