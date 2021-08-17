package webinars

import (
	"context"
	"injar/usecase/webinars"

	"gorm.io/gorm"
)

type mysqlWebinarRepository struct {
	DB *gorm.DB
}

func NewMySQLWebinarRepository(conn *gorm.DB) webinars.Repository {
	return &mysqlWebinarRepository{
		DB: conn,
	}
}

func (repo *mysqlWebinarRepository) GetAll(ctx context.Context, name string) ([]webinars.Domain, error) {
	rec := []Webinars{}

	err := repo.DB.Preload("Categories").Where("webinars.name LIKE ?", "%"+name+"%").Find(&rec).Error
	if err != nil {
		return []webinars.Domain{}, err
	}

	webinarDomain := []webinars.Domain{}
	for _, value := range rec {
		webinarDomain = append(webinarDomain, value.toDomain())
	}

	return webinarDomain, nil
}

func (repo *mysqlWebinarRepository) GetByID(ctx context.Context, ID int) (webinars.Domain, error) {
	rec := Webinars{}
	err := repo.DB.Preload("Categories").Where("webinars.id = ?", ID).First(&rec).Error
	if err != nil {
		return webinars.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlWebinarRepository) GetByName(ctx context.Context, webinarName string) (webinars.Domain, error) {
	rec := Webinars{}
	err := repo.DB.Joins("Categories").Where("name = ?", webinarName).First(&rec).Error
	if err != nil {
		return webinars.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlWebinarRepository) Store(ctx context.Context, webinarsDomain *webinars.Domain) (webinars.Domain, error) {
	rec := fromDomain(*webinarsDomain)

	result := repo.DB.Create(&rec)
	if result.Error != nil {
		return webinars.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (repo *mysqlWebinarRepository) Update(ctx context.Context, webinarsDomain *webinars.Domain) (webinars.Domain, error) {
	rec := fromDomain(*webinarsDomain)

	result := repo.DB.Updates(rec)
	if result.Error != nil {
		return webinars.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (repo *mysqlWebinarRepository) Delete(ctx context.Context, webinarsDomain *webinars.Domain) (webinars.Domain, error) {
	rec := fromDomain(*webinarsDomain)

	result := repo.DB.Delete(rec)

	if result.Error != nil {
		return webinars.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
