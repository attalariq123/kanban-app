package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {

	var categoryList []entity.Category
	err := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&categoryList).Error

	return categoryList, err
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {

	err = r.db.WithContext(ctx).Create(&category).Error
	if err != nil {
		return 0, err
	}

	return category.ID, err
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {

	return r.db.WithContext(ctx).Create(&categories).Error
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {

	var category entity.Category
	err := r.db.WithContext(ctx).Find(&category, id).Error
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {

	return r.db.WithContext(ctx).Model(&category).Updates(&category).Error
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {

	return r.db.WithContext(ctx).Delete(&entity.Category{}, id).Error
}
