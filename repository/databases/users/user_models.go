package users

import (
	"injar/usecase/users"

	"time"
)

type Users struct {
	ID        int
	Name      string
	Email     string
	Username  string
	Password  string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Username:  rec.Username,
		Password:  rec.Password,
		ImageUrl:  rec.ImageUrl,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		ImageUrl:  userDomain.ImageUrl,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
