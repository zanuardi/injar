package transactions

import (
	"injar/repository/databases/users"
	"injar/repository/databases/webinars"
	"injar/usecase/transactions"

	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	ID        int
	UserID    int
	Users     users.Users `gorm:"foreignKey:UserID;references:ID"`
	WebinarID int
	Webinars  webinars.Webinars `gorm:"foreignKey:WebinarID;references:ID"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Transactions) toDomain() transactions.Domain {
	return transactions.Domain{
		ID:        rec.ID,
		UserID:    rec.UserID,
		WebinarID: rec.WebinarID,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(transactionsDomain transactions.Domain) *Transactions {
	return &Transactions{
		ID:        transactionsDomain.ID,
		UserID:    transactionsDomain.UserID,
		WebinarID: transactionsDomain.WebinarID,
		Status:    transactionsDomain.Status,
		CreatedAt: transactionsDomain.CreatedAt,
		UpdatedAt: transactionsDomain.UpdatedAt,
		DeletedAt: transactionsDomain.DeletedAt,
	}
}
