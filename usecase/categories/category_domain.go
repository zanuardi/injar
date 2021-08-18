package categories

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	FindAll(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, categoryId int) (Domain, error)
	GetByName(ctx context.Context, categoryName string) (Domain, error)
	Store(ctx context.Context, categoryDomain *Domain) (Domain, error)
	Update(ctx context.Context, newsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, categoriesDomain *Domain) (*Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, categoryId int) (Domain, error)
	GetByName(ctx context.Context, categoryName string) (Domain, error)
	Store(ctx context.Context, categoriesDomain *Domain) (Domain, error)
	Update(ctx context.Context, categoriesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, categoriesDomain *Domain) (Domain, error)
}
