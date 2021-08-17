package favourites

import (
	"context"
	"injar/repository/databases/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          int
	UserID      int
	Users       users.Users
	WebinarID   int
	WebinarName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Usecase interface {
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	Store(ctx context.Context, favouriteDomain *Domain) (Domain, error)
	Delete(ctx context.Context, favouriteDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	Store(ctx context.Context, favouriteDomain *Domain) (Domain, error)
	Delete(ctx context.Context, favouriteDomain *Domain) (Domain, error)
}
