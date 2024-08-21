package model

type TransactionRequest struct {
	TransactionType  string  `json:"transaction_type"`
	CategoryTypeCode uint    `json:"category_type_code"`
	Description      string  `json:"description"`
	Title            string  `json:"title"`
	Amount           float64 `json:"amount"`
}

type TransactionResponse struct {
	ID               uint    `json:"id"`
	UUID             string  `json:"uuid"`
	TransactionCode  uint    `json:"transaction_code"`
	CategoryTypeCode uint    `json:"category_type_code"`
	Description      string  `json:"description"`
	Title            string  `json:"title"`
	Amount           float64 `json:"amount"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	DeletedAt        string  `json:"deleted_at"`
	CreatedBy        int     `json:"created_by"`
	UpdatedBy        int     `json:"updated_by"`
	DeletedBy        int     `json:"deleted_by"`
}

type TransactionUpdateRequest struct {
	TransactionType  string  `json:"transaction_type"`
	CategoryTypeCode uint    `json:"category_type_code"`
	Description      string  `json:"description"`
	Title            string  `json:"title"`
	Amount           float64 `json:"amount"`
}

// type TransactionPaginationRequest struct {
// 	model
// }
