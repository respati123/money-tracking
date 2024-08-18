package converter

import (
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
)

type CategoryConverter struct{}

func NewCategoryConverter() *CategoryConverter {
	return &CategoryConverter{}
}

func (cc *CategoryConverter) ToCategoryResponse(category entity.Category) model.CategoryResponse {

	return model.CategoryResponse{
		ID:           category.ID,
		UUID:         category.UUID,
		Alias:        category.Alias,
		Name:         category.Name,
		CategoryCode: category.CategoryCode,
		CreatedAt:    category.CreatedAt,
		CreatedBy:    category.CreatedBy,
		UpdatedAt:    category.UpdatedAt,
		UpdatedBy:    category.UpdatedBy,
		DeletedAt:    &category.DeletedAt.Time,
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
