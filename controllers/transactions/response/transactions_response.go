package response

import (
	"injar/repository/databases/users"
	"injar/repository/databases/webinars"
	"injar/usecase/transactions"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        int               `json:"id"`
	UserID    int               `json:"user_id"`
	Users     users.Users       `json:"users"`
	WebinarID int               `json:"webinar_id"`
	Webinars  webinars.Webinars `json:"webinars"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		ID:        domain.ID,
		UserID:    domain.UserID,
		WebinarID: domain.WebinarID,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
