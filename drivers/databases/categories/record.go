package categories

import (
	"injar/businesses/categories"
	categoriesUsecase "injar/businesses/categories"
	"time"
)

type Categories struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Categories) toDomain() categories.Domain {
	return categories.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(categoriesDomain *categoriesUsecase.Domain) *Categories {
	return &Categories{
		ID:        categoriesDomain.ID,
		Name:      categoriesDomain.Name,
		CreatedAt: categoriesDomain.CreatedAt,
		UpdatedAt: categoriesDomain.UpdatedAt,
	}
}
