package response

import (
	"injar/usecase/favourites"
	"time"

	"gorm.io/gorm"
)

type Favourite struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	WebinarID   int            `json:"webinar_id"`
	WebinarName string         `json:"webinar_name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain favourites.Domain) Favourite {
	return Favourite{
		ID:          domain.ID,
		UserID:      domain.UserID,
		WebinarID:   domain.WebinarID,
		WebinarName: domain.WebinarName,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}
