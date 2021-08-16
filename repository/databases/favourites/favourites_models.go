package favourites

import (
	"injar/repository/databases/users"
	"injar/repository/databases/webinars"
	"injar/usecase/favourites"

	"time"

	"gorm.io/gorm"
)

type Favourites struct {
	ID        int
	UserID    int
	Users     users.Users `gorm:"foreignKey:UserID;references:ID"`
	WebinarID int
	Webinars  webinars.Webinars `gorm:"foreignKey:WebinarID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Favourites) toDomain() favourites.Domain {
	return favourites.Domain{
		ID:        rec.ID,
		UserID:    rec.UserID,
		WebinarID: rec.WebinarID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(favouritesDomain favourites.Domain) *Favourites {
	return &Favourites{
		ID:        favouritesDomain.ID,
		UserID:    favouritesDomain.UserID,
		WebinarID: favouritesDomain.WebinarID,
		CreatedAt: favouritesDomain.CreatedAt,
		UpdatedAt: favouritesDomain.UpdatedAt,
		DeletedAt: favouritesDomain.DeletedAt,
	}
}
