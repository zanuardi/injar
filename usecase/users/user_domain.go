package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Username  string
	Password  string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Usecase interface {
	CreateToken(ctx context.Context, username, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, ID int) (Domain, error)
	Update(ctx context.Context, usersDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, ID int) (Domain, error)
	Update(ctx context.Context, usersDomain *Domain) (Domain, error)
}
