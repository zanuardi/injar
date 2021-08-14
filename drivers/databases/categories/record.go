package categories

import (
	"injar/businesses/categories"
	categoriesUsecase "injar/businesses/categories"
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Categories) toDomain() categories.Domain {
	return categories.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(categoriesDomain *categoriesUsecase.Domain) *Categories {
	return &Categories{
		ID:        categoriesDomain.ID,
		Name:      categoriesDomain.Name,
		CreatedAt: categoriesDomain.CreatedAt,
		UpdatedAt: categoriesDomain.UpdatedAt,
		DeletedAt: categoriesDomain.DeletedAt,
	}
}
