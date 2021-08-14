package categories

import (
	"context"
	"fmt"
	"injar/businesses/categories"

	"gorm.io/gorm"
)

type mysqlCategoriesRepository struct {
	DB *gorm.DB
}

func NewMySQLCategoryRepository(conn *gorm.DB) categories.Repository {
	return &mysqlCategoriesRepository{
		DB: conn,
	}
}

func (cr *mysqlCategoriesRepository) Fetch(ctx context.Context, page, perpage int) ([]categories.Domain, int, error) {
	rec := []Categories{}

	offset := (page - 1) * perpage
	err := cr.DB.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []categories.Domain{}, 0, err
	}

	var totalData int64
	err = cr.DB.Count(&totalData).Error
	if err != nil {
		return []categories.Domain{}, 0, err
	}

	var categoriesDomain []categories.Domain
	for _, value := range rec {
		categoriesDomain = append(categoriesDomain, value.toDomain())
	}
	return categoriesDomain, int(totalData), nil
}

func (cr *mysqlCategoriesRepository) Find(ctx context.Context) ([]categories.Domain, error) {
	rec := []Categories{}

	// query := cr.Conn.Where("deleted_at = ?", nil)

	// err := query.Find(&rec).Error
	// if err != nil {
	// 	return []categories.Domain{}, err
	// }

	categoryDomain := []categories.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.toDomain())
	}
	fmt.Println(categoryDomain)

	return categoryDomain, nil
}

func (cr *mysqlCategoriesRepository) GetByID(ctx context.Context, userId int) (categories.Domain, error) {
	rec := Categories{}
	err := cr.DB.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) GetByName(ctx context.Context, categoryName string) (categories.Domain, error) {
	rec := Categories{}
	err := nr.DB.Where("name = ?", categoryName).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) Store(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := nr.DB.Create(&rec)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := nr.DB.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) Update(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := nr.DB.Updates(rec)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := nr.DB.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) Delete(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := nr.DB.Delete(rec)

	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := nr.DB.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
