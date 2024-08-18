package model

import "time"

type CategoryCreateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type CategoryUpdateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type CategoryResponse struct {
	ID           uint       `json:"id"`
	UUID         string     `json:"uuid"`
	CategoryCode int        `json:"category_code"`
	Name         string     `json:"name"`
	Alias        string     `json:"alias"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	CreatedBy    int        `json:"created_by"`
	UpdatedBy    int        `json:"updated_by"`
	DeletedBy    int        `json:"deleted_by"`
}
