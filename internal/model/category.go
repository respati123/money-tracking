package model

type CategoryCreateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type CategoryUpdateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	UUID         string `json:"uuid"`
	CategoryCode int    `json:"category_code"`
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    int    `json:"created_by"`
	UpdatedAt    string `json:"updated_at"`
	UpdatedBy    int    `json:"updated_by"`
	DeletedAt    string `json:"deleted_at"`
	DeletedBy    int    `json:"deleted_by"`
}
