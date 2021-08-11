package categories

import (
	"context"
	"injar/businesses/categories"

	"gorm.io/gorm"
)

type mysqlCategoriesRepository struct {
	Conn *gorm.DB
}

func NewMySQLUserRepository(conn *gorm.DB) categories.Repository {
	return &mysqlCategoriesRepository{
		Conn: conn,
	}
}

func (cr *mysqlCategoriesRepository) Fetch(ctx context.Context, page, perpage int) ([]categories.Domain, int, error) {
	rec := []Categories{}

	offset := (page - 1) * perpage
	err := cr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []categories.Domain{}, 0, err
	}

	var totalData int64
	err = cr.Conn.Count(&totalData).Error
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

	query := cr.Conn.Where("deleted_at = ?", nil)

	err := query.Find(&rec).Error
	if err != nil {
		return []categories.Domain{}, err
	}

	categoryDomain := []categories.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.toDomain())
	}

	return categoryDomain, nil
}

func (cr *mysqlCategoriesRepository) GetByID(ctx context.Context, userId int) (categories.Domain, error) {
	rec := Categories{}
	err := cr.Conn.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) GetByName(ctx context.Context, categoryName string) (categories.Domain, error) {
	rec := Categories{}
	err := nr.Conn.Where("name = ?", categoryName).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) Store(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlCategoriesRepository) Update(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
