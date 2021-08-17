package response

import (
	"injar/repository/databases/webinars"
	"injar/usecase/favourites"
	"time"

	"gorm.io/gorm"
)

type Favourite struct {
	ID        int               `json:"id"`
	UserID    int               `json:"user_id"`
	WebinarID int               `json:"webinar_id"`
	Webinars  webinars.Webinars `json:"webinars"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

func FromDomain(domain favourites.Domain) Favourite {
	return Favourite{
		ID:        domain.ID,
		UserID:    domain.UserID,
		WebinarID: domain.WebinarID,
		Webinars:  domain.Webinars,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
