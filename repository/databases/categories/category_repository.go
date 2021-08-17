package categories

import (
	"context"
	"injar/usecase/categories"

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

func (repo *mysqlCategoriesRepository) Fetch(ctx context.Context, page, perpage int) ([]categories.Domain, int, error) {
	rec := []Categories{}

	offset := (page - 1) * perpage
	err := repo.DB.Find(&rec).Offset(offset).Limit(perpage).Error
	if err != nil {
		return []categories.Domain{}, 0, err
	}

	var totalData int64
	err = repo.DB.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []categories.Domain{}, 0, err
	}

	var domainCategory []categories.Domain
	for _, value := range rec {
		domainCategory = append(domainCategory, value.toDomain())
	}
	return domainCategory, int(totalData), nil
}

func (repo *mysqlCategoriesRepository) Find(ctx context.Context) ([]categories.Domain, error) {
	rec := []Categories{}

	repo.DB.Find(&rec)
	categoryDomain := []categories.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.toDomain())
	}

	return categoryDomain, nil
}

func (repo *mysqlCategoriesRepository) GetByID(ctx context.Context, userId int) (categories.Domain, error) {
	rec := Categories{}
	err := repo.DB.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlCategoriesRepository) GetByName(ctx context.Context, categoryName string) (categories.Domain, error) {
	rec := Categories{}
	err := repo.DB.Where("name = ?", categoryName).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlCategoriesRepository) Store(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := repo.DB.Select("Name", "repoeatedAt").Create(&rec)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := repo.DB.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (repo *mysqlCategoriesRepository) Update(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := repo.DB.Updates(rec)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := repo.DB.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (repo *mysqlCategoriesRepository) Delete(ctx context.Context, categoriesDomain *categories.Domain) (categories.Domain, error) {
	rec := fromDomain(categoriesDomain)

	result := repo.DB.Delete(rec)

	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	err := repo.DB.Preload("Category").First(&rec, rec.ID).Error
	if err != nil {
		return categories.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
