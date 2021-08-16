package favourites

import (
	"context"
	"injar/usecase/favourites"

	"gorm.io/gorm"
)

type mysqlFavouritesRepository struct {
	DB *gorm.DB
}

func NewMySQLFavouritesRepository(conn *gorm.DB) favourites.Repository {
	return &mysqlFavouritesRepository{
		DB: conn,
	}
}

func (cr *mysqlFavouritesRepository) GetByUserID(ctx context.Context, userID int) ([]favourites.Domain, error) {
	rec := []Favourites{}

	err := cr.DB.Joins("Users").Joins("Webinars").Where("favourites.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []favourites.Domain{}, err
	}

	favouriteDomain := []favourites.Domain{}
	for _, value := range rec {
		favouriteDomain = append(favouriteDomain, value.toDomain())
	}

	return favouriteDomain, nil
}

func (cr *mysqlFavouritesRepository) GetByID(ctx context.Context, ID int) (favourites.Domain, error) {
	rec := Favourites{}
	err := cr.DB.Joins("Users").Joins("Webinars").Where("favourites.id = ?", ID).First(&rec).Error
	if err != nil {
		return favourites.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlFavouritesRepository) Store(ctx context.Context, favouritesDomain *favourites.Domain) (favourites.Domain, error) {
	rec := fromDomain(*favouritesDomain)

	result := nr.DB.Create(&rec)
	if result.Error != nil {
		return favourites.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlFavouritesRepository) Delete(ctx context.Context, favouritesDomain *favourites.Domain) (favourites.Domain, error) {
	rec := fromDomain(*favouritesDomain)

	result := nr.DB.Delete(rec)

	if result.Error != nil {
		return favourites.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
