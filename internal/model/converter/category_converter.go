package converter

import (
	"time"

	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
)

type CategoryConverter struct{}

func NewCategoryConverter() *CategoryConverter {
	return &CategoryConverter{}
}

func (cc *CategoryConverter) ToCategoryResponse(category entity.Category) model.CategoryResponse {

	deletedAt := ""
	if category.DeletedAt.Valid {
		deletedAt = category.DeletedAt.Time.Format(time.RFC3339)
	}

	return model.CategoryResponse{
		ID:           category.ID,
		UUID:         category.UUID,
		Alias:        category.Alias,
		Name:         category.Name,
		CategoryCode: category.CategoryCode,
		CreatedAt:    category.CreatedAt.Format(time.RFC3339),
		CreatedBy:    category.CreatedBy,
		UpdatedAt:    category.UpdatedAt.Format(time.RFC3339),
		UpdatedBy:    category.UpdatedBy,
		DeletedAt:    deletedAt,
		DeletedBy:    category.DeletedBy,
	}
}

func (cc *CategoryConverter) ToCategoryResponses(categories *[]entity.Category) []model.CategoryResponse {
	var category []model.CategoryResponse
	for _, value := range *categories {
		category = append(category, cc.ToCategoryResponse(value))
	}
	return category
}
