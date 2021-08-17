package favourites

import (
	"injar/repository/databases/users"
	"injar/repository/databases/webinars"
	"injar/usecase/favourites"

	"time"

	"gorm.io/gorm"
)

type Favourites struct {
	ID        int               `json:"id"`
	UserID    int               `json:"user_id"`
	Users     users.Users       `json:"users" gorm:"foreignKey:UserID;references:ID"`
	WebinarID int               `json:"webinar_id"`
	Webinars  webinars.Webinars `json:"webinars" gorm:"foreignKey:WebinarID;references:ID"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

func (rec *Favourites) toDomain() favourites.Domain {
	return favourites.Domain{
		ID:        rec.ID,
		UserID:    rec.UserID,
		Users:     rec.Users,
		WebinarID: rec.WebinarID,
		Webinars:  rec.Webinars,
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
		Webinars:  favouritesDomain.Webinars,
		CreatedAt: favouritesDomain.CreatedAt,
		UpdatedAt: favouritesDomain.UpdatedAt,
		DeletedAt: favouritesDomain.DeletedAt,
	}
}
