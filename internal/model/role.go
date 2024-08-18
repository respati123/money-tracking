package model

type RoleCreateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type RoleUpdateRequest struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type RoleResponse struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	CreatedBy int    `json:"created_by"`
	UpdatedBy int    `json:"updated_by"`
	DeletedBy int    `json:"deleted_by"`
}
