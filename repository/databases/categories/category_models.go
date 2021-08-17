package categories

import (
	"injar/usecase/categories"
	categoriesUsecase "injar/usecase/categories"
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
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
