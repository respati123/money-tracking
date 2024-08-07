package repository

import (
	"errors"
	"reflect"

	"github.com/respati123/money-tracking/internal/model"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, data *T) error {
	return db.Create(data).Error
}

func (r *Repository[T]) Update(db *gorm.DB, data *T) error {
	return db.Save(data).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, data *T) error {
	return db.Delete(data).Error
}

func (r *Repository[T]) CountById(db *gorm.DB, uuid string) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("uuid = ?", uuid).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindAll(db *gorm.DB, result *[]T, request model.PaginationRequest) (*[]T, model.PaginationModel, error) {

	var (
		totalData int64
		totalPage int
	)
	db = db.Model(new(T))

	if request.Filter != nil {
		for key, value := range request.Filter {
			_, isExist := reflect.TypeOf(new(T)).Elem().FieldByName(key)
			if !isExist {
				return nil, model.PaginationModel{}, errors.New(" field " + key + " is not exist")
			}
			db = db.Where(key+"=?", value)
		}
	}

	offset := (request.Page - 1) * request.PerPage
	db.Count(&totalData)
	err := db.Limit(request.PerPage).Offset(offset).Find(&result).Error

	if int(totalData)/request.PerPage == 0 {
		totalPage = 1
	} else {
		totalPage = int(totalData) / request.PerPage
	}

	metadata := model.PaginationModel{
		TotalData:   int(totalData),
		TotalPage:   totalPage,
		CurrentPage: request.Page,
		PerPage:     request.PerPage,
	}
	if err != nil {
		return nil, model.PaginationModel{}, err
	}

	return result, metadata, nil
}
