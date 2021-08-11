package categories

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, categoryId int) (Domain, error)
	GetByName(ctx context.Context, categoryName string) (Domain, error)
	Store(ctx context.Context, categoryDomain *Domain) (Domain, error)
	Update(ctx context.Context, newsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, categoryId int) (Domain, error)
	GetByName(ctx context.Context, categoryName string) (Domain, error)
	Store(ctx context.Context, categoriesDomain *Domain) (Domain, error)
	Update(ctx context.Context, categoriesDomain *Domain) (Domain, error)
}
